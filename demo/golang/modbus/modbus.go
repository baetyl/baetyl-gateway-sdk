package modbus

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"time"

	"github.com/baetyl/baetyl-gateway/config"
	dm "github.com/baetyl/baetyl-go/v2/dmcontext"
	"github.com/baetyl/baetyl-go/v2/errors"
	"github.com/hashicorp/go-hclog"
	"github.com/panjf2000/ants/v2"
)

type Modbus struct {
	ctx     dm.Context
	log     hclog.Logger
	slaves  map[string]*Slave
	ws      map[string]*Worker
	pool    *ants.PoolWithFunc
	tickers []*time.Ticker
}

func NewModbus(d *Driver, ctx dm.Context, cfg *Config) (*Modbus, error) {
	devMap := map[string]dm.DeviceInfo{}
	for _, dev := range ctx.GetAllDevices(cfg.DriverName) {
		devMap[dev.Name] = dev
	}
	slaves := map[string]*Slave{}
	for _, slaveConfig := range cfg.Slaves {
		client, err := NewClient(slaveConfig)
		if err != nil {
			return nil, err
		}
		err = client.Connect()
		if err != nil {
			d.log.Error("connect failed", "slave id", slaveConfig.Id, "error", err)
			continue
		}
		dev, ok := devMap[slaveConfig.Device]
		if !ok {
			d.log.Error("can not find device according to job config", "device", slaveConfig.Device)
			continue
		}
		slave := NewSlave(ctx, d, &dev, slaveConfig, client)
		slaves[slaveConfig.Device] = slave
		if err = slave.UpdateStatus(SlaveOnline); err != nil {
			d.log.Error("failed to update status", "error", err)
		}
	}
	mod := &Modbus{
		ctx:    ctx,
		ws:     make(map[string]*Worker),
		log:    d.log,
		slaves: slaves,
	}
	for _, job := range cfg.Jobs {
		if slave := slaves[job.Device]; slave != nil {
			dev, ok := devMap[slave.cfg.Device]
			if !ok {
				d.log.Error("can not find device according to job config", "device", slave.cfg.Device)
				continue
			}
			mod.ws[dev.Name] = NewWorker(d, ctx, job, slave, d.log)
		} else {
			d.log.Error("slave id of job is invalid", "device", job.Device)
		}
	}
	pool, err := ants.NewPoolWithFunc(DefaultAntsPoolSize, mod.working, ants.WithPreAlloc(true))
	if err != nil {
		d.log.Error("failed to allocate goroutine pool", "error", err)
		return nil, err
	}
	mod.pool = pool
	return mod, nil
}

func (mod *Modbus) Start() {
	for _, worker := range mod.ws {
		err := mod.pool.Invoke(worker)
		if err != nil {
			mod.log.Error("failed to invoke modbus worker into go pool", "error", err)
			continue
		}
	}
}

func (mod *Modbus) Restart() {
	for _, ticker := range mod.tickers {
		ticker.Stop()
	}
	mod.Start()
}

func (mod *Modbus) Stop() {
	for _, ticker := range mod.tickers {
		ticker.Stop()
	}
	for _, slave := range mod.slaves {
		if err := slave.client.Close(); err != nil {
			mod.log.Warn("failed to close slave", "slave id", slave.cfg.Id, "error", err)
		}
	}
}

func (mod *Modbus) Set(_ string, info *dm.DeviceInfo, props []config.DeviceProperty) error {
	var err error
	w, ok := mod.ws[info.Name]
	if !ok {
		mod.log.Warn("worker not exist according to device", "device", info.Name)
		return ErrWorkerNotExist
	}

	ms := map[string]MapConfig{}
	for _, m := range w.job.Maps {
		ms[m.Name] = m
	}
	for _, prop := range props {
		slave, ok := mod.slaves[w.job.Device]
		if !ok {
			mod.log.Warn("did not find slave to write", "device", w.job.Device)
			continue
		}
		if slave.client.Client == nil {
			mod.log.Warn("slave client is nil", "device", w.job.Device)
			continue
		}

		cfg, ok := ms[prop.PropName]
		if !ok {
			mod.log.Warn("did not find prop", "name", prop.PropName)
			continue
		}

		var value any
		switch prop.PropVal.(type) {
		case float64:
			value, err = dm.ParsePropertyValue(cfg.Type, prop.PropVal.(float64))
		default:
			value = prop.PropVal
		}
		if err != nil {
			mod.log.Warn("parse property value err", "error", err)
			continue
		}
		bs, err := transform(value, cfg)
		if err != nil {
			mod.log.Warn("ignore illegal data type of value", "value", value, "type", cfg.Type, "error", err)
			continue
		}
		switch cfg.Function {
		case DiscreteInput:
		case InputRegister:
			return fmt.Errorf("can not write data with illegal function code: [%d]", cfg.Function)
		case Coil:
			if _, err := slave.client.WriteMultipleCoils(cfg.Address, cfg.Quantity, bs); err != nil {
				return err
			}
		case HoldingRegister:
			if _, err := slave.client.WriteMultipleRegisters(cfg.Address, cfg.Quantity, bs); err != nil {
				return err
			}
		}
	}
	return nil
}

func transform(value any, cfg MapConfig) ([]byte, error) {
	var order binary.ByteOrder = binary.BigEndian
	if cfg.Function == HoldingRegister && cfg.SwapByte {
		order = binary.LittleEndian
	}
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, order, value)
	if err != nil {
		return nil, errors.Trace(err)
	}
	bs := buf.Bytes()
	if cfg.Function == HoldingRegister && cfg.SwapRegister {
		for i := 0; i < len(bs)-1; i += 2 {
			bs[i], bs[i+1] = bs[i+1], bs[i]
		}
	}
	return bs, nil
}

func (mod *Modbus) Event(info *dm.DeviceInfo, event *dm.Event) error {
	switch event.Type {
	case dm.TypeReportEvent:
		return mod.PropertyGet(info, nil)
	default:
		return errors.New("event type not supported yet")
	}
}

func (mod *Modbus) PropertyGet(info *dm.DeviceInfo, _ []string) error {
	w, ok := mod.ws[info.Name]
	if !ok {
		mod.log.Warn("worker not exist according to device", "device", info.Name)
		return ErrWorkerNotExist
	}
	if err := w.Execute(false); err != nil {
		return err
	}
	return nil
}

func (mod *Modbus) working(w interface{}) {
	worker, ok := w.(*Worker)
	if !ok {
		return
	}
	ticker := time.NewTicker(worker.job.Interval)
	mod.tickers = append(mod.tickers, ticker)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			err := worker.Execute(true)
			if err != nil {
				mod.log.Error("failed to execute job", "error", err)
			}
		}
	}
}
