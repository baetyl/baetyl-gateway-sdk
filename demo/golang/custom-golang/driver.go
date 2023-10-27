package custom

import (
	"context"
	"encoding/json"
	"fmt"

	plugin "github.com/baetyl/baetyl-gateway-sdk/sdk/golang"
	dm "github.com/baetyl/baetyl-go/v2/dmcontext"
	"github.com/baetyl/baetyl-go/v2/spec/v1"
	"github.com/baetyl/baetyl-go/v2/utils"
)

var _ plugin.Driver = &Driver{}

type Driver struct {
	driverName string
	configPath string
	config     *Config
	report     plugin.Report
	custom     *Custom
}

// NewDriver 初始化驱动，注册到插件，由 go-plugin 框架调用
// 插件启动后，gateway/driver.go 会依次调用 Setup() SetConfig() Start() 来启动驱动
func NewDriver(_ context.Context, cfg *plugin.BackendConfig) (plugin.Driver, error) {
	d := &Driver{
		driverName: cfg.DriverName,
	}
	InitL(cfg.Log)
	L().Debug("NewDriver function")
	return d, nil
}

// GetDriverInfo 获取驱动信息
func (d *Driver) GetDriverInfo(_ *plugin.Request) (*plugin.Response, error) {
	L().Debug("GetDriverInfo function")
	info := map[string]any{
		"name":       d.driverName,
		"configPath": d.configPath,
		"config":     d.config,
	}
	dt, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}
	return &plugin.Response{Data: string(dt)}, nil
}

// SetConfig 设置自定义驱动的配置文件路径
func (d *Driver) SetConfig(req *plugin.Request) (*plugin.Response, error) {
	L().Debug("SetConfig function")
	d.configPath = req.Req
	return &plugin.Response{Data: fmt.Sprintf("Plugin %s SetConfig success", d.driverName)}, nil
}

// Setup 设置驱动名称及上报函数的实现
func (d *Driver) Setup(cfg *plugin.BackendConfig) (*plugin.Response, error) {
	L().Debug("Setup function")

	d.driverName = cfg.DriverName
	d.report = cfg.ReportSvc
	return &plugin.Response{Data: fmt.Sprintf("Plugin %s Setup success", d.driverName)}, nil
}

// Start 启动驱动
func (d *Driver) Start(_ *plugin.Request) (*plugin.Response, error) {
	L().Debug("Start function")

	dm.Run(func(ctx dm.Context) error {
		err := ctx.LoadDriverConfig(d.configPath, d.driverName)
		if err != nil {
			return err
		}
		cfg, err := generateConfig(ctx, d.driverName)
		if err != nil {
			return err
		}
		d.config = cfg

		L().Debug("Start config", cfg)

		d.custom, err = NewCustom(ctx, cfg, d.report)
		if err != nil {
			L().Debug("Start NewCustom error", err)
			return err
		}

		d.custom.Start()
		return nil
	})
	return nil, nil
}

func (d *Driver) Restart(_ *plugin.Request) (*plugin.Response, error) {
	L().Debug("Restart function")
	d.custom.Restart()
	return nil, nil
}

func (d *Driver) Stop(_ *plugin.Request) (*plugin.Response, error) {
	L().Debug("Stop function")
	d.custom.Stop()
	return nil, nil
}

func (d *Driver) Get(req *plugin.Request) (*plugin.Response, error) {
	L().Debug("Get function")
	L().Debug("request", req)

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

	info, err := d.custom.ctx.GetDevice(driverName, devName)
	if err != nil {
		return nil, err
	}

	switch msg.Kind {
	case v1.MessageDeviceEvent:
		var event dm.Event
		if err = msg.Content.Unmarshal(&event); err != nil {
			return nil, err
		}
		return nil, d.custom.Event(info, &event)
	case v1.MessageDevicePropertyGet:
		return nil, d.custom.PropertyGet(info)
	default:
		return nil, ErrMessageTypeNotSupported
	}
}

func (d *Driver) Set(req *plugin.Request) (*plugin.Response, error) {
	L().Debug("Set function")
	L().Debug("request", req)

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

	info, err := d.custom.ctx.GetDevice(driverName, devName)
	if err != nil {
		return nil, err
	}

	var props map[string]any
	if err = msg.Content.Unmarshal(&props); err != nil {
		return nil, err
	}

	return nil, d.custom.Set(info, props)
}

func generateConfig(ctx dm.Context, driverName string) (*Config, error) {
	cfg := &Config{}
	var devices []DeviceConfig
	var jobs []Job

	for _, info := range ctx.GetAllDevices(driverName) {
		accessConfig := info.AccessConfig
		if accessConfig.Custom == nil {
			continue
		}

		device := DeviceConfig{
			Device: info.Name,
		}

		devices = append(devices, device)

		var jobProps []Property

		tpl, err := ctx.GetAccessTemplates(driverName, info.AccessTemplate)
		if err != nil {
			return nil, err
		}
		if tpl != nil && tpl.Properties != nil {
			for _, prop := range tpl.Properties {
				if visitor := prop.Visitor.Custom; visitor != nil {
					// 自定义的针对一个设备的点位信息数据解析
					// e.g. {"name": "pressure", "type": "float32", "index": 2}
					var jobProp Property
					err = json.Unmarshal([]byte(*visitor), &jobProp)
					if err != nil {
						return nil, err
					}
					jobProps = append(jobProps, jobProp)
				}
			}
		}

		// 自定义针对一个设备的采集配置的解析
		// e.g. {"interval": 3000000000}
		var job Job
		err = json.Unmarshal([]byte(*accessConfig.Custom), &job)
		if err != nil {
			return nil, err
		}
		job.Device = info.Name
		job.Properties = jobProps

		jobs = append(jobs, job)
	}
	cfg.Jobs = jobs
	cfg.Devices = devices
	cfg.DriverName = driverName
	if err := utils.SetDefaults(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
