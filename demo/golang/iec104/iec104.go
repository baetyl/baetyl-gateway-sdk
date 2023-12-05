package iec104

import (
	"time"

	"github.com/baetyl/baetyl-gateway/config"
	dm "github.com/baetyl/baetyl-go/v2/dmcontext"
	"github.com/baetyl/baetyl-go/v2/errors"
	"github.com/hashicorp/go-hclog"
	"github.com/panjf2000/ants/v2"
	"github.com/spf13/cast"
	"github.com/thinkgos/go-iecp5/cs104"
)

type IEC104 struct {
	ctx     dm.Context
	ws      map[string]*Worker
	pool    *ants.PoolWithFunc
	tickers []*time.Ticker
	log     hclog.Logger
}

func NewIEC104(d *Driver, ctx dm.Context, cfg *Config) (*IEC104, error) {
	infos := make(map[string]dm.DeviceInfo)
	for _, info := range ctx.GetAllDevices(cfg.DriverName) {
		infos[info.Name] = info
	}
	slaves := make(map[string]*Slave)
	for _, dCfg := range cfg.Slaves {
		if info, ok := infos[dCfg.Device]; ok {
			// slave异步初始化
			slave, err := NewSlave(d, &info, dCfg)
			if err != nil {
				d.log.Error("ignore device which failed to establish connection", "device", dCfg.Device, "error", err)
				continue
			}
			slaves[dCfg.Device] = slave
			flag := false
			// 监听slave初始化成功
			slave.client.SetOnConnectHandler(func(c *cs104.Client) {
				if flag {
					return
				}
				flag = true
				slaves[dCfg.Device] = slave
				if err = slave.UpdateStatus(SlaveOnline); err != nil {
					d.log.Error("failed to update status", "error", err)
				}
			})
		}
	}

	ws := make(map[string]*Worker)
	for _, job := range cfg.Jobs {
		if slave := slaves[job.Device]; slave != nil {
			ws[slave.info.Name] = NewWorker(d, ctx, job, slave, d.log)
		} else {
			d.log.Error("device of job not exist", "device id", job.Device)
		}
	}
	iec := &IEC104{
		ctx: ctx,
		log: d.log,
		ws:  ws,
	}
	pool, err := ants.NewPoolWithFunc(DefaultAntsPoolSize, iec.working, ants.WithPreAlloc(true))
	if err != nil {
		d.log.Error("failed to allocate goroutine pool", "error", err)
		return nil, err
	}
	iec.pool = pool
	return iec, nil
}

func (i *IEC104) Start() {
	for _, worker := range i.ws {
		err := i.pool.Invoke(worker)
		if err != nil {
			i.log.Error("failed to invoke iec104 worker into go pool", "error", err)
			continue
		}
	}
}

func (i *IEC104) Restart() {
	i.Stop()
	i.Start()
}

func (i *IEC104) Stop() {
	for _, ticker := range i.tickers {
		ticker.Stop()
	}
}

func (i *IEC104) working(w interface{}) {
	worker, ok := w.(*Worker)
	if !ok {
		return
	}
	ticker := time.NewTicker(worker.job.Interval)
	i.tickers = append(i.tickers, ticker)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			err := worker.Execute()
			if err != nil {
				i.log.Error("failed to execute job", "error", err)
			}
		}
	}
}

func (i *IEC104) Set(_ string, info *dm.DeviceInfo, delta []config.DeviceProperty) error {
	var err error
	w, ok := i.ws[info.Name]
	if !ok {
		i.log.Warn("worker not exist according to device", "device", info.Name)
		return ErrWorkerNotExist
	}
	// 遍历delta，解析出对应的采集点，进行置数设置
	for _, val := range delta {
		// 对采集点进行置数操作。根据采集点类型AO、DO选择不同命令。
		for _, point := range w.job.Points {
			if val.PropName == point.Name {
				var value interface{}
				switch val.PropVal.(type) {
				case float64:
					value, err = dm.ParsePropertyValue(point.Type, val.PropVal.(float64))
				default:
					value = val.PropVal
				}
				if err != nil {
					return errors.Trace(err)
				}
				if point.PointType == "DO" {
					// 单点遥控
					parseVal, parseErr := cast.ToBoolE(value)
					if parseErr == nil {
						err = w.SingleCmd(point, parseVal)
					} else {
						err = parseErr
					}
				} else if point.PointType == "AO" {
					// 遥调设点
					parseVal, parseErr := cast.ToFloat32E(value)
					if parseErr == nil {
						err = w.SetPointCmdFloat(point, parseVal)
					} else {
						err = parseErr
					}
				}
				if err != nil {
					return errors.Trace(err)
				}
			}
		}
	}
	return nil
}

func (i *IEC104) Event(info *dm.DeviceInfo, event *dm.Event) error {
	return nil
}

func (i *IEC104) PropertyGet(info *dm.DeviceInfo, _ []string) error {
	w, ok := i.ws[info.Name]
	if !ok {
		i.log.Warn("worker not exist according to device", "device", info.Name)
		return ErrWorkerNotExist
	}
	if err := w.Execute(); err != nil {
		return err
	}
	return nil
}
