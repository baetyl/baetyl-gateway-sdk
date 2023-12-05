package iec104

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	plugin "github.com/baetyl/baetyl-gateway-sdk/sdk/golang"
	"github.com/baetyl/baetyl-gateway/config"
	dm "github.com/baetyl/baetyl-go/v2/dmcontext"
	"github.com/baetyl/baetyl-go/v2/errors"
	"github.com/baetyl/baetyl-go/v2/log"
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
	iec104     *IEC104
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

		d.iec104, err = NewIEC104(d, ctx, cfg)
		if err != nil {
			return err
		}
		d.iec104.Start()
		return nil
	})
	return nil, nil
}

// Restart 驱动重启
func (d *Driver) Restart(req *plugin.Request) (*plugin.Response, error) {
	d.iec104.Restart()
	return nil, nil
}

// Stop 驱动停止
func (d *Driver) Stop(req *plugin.Request) (*plugin.Response, error) {
	d.iec104.Stop()
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

	devInfo, err := d.iec104.ctx.GetDevice(driverName, devName)
	if err != nil {
		return nil, err
	}
	switch msg.Kind {
	case v1.MessageDeviceEvent:
		var event dm.Event
		if err = msg.Content.Unmarshal(&event); err != nil {
			return nil, err
		}
		return nil, d.iec104.Event(devInfo, &event)
	case v1.MessageDevicePropertyGet:
		var props []string
		if err = msg.Content.Unmarshal(&props); err != nil {
			return nil, err
		}
		return nil, d.iec104.PropertyGet(devInfo, props)
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

	devInfo, err := d.iec104.ctx.GetDevice(driverName, devName)
	if err != nil {
		return nil, err
	}
	return nil, d.iec104.Set(driverName, devInfo, props)
}

func genConfig(ctx dm.Context, driverName string) (*Config, error) {
	cfg := &Config{}
	var slaves []SlaveConfig
	var jobs []Job
	// 获取环境变量中的偏移
	pointStart, err := genPointStart(ctx)
	if err != nil {
		return nil, err
	}

	for _, deviceInfo := range ctx.GetAllDevices(driverName) {
		accessConfig := deviceInfo.AccessConfig
		if accessConfig.IEC104 == nil {
			continue
		}
		slave := SlaveConfig{
			Device: deviceInfo.Name,
		}
		if err := copier.Copy(&slave, accessConfig.IEC104); err != nil {
			return nil, err
		}
		slaves = append(slaves, slave)

		var jobProps []Point
		deviceTemplate, _ := ctx.GetAccessTemplates(driverName, deviceInfo.AccessTemplate)

		if deviceTemplate != nil && deviceTemplate.Properties != nil && len(deviceTemplate.Properties) > 0 {
			pointOffset, err := genPointOffset(&deviceInfo)
			if err != nil {
				return nil, err
			}
			for _, prop := range deviceTemplate.Properties {
				if visitor := prop.Visitor.IEC104; visitor != nil {
					jobProps = append(jobProps, Point{
						Name:      prop.Name,
						PointNum:  visitor.PointNum + pointOffset[visitor.PointType] + pointStart[visitor.PointType],
						PointType: visitor.PointType,
						Type:      visitor.Type,
					})
				}
			}
		}
		job := Job{
			Device:   deviceInfo.Name,
			Interval: accessConfig.IEC104.Interval,
			Points:   jobProps,
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

func genPointOffset(deviceInfo *dm.DeviceInfo) (map[string]uint, error) {
	return map[string]uint{
		DI: uint(deviceInfo.AccessConfig.IEC104.DIOffset),
		AI: uint(deviceInfo.AccessConfig.IEC104.AIOffset),
		DO: uint(deviceInfo.AccessConfig.IEC104.DOOffset),
		AO: uint(deviceInfo.AccessConfig.IEC104.AOOffset),
	}, nil
}

func genPointStart(ctx dm.Context) (map[string]uint, error) {
	diStartStr := os.Getenv(DI_Start)
	if len(diStartStr) < 3 || !strings.HasPrefix(diStartStr, "0x") {
		ctx.Log().Warn("env DI_Start is invalid, use default 0x1", log.Any("DI_Start", diStartStr))
		diStartStr = "0x1"
	}
	aiStartStr := os.Getenv(AI_Start)
	if len(aiStartStr) < 3 || !strings.HasPrefix(aiStartStr, "0x") {
		ctx.Log().Warn("env AI_Start is invalid, use default 0x4001", log.Any("AI_Start", aiStartStr))
		aiStartStr = "0x4001"
	}
	piStartStr := os.Getenv(PI_Start)
	if len(piStartStr) < 3 || !strings.HasPrefix(piStartStr, "0x") {
		ctx.Log().Warn("env PI_Start is invalid, use default 0x6401", log.Any("PI_Start", piStartStr))
		piStartStr = "0x6401"
	}
	doStartStr := os.Getenv(DO_Start)
	if len(doStartStr) < 3 || !strings.HasPrefix(doStartStr, "0x") {
		ctx.Log().Warn("env DO_Start is invalid, use default 0x6001", log.Any("DO_Start", doStartStr))
		doStartStr = "0x6001"
	}
	aoStartStr := os.Getenv(AO_Start)
	if len(aoStartStr) < 3 || !strings.HasPrefix(aoStartStr, "0x") {
		ctx.Log().Warn("env AO_Start is invalid, use default 0x6201", log.Any("AO_Start", aoStartStr))
		aoStartStr = "0x6201"
	}

	diStart, err := strconv.ParseUint(diStartStr[2:], 16, 16)
	if err != nil {
		return nil, err
	}
	aiStart, err := strconv.ParseUint(aiStartStr[2:], 16, 16)
	if err != nil {
		return nil, err
	}
	piStart, err := strconv.ParseUint(piStartStr[2:], 16, 16)
	if err != nil {
		return nil, err
	}
	doStart, err := strconv.ParseUint(doStartStr[2:], 16, 16)
	if err != nil {
		return nil, err
	}
	aoStart, err := strconv.ParseUint(aoStartStr[2:], 16, 16)
	if err != nil {
		return nil, err
	}

	return map[string]uint{
		DI: uint(diStart),
		AI: uint(aiStart),
		PI: uint(piStart),
		DO: uint(doStart),
		AO: uint(aoStart),
	}, nil
}
