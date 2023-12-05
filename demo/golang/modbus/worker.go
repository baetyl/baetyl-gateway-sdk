package modbus

import (
	"encoding/json"
	"math"
	"time"

	plugin "github.com/baetyl/baetyl-gateway-sdk/sdk/golang"
	dm "github.com/baetyl/baetyl-go/v2/dmcontext"
	"github.com/baetyl/baetyl-go/v2/spec/v1"
	log "github.com/hashicorp/go-hclog"
)

type Worker struct {
	ctx    dm.Context
	driver *Driver
	job    Job
	maps   []*Map
	slave  *Slave
	log    log.Logger
}

func NewWorker(d *Driver, ctx dm.Context, job Job, slave *Slave, log log.Logger) *Worker {
	w := &Worker{
		driver: d,
		ctx:    ctx,
		job:    job,
		slave:  slave,
		log:    log,
	}
	for _, v := range job.Maps {
		m := NewMap(ctx, v, slave, log)
		w.maps = append(w.maps, m)
	}
	return w
}

func (w *Worker) Execute(report bool) error {
	msg := v1.Message{}
	temp := make(map[string]any)
	cur := time.Now()
	for _, m := range w.maps {
		p, err := m.Collect()
		if err != nil {
			if err1 := w.slave.UpdateStatus(SlaveOffline); err1 != nil {
				w.log.Error("failed to update slave status", "error", err1, "device", "offline")
			}
			return err
		}
		pa, err := m.Parse(p[4:])
		if err != nil {
			return err
		}
		v, _ := dm.ParseValueToFloat64(pa)
		if math.IsInf(v, 0) || math.IsNaN(v) {
			w.log.Debug("value Inf or NaN adopted", "origin", string(p), "props", m.cfg.Name, "value", v)
			continue
		}
		temp[m.cfg.Name] = pa
	}

	cost := time.Since(cur).String()
	w.log.Debug("collect single device cost time", "cost", cost, "device", w.slave.dev.Name)
	// Report or property get
	if report {
		msg.Kind = v1.MessageDeviceReport
	} else {
		msg.Kind = v1.MessageDeviceDesire
	}
	msg.Metadata = make(map[string]string)
	msg.Metadata[dm.KeyDriverName] = w.driver.driverName
	msg.Metadata[dm.KeyDeviceName] = w.slave.dev.Name
	msg.Content = v1.LazyValue{Value: temp}

	msgData, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	res, err := w.driver.report.Post(&plugin.Request{Req: string(msgData)})
	if err != nil {
		return err
	}
	w.log.Debug("modbus driver report message", "msgdata", string(msgData), "rsp", res.Data)
	if err := w.slave.UpdateStatus(SlaveOnline); err != nil {
		w.log.Error("failed to update slave status", "error", err, "status", "online")
	}
	return nil
}
