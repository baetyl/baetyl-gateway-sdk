package modbus

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	plugin "github.com/baetyl/baetyl-gateway-sdk/sdk/golang"
	cfg "github.com/baetyl/baetyl-gateway/config"
	dm "github.com/baetyl/baetyl-go/v2/dmcontext"
	"github.com/baetyl/baetyl-go/v2/spec/v1"
	"github.com/baetyl/baetyl-go/v2/utils"
	"github.com/hashicorp/go-hclog"
)

var _ plugin.Driver = &Driver{}

type Driver struct {
	driverName string
	configPath string
	report     plugin.Report
	mds        *Modbus
	log        hclog.Logger
}

func NewDriver(_ context.Context, cfg *plugin.BackendConfig) (plugin.Driver, error) {
	b := &Driver{log: cfg.Log}
	return b, nil
}

func (d *Driver) GetDriverInfo(req *plugin.Request) (*plugin.Response, error) {
	return nil, nil
}

// SetConfig 配置驱动配置路径
func (d *Driver) SetConfig(req *plugin.Request) (*plugin.Response, error) {
	d.configPath = req.Req
	return &plugin.Response{Data: fmt.Sprintf("Plugin %s SetConfig success", d.driverName)}, nil
}

// Setup 宿主进程上报接口传递
func (d *Driver) Setup(config *plugin.BackendConfig) (*plugin.Response, error) {
	d.driverName = config.DriverName
	d.report = config.ReportSvc
	return &plugin.Response{Data: fmt.Sprintf("Plugin %s Setup success", d.driverName)}, nil
}

// Start 驱动采集启动
func (d *Driver) Start(req *plugin.Request) (*plugin.Response, error) {
	dm.Run(func(ctx dm.Context) error {
		err := ctx.LoadDriverConfig(d.configPath, d.driverName)
		if err != nil {
			return err
		}
		cfg, err := genConfig(ctx, d.driverName)
		if err != nil {
			return err
		}

		d.mds, err = NewModbus(d, ctx, cfg)
		if err != nil {
			return err
		}
		d.mds.Start()
		return nil
	})
	return nil, nil
}

// Restart 驱动重启
func (d *Driver) Restart(req *plugin.Request) (*plugin.Response, error) {
	d.mds.Restart()
	return nil, nil
}

// Stop 驱动停止
func (d *Driver) Stop(req *plugin.Request) (*plugin.Response, error) {
	d.mds.Stop()
	return nil, nil
}

// Get 召测
func (d *Driver) Get(req *plugin.Request) (*plugin.Response, error) {
	msg := &v1.Message{}
	err := json.Unmarshal([]byte(req.Req), msg)
	if err != nil {
		return nil, err
	}
	var driverName, devName string
	var ok bool

	if driverName, ok = msg.Metadata[dm.KeyDriverName]; !ok {
		return nil, ErrDriverNameNotExist
	}
	if devName, ok = msg.Metadata[dm.KeyDeviceName]; !ok {
		return nil, ErrDevNameNotExist
	}

	devInfo, err := d.mds.ctx.GetDevice(driverName, devName)
	if err != nil {
		return nil, err
	}
	switch msg.Kind {
	case v1.MessageDeviceEvent:
		var event dm.Event
		if err = msg.Content.Unmarshal(&event); err != nil {
			return nil, err
		}
		return nil, d.mds.Event(devInfo, &event)
	case v1.MessageDevicePropertyGet:
		var props []string
		if err = msg.Content.Unmarshal(&props); err != nil {
			return nil, err
		}
		return nil, d.mds.PropertyGet(devInfo, props)
	default:
		return nil, errors.New("unexpected message kind")
	}
}

// Set 置数
func (d *Driver) Set(req *plugin.Request) (*plugin.Response, error) {
	msg := &v1.Message{}
	err := json.Unmarshal([]byte(req.Req), msg)
	if err != nil {
		return nil, err
	}
	var driverName, devName string
	var ok bool
	var props []cfg.DeviceProperty

	if driverName, ok = msg.Metadata[dm.KeyDriverName]; !ok {
		return nil, ErrDriverNameNotExist
	}
	if devName, ok = msg.Metadata[dm.KeyDeviceName]; !ok {
		return nil, ErrDevNameNotExist
	}

	if err = msg.Content.Unmarshal(&props); err != nil {
		return nil, err
	}

	devInfo, err := d.mds.ctx.GetDevice(driverName, devName)
	if err != nil {
		return nil, err
	}
	return nil, d.mds.Set(driverName, devInfo, props)
}

func genConfig(ctx dm.Context, driverName string) (*Config, error) {
	cfg := &Config{}
	var slaves []SlaveConfig
	var jobs []Job

	for _, deviceInfo := range ctx.GetAllDevices(driverName) {
		accessConfig := deviceInfo.AccessConfig
		if accessConfig.Modbus == nil {
			continue
		}
		slave := SlaveConfig{
			Device:      deviceInfo.Name,
			Id:          accessConfig.Modbus.ID,
			Timeout:     accessConfig.Modbus.Timeout,
			IdleTimeout: accessConfig.Modbus.IdleTimeout,
		}
		if tcp := accessConfig.Modbus.TCP; tcp != nil {
			slave.Mode = string(ModeTCP)
			slave.Address = fmt.Sprintf("%s:%d", tcp.Address, tcp.Port)
		} else if rtu := accessConfig.Modbus.RTU; rtu != nil {
			slave.Mode = string(ModeRTU)
			slave.Address = rtu.Port
			slave.BaudRate = rtu.BaudRate
			slave.DataBits = rtu.DataBit
			slave.StopBits = rtu.StopBit
			slave.Parity = rtu.Parity
		}
		slaves = append(slaves, slave)

		var jobMaps []MapConfig
		deviceTemplate, err := ctx.GetAccessTemplates(driverName, deviceInfo.AccessTemplate)
		if err != nil {
			return nil, err
		}
		if deviceTemplate != nil && deviceTemplate.Properties != nil && len(deviceTemplate.Properties) > 0 {
			for _, prop := range deviceTemplate.Properties {
				if visitor := prop.Visitor.Modbus; visitor != nil {
					address, aErr := strconv.ParseUint(visitor.Address[2:], 16, 16)
					if aErr != nil {
						return nil, aErr
					}
					m := MapConfig{
						Id:           prop.ID,
						Name:         prop.Name,
						Type:         visitor.Type,
						Function:     visitor.Function,
						Address:      uint16(address),
						Quantity:     visitor.Quantity,
						SwapRegister: visitor.SwapRegister,
						SwapByte:     visitor.SwapByte,
					}
					jobMaps = append(jobMaps, m)
				}
			}
		}
		job := Job{
			Device:   deviceInfo.Name,
			Interval: accessConfig.Modbus.Interval,
			Maps:     jobMaps,
		}
		jobs = append(jobs, job)
	}
	cfg.Jobs = jobs
	cfg.Slaves = slaves
	cfg.DriverName = driverName
	if err := utils.SetDefaults(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
