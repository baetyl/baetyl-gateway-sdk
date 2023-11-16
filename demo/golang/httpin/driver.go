package httpin

import (
	"context"
	"encoding/json"
	"fmt"

	dm "github.com/baetyl/baetyl-go/v2/dmcontext"
	"github.com/baetyl/baetyl-go/v2/utils"

	"github.com/baetyl/baetyl-gateway-sdk/sdk/golang"
)

var _ plugin.Driver = &Driver{}

type Driver struct {
	driverName string
	configPath string
	config     *Config
	report     plugin.Report
	engine     *Engine
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

		d.engine, err = NewEngine(cfg, d.report)
		if err != nil {
			L().Debug("driver NewEngine error", err)
			return err
		}

		d.engine.Start()
		return nil
	})
	return nil, nil
}

func (d *Driver) Restart(_ *plugin.Request) (*plugin.Response, error) {
	L().Debug("Restart function")
	d.engine.Restart()
	return nil, nil
}

func (d *Driver) Stop(_ *plugin.Request) (*plugin.Response, error) {
	L().Debug("Stop function")
	d.engine.Stop()
	return nil, nil
}

func (d *Driver) Get(_ *plugin.Request) (*plugin.Response, error) {
	L().Debug("Get function")
	return nil, ErrEventTypeNotSupported
}

func (d *Driver) Set(_ *plugin.Request) (*plugin.Response, error) {
	L().Debug("Set function")
	return nil, ErrEventTypeNotSupported
}

func generateConfig(ctx dm.Context, driverName string) (*Config, error) {
	cfg := &Config{}
	var devs []DeviceConfig

	for _, info := range ctx.GetAllDevices(driverName) {
		accessConfig := info.AccessConfig
		if accessConfig.Custom == nil {
			continue
		}

		var sc ServerConfig
		err := json.Unmarshal([]byte(*info.AccessConfig.Custom), &sc)
		if err != nil {
			return nil, err
		}
		devs = append(devs, DeviceConfig{
			ServerConfig: sc,
			DeviceName:   info.Name,
		})
	}
	cfg.Devices = devs
	cfg.DriverName = driverName
	if err := utils.SetDefaults(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
