package opcua

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
	opcua      *Opcua
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

		d.opcua, err = NewOpcua(d, ctx, cfg)
		if err != nil {
			return err
		}
		d.opcua.Start()
		return nil
	})
	return nil, nil
}

// Restart 驱动重启
func (d *Driver) Restart(req *plugin.Request) (*plugin.Response, error) {
	d.opcua.Restart()
	return nil, nil
}

// Stop 驱动停止
func (d *Driver) Stop(req *plugin.Request) (*plugin.Response, error) {
	d.opcua.Stop()
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

	devInfo, err := d.opcua.ctx.GetDevice(driverName, devName)
	if err != nil {
		return nil, err
	}
	switch msg.Kind {
	case v1.MessageDeviceEvent:
		var event dm.Event
		if err = msg.Content.Unmarshal(&event); err != nil {
			return nil, err
		}
		return nil, d.opcua.Event(devInfo, &event)
	case v1.MessageDevicePropertyGet:
		var props []string
		if err = msg.Content.Unmarshal(&props); err != nil {
			return nil, err
		}
		return nil, d.opcua.PropertyGet(devInfo, props)
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

	devInfo, err := d.opcua.ctx.GetDevice(driverName, devName)
	if err != nil {
		return nil, err
	}
	return nil, d.opcua.Set(driverName, devInfo, props)
}

func genConfig(ctx dm.Context, driverName string) (*Config, error) {
	cfg := &Config{}
	var devices []DeviceConfig
	var jobs []Job

	for _, deviceInfo := range ctx.GetAllDevices(driverName) {
		accessConfig := deviceInfo.AccessConfig
		if accessConfig.Opcua == nil {
			continue
		}
		device := DeviceConfig{
			Device: deviceInfo.Name,
		}
		if err := copier.Copy(&device, accessConfig.Opcua); err != nil {
			return nil, err
		}
		devices = append(devices, device)

		var jobProps []Property
		deviceTemplate, _ := ctx.GetAccessTemplates(driverName, deviceInfo.AccessTemplate)
		if deviceTemplate != nil && deviceTemplate.Properties != nil && len(deviceTemplate.Properties) > 0 {
			for _, prop := range deviceTemplate.Properties {
				if visitor := prop.Visitor.Opcua; visitor != nil {
					var nodeId string
					ns := deviceInfo.AccessConfig.Opcua.NsOffset + visitor.NsBase
					switch visitor.IDType {
					case OpcuaIdTypeI:
						idBase, err := strconv.Atoi(visitor.IDBase)
						if err != nil {
							continue
						}
						nodeId = fmt.Sprintf("ns=%d;i=%d", ns, deviceInfo.AccessConfig.Opcua.IDOffset+idBase)
					case OpcuaIdTypeS, OpcuaIdTypeG, OpcuaIdTypeB:
						nodeId = fmt.Sprintf("ns=%d;%s=%s", ns, visitor.IDType, visitor.IDBase)
					default:
						continue
					}
					jobProps = append(jobProps, Property{
						Name:   prop.Name,
						Type:   visitor.Type,
						NodeID: nodeId,
					})
				}
			}
		}
		job := Job{
			Device:     deviceInfo.Name,
			Subscribe:  accessConfig.Opcua.Subscribe,
			Interval:   accessConfig.Opcua.Interval,
			Properties: jobProps,
		}
		jobs = append(jobs, job)
	}
	cfg.Devices = devices
	cfg.Jobs = jobs
	cfg.DriverName = driverName
	if err := utils.SetDefaults(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
