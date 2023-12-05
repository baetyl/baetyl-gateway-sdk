package bacnet

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	plugin "github.com/baetyl/baetyl-gateway-sdk/sdk/golang"
	"github.com/baetyl/baetyl-gateway/config"
	dm "github.com/baetyl/baetyl-go/v2/dmcontext"
	"github.com/baetyl/baetyl-go/v2/spec/v1"
	"github.com/baetyl/baetyl-go/v2/utils"
	"github.com/hashicorp/go-hclog"
	"github.com/jinzhu/copier"
)

var _ plugin.Driver = &Driver{}

type Driver struct {
	driverName string
	configPath string
	report     plugin.Report
	bacnet     *Bacnet
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

		d.bacnet, err = NewBacnet(d, ctx, cfg)
		if err != nil {
			return err
		}
		d.bacnet.Start()
		return nil
	})
	return nil, nil
}

// Restart 驱动重启
func (d *Driver) Restart(req *plugin.Request) (*plugin.Response, error) {
	d.bacnet.Restart()
	return nil, nil
}

// Stop 驱动停止
func (d *Driver) Stop(req *plugin.Request) (*plugin.Response, error) {
	d.bacnet.Stop()
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

	devInfo, err := d.bacnet.ctx.GetDevice(driverName, devName)
	if err != nil {
		return nil, err
	}
	switch msg.Kind {
	case v1.MessageDeviceEvent:
		var event dm.Event
		if err = msg.Content.Unmarshal(&event); err != nil {
			return nil, err
		}
		return nil, d.bacnet.Event(devInfo, &event)
	case v1.MessageDevicePropertyGet:
		var props []string
		if err = msg.Content.Unmarshal(&props); err != nil {
			return nil, err
		}
		return nil, d.bacnet.PropertyGet(devInfo, props)
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
	var props []config.DeviceProperty

	if driverName, ok = msg.Metadata[dm.KeyDriverName]; !ok {
		return nil, ErrDriverNameNotExist
	}
	if devName, ok = msg.Metadata[dm.KeyDeviceName]; !ok {
		return nil, ErrDevNameNotExist
	}

	if err = msg.Content.Unmarshal(&props); err != nil {
		return nil, err
	}

	devInfo, err := d.bacnet.ctx.GetDevice(driverName, devName)
	if err != nil {
		return nil, err
	}
	return nil, d.bacnet.Set(driverName, devInfo, props)
}

func genConfig(ctx dm.Context, driverName string) (*Config, error) {
	cfg := &Config{}
	var slaves []SlaveConfig
	var jobs []Job

	// generate job
	for _, devInfo := range ctx.GetAllDevices(driverName) {
		accessConfig := devInfo.AccessConfig
		if accessConfig.Bacnet == nil {
			continue
		}
		slave := SlaveConfig{
			Device: devInfo.Name,
		}
		if err := copier.Copy(&slave, accessConfig.Bacnet); err != nil {
			return nil, err
		}
		slaves = append(slaves, slave)

		// generate jobMap
		jobMaps := make(map[string]Property)
		devTpl, err := ctx.GetAccessTemplates(driverName, devInfo.AccessTemplate)
		if err != nil {
			return nil, err
		}
		if devTpl != nil && devTpl.Properties != nil && len(devTpl.Properties) > 0 {
			for _, prop := range devTpl.Properties {
				if visitor := prop.Visitor.Bacnet; visitor != nil {
					var jobMap Property
					jobMap.Id = prop.ID
					jobMap.Name = prop.Name
					jobMap.Type = prop.Type
					jobMap.Mode = prop.Mode
					jobMap.BacnetType = visitor.BacnetType
					jobMap.ApplicationTagNumber = visitor.ApplicationTagNumber
					jobMap.BacnetAddress = visitor.BacnetAddress + accessConfig.Bacnet.AddressOffset
					jobMaps[strconv.FormatUint(uint64(jobMap.BacnetType), 10)+":"+
						strconv.FormatUint(uint64(jobMap.BacnetAddress), 10)] = jobMap
				}
			}
		}
		job := Job{
			Device:        devInfo.Name,
			Interval:      accessConfig.Bacnet.Interval,
			Properties:    jobMaps,
			DeviceId:      accessConfig.Bacnet.DeviceID,
			AddressOffset: accessConfig.Bacnet.AddressOffset,
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
