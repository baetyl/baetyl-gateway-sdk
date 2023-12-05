package bacnet

import (
	"context"
	"encoding/json"
	"math"
	"time"

	plugin "github.com/baetyl/baetyl-gateway-sdk/sdk/golang"
	dm "github.com/baetyl/baetyl-go/v2/dmcontext"
	"github.com/baetyl/baetyl-go/v2/spec/v1"
	"github.com/chen-Leo/bacnet"
	"github.com/chen-Leo/bacnet/bacip"
	log "github.com/hashicorp/go-hclog"
)

type Worker struct {
	ctx    dm.Context
	driver *Driver
	job    Job
	slave  *Slave
	log    log.Logger
}

func NewWorker(d *Driver, ctx dm.Context, job Job, slave *Slave, log log.Logger) *Worker {
	return &Worker{
		driver: d,
		ctx:    ctx,
		job:    job,
		slave:  slave,
		log:    log,
	}
}

func (w *Worker) Execute(report bool) error {
	msg := v1.Message{}
	temp := make(map[string]interface{})
	for _, prop := range w.job.Properties {
		objID := bacnet.ObjectID{
			Type:     bacnet.ObjectType(prop.BacnetType),
			Instance: bacnet.ObjectInstance(prop.BacnetAddress),
		}
		d, err := readValue(w.slave.bacnetClient, w.slave.device, objID)
		if err != nil {
			w.log.Error("failed to read", "error", err, "device", w.slave.device, "prop", prop.Name)
			if err1 := w.slave.UpdateStatus(SlaveOffline); err1 != nil {
				w.log.Error("failed to update slave status", "error", err1, "status", "offline")
			}
			continue
		}
		v, _ := dm.ParseValueToFloat64(d)
		if math.IsInf(v, 0) || math.IsNaN(v) {
			w.log.Warn("value Inf or NaN adopted", "value", v)
			continue
		}
		temp[prop.Name] = d
	}

	if report {
		msg.Kind = v1.MessageDeviceReport
	} else {
		msg.Kind = v1.MessageDeviceDesire
	}
	msg.Metadata = make(map[string]string)
	msg.Metadata[dm.KeyDriverName] = w.driver.driverName
	msg.Metadata[dm.KeyDeviceName] = w.slave.info.Name
	msg.Content = v1.LazyValue{Value: temp}

	msgData, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	res, err := w.driver.report.Post(&plugin.Request{Req: string(msgData)})
	if err != nil {
		return err
	}
	w.log.Debug("bacnet driver report message", "msgdata", string(msgData), "rsp", res.Data)

	if err = w.slave.UpdateStatus(SlaveOnline); err != nil {
		w.log.Error("failed to update slave status", "error", err, "status", "online")
	}
	return nil
}

func readValue(c *bacip.Client, device bacnet.Device, object bacnet.ObjectID) (interface{}, error) {
	rp := bacip.ReadProperty{
		ObjectID: object,
		Property: bacnet.PropertyIdentifier{
			Type: bacnet.PresentValue,
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	d, err := c.ReadProperty(ctx, device, rp)
	if err != nil {
		return nil, err
	}
	return d, nil
}
