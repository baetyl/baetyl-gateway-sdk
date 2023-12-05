package opcua

import (
	"context"
	"encoding/json"
	"time"

	plugin "github.com/baetyl/baetyl-gateway-sdk/sdk/golang"
	dm "github.com/baetyl/baetyl-go/v2/dmcontext"
	"github.com/baetyl/baetyl-go/v2/errors"
	"github.com/baetyl/baetyl-go/v2/spec/v1"
	"github.com/gopcua/opcua"
	"github.com/gopcua/opcua/monitor"
	"github.com/gopcua/opcua/ua"
	log "github.com/hashicorp/go-hclog"
)

type Worker struct {
	ctx    dm.Context
	driver *Driver
	job    Job
	device *Device
	log    log.Logger
}

func NewWorker(d *Driver, ctx dm.Context, job Job, device *Device, logger log.Logger) *Worker {
	w := &Worker{
		device: device,
		job:    job,
		ctx:    ctx,
		driver: d,
		log:    logger,
	}
	return w
}

func (w *Worker) Execute(report bool) error {
	msg := v1.Message{}
	cur := time.Now()

	temp, err := w.read(w.job.Properties)
	if err != nil {
		if sErr := w.device.UpdateStatus(DeviceOffline); sErr != nil {
			w.log.Error("failed to update device status", "error", sErr, "status", "offline")
		}
		w.log.Error("failed to read", "error", err)
		return err
	}

	cost := time.Since(cur).String()
	w.log.Debug("collect single device cost time", "cost", cost, "device", w.job.Device)

	// Report or property get
	if report {
		msg.Kind = v1.MessageDeviceReport
	} else {
		msg.Kind = v1.MessageDeviceDesire
	}
	msg.Metadata = make(map[string]string)
	msg.Metadata[dm.KeyDriverName] = w.driver.driverName
	msg.Metadata[dm.KeyDeviceName] = w.device.info.Name
	msg.Content = v1.LazyValue{Value: temp}

	msgData, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	res, err := w.driver.report.Post(&plugin.Request{Req: string(msgData)})
	if err != nil {
		return err
	}
	w.log.Debug("opcua driver report message", "msgdata", string(msgData), "rsp", res.Data)
	if err := w.device.UpdateStatus(DeviceOnline); err != nil {
		w.log.Error("failed to update device status", "error", err, "status", "online")
	}
	return nil
}

func (w *Worker) read(props []Property) (map[string]any, error) {
	var ids []*ua.ReadValueID
	var fP []FilterProps
	temp := make(map[string]any)

	for _, prop := range props {
		id, err := ua.ParseNodeID(prop.NodeID)
		if err != nil {
			w.log.Error("invalid node id", "nodeid", prop.NodeID)
			continue
		}
		readID := &ua.ReadValueID{
			NodeID: id,
		}
		ids = append(ids, readID)

		var filterProp FilterProps
		filterProp.Name = prop.Name
		filterProp.Type = prop.Type
		fP = append(fP, filterProp)
	}

	req := &ua.ReadRequest{
		MaxAge:             2000,
		NodesToRead:        ids,
		TimestampsToReturn: ua.TimestampsToReturnNeither,
	}
	resp, err := w.device.opcuaClient.Read(req)
	if err != nil {
		w.log.Error("failed to read", "nodeid", ids, "error", err)
		return nil, err
	}
	if resp == nil || len(resp.Results) == 0 {
		w.log.Error("invalid read response", "nodeid", ids)
		return nil, errors.Errorf("invalid read response")
	}

	for i, r := range resp.Results {
		if r.Status != ua.StatusOK {
			w.log.Error("Node Status Not OK", "props", fP[i].Name)
			continue
		}
		value, err := variant2Value(fP[i].Type, r.Value)
		if err != nil {
			w.log.Error("failed to parse", "error", err)
			continue
		}
		temp[fP[i].Name] = value
	}

	return temp, nil
}

func (w *Worker) Subscribe() error {
	m, err := monitor.NewNodeMonitor(w.device.opcuaClient)
	if err != nil {
		w.log.Error("failed to create node monitor")
		return err
	}

	var nodeID []string
	for _, p := range w.job.Properties {
		nodeID = append(nodeID, p.NodeID)
	}

	w.startChanSub(context.Background(), m, 1*time.Second, 0, nodeID...)

	return nil
}

func (w *Worker) startChanSub(ctx context.Context, m *monitor.NodeMonitor, interval, lag time.Duration, nodes ...string) {
	ch := make(chan *monitor.DataChangeMessage, 1024)

	sub, err := m.ChanSubscribe(ctx, &opcua.SubscriptionParameters{Interval: interval}, ch, nodes...)
	if err != nil {
		w.log.Error("failed to start channel subscribe", "error", err)
		return
	}

	defer cleanup(ctx, sub)

	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-ch:
			if msg.Error != nil {
				w.log.Error("subscribe channel message error", "error", msg.Error)
			} else {
				temp := make(map[string]any)
				for _, p := range w.job.Properties {
					if p.NodeID == msg.NodeID.String() {
						temp[p.Name] = msg.Value.Value()
					}
				}
				reportMsg := v1.Message{}
				reportMsg.Kind = v1.MessageDeviceReport
				reportMsg.Metadata = make(map[string]string)
				reportMsg.Metadata[dm.KeyDriverName] = w.driver.driverName
				reportMsg.Metadata[dm.KeyDeviceName] = w.device.info.Name
				reportMsg.Content = v1.LazyValue{Value: temp}

				msgData, err := json.Marshal(reportMsg)
				if err != nil {
					return
				}

				res, err := w.driver.report.Post(&plugin.Request{Req: string(msgData)})
				if err != nil {
					return
				}
				w.log.Debug("opcua driver report message", "msgdata", string(msgData), "rsp", res.Data)
			}
			time.Sleep(lag)
		}
	}
}

func cleanup(ctx context.Context, sub *monitor.Subscription) {
	sub.Unsubscribe(ctx)
}
