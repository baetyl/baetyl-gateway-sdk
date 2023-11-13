package custom

import (
	"github.com/baetyl/baetyl-gateway-sdk/sdk/golang"
	dm "github.com/baetyl/baetyl-go/v2/dmcontext"
	"github.com/baetyl/baetyl-go/v2/utils"
)

type Custom struct {
	cfg  Config
	ctx  dm.Context
	ws   map[string]*Worker
	devs map[string]*Device
	tomb utils.Tomb
}

func NewCustom(ctx dm.Context, cfg *Config, report plugin.Report) (*Custom, error) {
	L().Debug("NewCustom function")

	infos := make(map[string]dm.DeviceInfo)
	for _, info := range ctx.GetAllDevices(cfg.DriverName) {
		infos[info.Name] = info
	}
	L().Debug("NewCustom infos", infos)

	devs := make(map[string]*Device)
	for _, item := range cfg.Devices {
		if info, ok := infos[item.Device]; ok {
			dev, err := NewDevice(&info, item)
			if err != nil {
				L().Error("ignore device which failed to establish connection", "device", item.Device, "error", err)
				continue
			}
			devs[item.Device] = dev

			L().Debug("NewCustom info", info)
		}
	}
	ws := make(map[string]*Worker)
	for _, job := range cfg.Jobs {
		if dev := devs[job.Device]; dev != nil {
			ws[dev.info.Name] = NewWorker(ctx, cfg.DriverName, job, dev, report)
			L().Debug("NewCustom set work", dev.info.Name)
			L().Debug("NewCustom work instance", ws[dev.info.Name])
		} else {
			L().Error("device of job not exist", "device id", job.Device)
		}
	}
	c := &Custom{
		ctx:  ctx,
		ws:   ws,
		devs: devs,
	}
	return c, nil
}

// Start 启动对每个设备的周期性采集上报
func (c *Custom) Start() {
	for _, w := range c.ws {
		c.tomb.Go(func() error {
			w.Working()
			return nil
		})
	}
}

func (c *Custom) Restart() {
	c.Stop()
	c.Start()
}

func (c *Custom) Stop() {
	c.tomb.Kill(nil)
	c.tomb.Wait()
}

func (c *Custom) Set(info *dm.DeviceInfo, props map[string]any) error {
	L().Debug("Set props", props)

	var err error
	w, ok := c.ws[info.Name]
	if !ok {
		L().Warn("worker not exist according to device", "device", info.Name)
		return ErrWorkerNotExist
	}
	for k, v := range props {
		for _, p := range w.job.Properties {
			if k == p.Name {
				var value any
				switch v.(type) {
				case float64:
					value, err = dm.ParsePropertyValue(p.Type, v.(float64))
				default:
					value = v
				}
				if err != nil {
					return err
				}

				w.device.Set(value.(float32), p.Index)
			}
		}
	}
	return nil
}

func (c *Custom) Event(info *dm.DeviceInfo, event *dm.Event) error {
	L().Debug("Event props", event)

	switch event.Type {
	case dm.TypeReportEvent:
		return c.PropertyGet(info)
	default:
		return ErrEventTypeNotSupported
	}
}

func (c *Custom) PropertyGet(info *dm.DeviceInfo) error {
	w, ok := c.ws[info.Name]
	if !ok {
		L().Warn("worker not exist according to device", "device", info.Name)
		return ErrWorkerNotExist
	}
	if err := w.Report(false); err != nil {
		return err
	}
	return nil
}
