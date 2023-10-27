package custom

import (
	"encoding/json"
	"time"

	plugin "github.com/baetyl/baetyl-gateway-sdk/sdk/golang"
	dm "github.com/baetyl/baetyl-go/v2/dmcontext"
	"github.com/baetyl/baetyl-go/v2/spec/v1"
)

type Worker struct {
	ctx        dm.Context
	job        Job
	device     *Device
	driverName string
	report     plugin.Report
}

func NewWorker(ctx dm.Context, driverName string, job Job, device *Device, report plugin.Report) *Worker {
	L().Debug("NewWorker function job", job)
	L().Debug("NewWorker function dev", device)
	w := &Worker{
		ctx:        ctx,
		job:        job,
		device:     device,
		driverName: driverName,
		report:     report,
	}
	return w
}

func (w *Worker) Working() {
	L().Debug("working start")

	ticker := time.NewTicker(w.job.Interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			L().Debug("working report")
			err := w.Report(true)
			if err != nil {
				L().Error("failed to report", "error", err)
			}
		case <-w.ctx.WaitChan():
			L().Warn("worker stopped", "worker", w)
			return
		}
	}
}

func (w *Worker) Report(report bool) error {
	L().Debug("custom driver report start")

	props := make(map[string]any)
	// 根据 job 列表罗列的待采集点的信息
	for _, p := range w.job.Properties {
		val, err := w.device.ReadProperty(p.Name)
		if err != nil {
			L().Error("failed to read", "error", err)
			continue
		}
		value, err := dm.ParseValue(p.Type, val, "")
		if err != nil {
			L().Error("failed to parse", "error", err)
			continue
		}
		props[p.Name] = value
		L().Debug("custom driver parse prop key", p.Name)
		L().Debug("custom driver parse prop value", value)
	}

	msg := v1.Message{
		Metadata: map[string]string{
			dm.KeyDriverName: w.driverName,
			dm.KeyDeviceName: w.device.info.Name,
		},
		Content: v1.LazyValue{Value: props},
	}
	// Report or property get
	if report {
		msg.Kind = v1.MessageDeviceReport
	} else {
		msg.Kind = v1.MessageDeviceDesire
	}

	dt, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	L().Debug("custom driver report meta data json", "msg", string(dt))

	res, err := w.report.Post(&plugin.Request{Req: string(dt)})
	if err != nil {
		return err
	}
	L().Debug("custom driver report message", "msg", string(dt), "rsp", res.Data)

	return nil
}
