package opcua

import (
	"time"

	"github.com/baetyl/baetyl-gateway/config"
	dm "github.com/baetyl/baetyl-go/v2/dmcontext"
	"github.com/baetyl/baetyl-go/v2/errors"
	"github.com/gopcua/opcua/ua"
	"github.com/hashicorp/go-hclog"
	"github.com/panjf2000/ants/v2"
)

type Opcua struct {
	cfg     Config
	ctx     dm.Context
	ws      map[string]*Worker
	devs    map[string]*Device
	pool    *ants.PoolWithFunc
	tickers []*time.Ticker
	log     hclog.Logger
}

func NewOpcua(d *Driver, ctx dm.Context, cfg *Config) (*Opcua, error) {
	infos := make(map[string]dm.DeviceInfo)
	for _, info := range ctx.GetAllDevices(cfg.DriverName) {
		infos[info.Name] = info
	}

	devs := make(map[string]*Device)
	for _, dCfg := range cfg.Devices {
		if info, ok := infos[dCfg.Device]; ok {
			dev, err := NewDevice(d, &info, dCfg)
			if err != nil {
				d.log.Error("ignore device which failed to establish connection", "device", dCfg.Device, "error", err)
				continue
			}
			devs[dCfg.Device] = dev
			if err = dev.UpdateStatus(DeviceOnline); err != nil {
				d.log.Error("failed to update status", "error", err)
			}
		}
	}
	ws := make(map[string]*Worker)
	for _, job := range cfg.Jobs {
		if dev := devs[job.Device]; dev != nil {
			ws[dev.info.Name] = NewWorker(d, ctx, job, dev, d.log)
		} else {
			d.log.Error("device of job not exist", "device id", job.Device)
		}
	}
	o := &Opcua{
		ctx:  ctx,
		ws:   ws,
		devs: devs,
		log:  d.log,
	}
	pool, err := ants.NewPoolWithFunc(DefaultAntsPoolSize, o.working, ants.WithPreAlloc(true))
	if err != nil {
		d.log.Error("failed to allocate goroutine pool", "error", err)
		return nil, err
	}
	o.pool = pool
	return o, nil
}

func (o *Opcua) Start() {
	for _, worker := range o.ws {
		err := o.pool.Invoke(worker)
		if err != nil {
			o.log.Error("failed to invoke opcua worker into go pool", "error", err)
			continue
		}
	}
}

func (o *Opcua) Restart() {
	o.Stop()
	o.Start()
}

func (o *Opcua) Stop() {
	o.log.Debug("opcua stopped")
	for _, ticker := range o.tickers {
		ticker.Stop()
	}
	for _, worker := range o.ws {
		worker.device.opcuaClient.Close()
	}
}

func (o *Opcua) working(w interface{}) {
	worker, ok := w.(*Worker)
	if !ok {
		return
	}
	if worker.job.Subscribe {
		err := worker.Subscribe()
		if err != nil {
			o.log.Error("failed to execute job", "error", err)
		}
	} else {
		ticker := time.NewTicker(worker.job.Interval)
		o.tickers = append(o.tickers, ticker)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				err := worker.Execute(true)
				if err != nil {
					o.log.Error("failed to execute job", "error", err)
				}
			}
		}
	}
}

func (o *Opcua) Set(_ string, info *dm.DeviceInfo, props []config.DeviceProperty) error {
	var err error
	w, ok := o.ws[info.Name]
	if !ok {
		o.log.Warn("worker not exist according to device", "device", info.Name)
		return ErrWorkerNotExist
	}
	for _, prop := range props {
		for _, p := range w.job.Properties {
			if prop.PropName == p.Name {
				var value interface{}
				switch prop.PropVal.(type) {
				case float64:
					value, err = dm.ParsePropertyValue(p.Type, prop.PropVal.(float64))
				default:
					value = prop.PropVal
				}
				if err != nil {
					return errors.Trace(err)
				}
				variant, varErr := ua.NewVariant(value)
				if varErr != nil {
					return errors.Trace(varErr)
				}
				nid, parErr := ua.ParseNodeID(p.NodeID)
				if parErr != nil {
					return errors.Trace(parErr)
				}
				req := &ua.WriteRequest{
					NodesToWrite: []*ua.WriteValue{{
						NodeID:      nid,
						AttributeID: ua.AttributeIDValue,
						Value: &ua.DataValue{
							EncodingMask: ua.DataValueValue,
							Value:        variant,
						}}},
				}
				_, err = w.device.opcuaClient.Write(req)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (o *Opcua) Event(info *dm.DeviceInfo, event *dm.Event) error {
	switch event.Type {
	case dm.TypeReportEvent:
		return o.PropertyGet(info, nil)
	default:
		return ErrEventTypeNotSupported
	}
}

func (o *Opcua) PropertyGet(info *dm.DeviceInfo, _ []string) error {
	w, ok := o.ws[info.Name]
	if !ok {
		o.log.Warn("worker not exist according to device", "device", info.Name)
		return ErrWorkerNotExist
	}
	if err := w.Execute(false); err != nil {
		return err
	}
	return nil
}
