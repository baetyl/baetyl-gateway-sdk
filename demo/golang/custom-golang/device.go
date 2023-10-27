package custom

import (
	dm "github.com/baetyl/baetyl-go/v2/dmcontext"
)

type Device struct {
	cfg  DeviceConfig
	info *dm.DeviceInfo
	cli  *Simulator
}

// NewDevice 此处为模拟设备采集初始化逻辑
func NewDevice(info *dm.DeviceInfo, cfg DeviceConfig) (*Device, error) {
	d := &Device{
		cfg:  cfg,
		info: info,
		cli:  NewSimulator(cfg.Device),
	}
	return d, nil
}

func (d *Device) ReadProperty(name string) (any, error) {
	return d.cli.Get(name)
}

func (d *Device) Set(val float32, idx int) {
	d.cli.Set(val, idx)
}
