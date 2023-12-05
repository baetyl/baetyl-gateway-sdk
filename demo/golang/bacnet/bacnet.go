package bacnet

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/baetyl/baetyl-gateway/config"
	dm "github.com/baetyl/baetyl-go/v2/dmcontext"
	"github.com/baetyl/baetyl-go/v2/errors"
	"github.com/chen-Leo/bacnet"
	"github.com/chen-Leo/bacnet/bacip"
	"github.com/hashicorp/go-hclog"
	"github.com/panjf2000/ants/v2"
)

type Bacnet struct {
	ctx     dm.Context
	log     hclog.Logger
	ws      map[string]*Worker
	pool    *ants.PoolWithFunc
	tickers []*time.Ticker
}

func NewBacnet(d *Driver, ctx dm.Context, cfg *Config) (*Bacnet, error) {
	infos := make(map[string]dm.DeviceInfo)
	for _, info := range ctx.GetAllDevices(d.driverName) {
		infos[info.Name] = info
	}

	slaves := make(map[string]*Slave)
	for _, dCfg := range cfg.Slaves {
		if info, ok := infos[dCfg.Device]; ok {
			slave, err := NewSlave(d, &info, dCfg)
			if err != nil {
				d.log.Error("failed to create device instance", "device", dCfg.Device, "error", err)
				continue
			}
			slaves[dCfg.Device] = slave
			if err = slave.UpdateStatus(SlaveOnline); err != nil {
				d.log.Error("failed to update status", "error", err)
			}
		}
	}
	ws := make(map[string]*Worker)
	for _, job := range cfg.Jobs {
		if dev := slaves[job.Device]; dev != nil {
			ws[job.Device] = NewWorker(d, ctx, job, dev, d.log)
		} else {
			d.log.Error("device of job not exist", "device id", job.Device)
		}
	}

	bac := &Bacnet{
		ctx: ctx,
		log: d.log,
		ws:  ws,
	}
	pool, err := ants.NewPoolWithFunc(DefaultAntsPoolSize, bac.working, ants.WithPreAlloc(true))
	if err != nil {
		d.log.Error("failed to allocate goroutine pool", "error", err)
		return nil, err
	}
	bac.pool = pool
	return bac, nil
}

func (bac *Bacnet) Start() {
	for _, worker := range bac.ws {
		err := bac.pool.Invoke(worker)
		if err != nil {
			bac.log.Error("failed to invoke bacnet worker into go pool", "error", err)
			continue
		}
	}
}

func (bac *Bacnet) Restart() {
	for _, ticker := range bac.tickers {
		ticker.Stop()
	}
	bac.Start()
}

func (bac *Bacnet) Stop() {
	for _, ticker := range bac.tickers {
		ticker.Stop()
	}
}

func (bac *Bacnet) Set(_ string, info *dm.DeviceInfo, props []config.DeviceProperty) error {
	var err error
	w, ok := bac.ws[info.Name]
	if !ok {
		bac.log.Warn("worker not exist according to device", "device", info.Name)
		return ErrWorkerNotExist
	}

	for _, p := range props {
		for _, prop := range w.job.Properties {
			if p.PropName == prop.Name {
				var value interface{}
				switch bacnet.PropertyValueType(prop.ApplicationTagNumber) {
				case bacnet.TypeBoolean:
					value, err = dm.ParseValueToBool(p.PropVal)
					if err != nil {
						return err
					}
				case bacnet.TypeReal:
					value, err = dm.ParseValueToFloat32(p.PropVal)
					if err != nil {
						return err
					}
				case bacnet.TypeEnumerated:
					value, err = dm.ParseValueToUint32(p.PropVal)
					if err != nil {
						return err
					}
				default:
					return errors.New(fmt.Sprintf("unsupported type conversion.prop name: %v, prop type: %v, application tag num: %v",
						p.PropName, reflect.TypeOf(p.PropVal).Name(), prop.ApplicationTagNumber))
				}

				objID := bacnet.ObjectID{
					Type:     bacnet.ObjectType(prop.BacnetType),
					Instance: bacnet.ObjectInstance(prop.BacnetAddress),
				}
				err = writeValue(w.slave.bacnetClient, w.slave.device, objID, bacnet.PropertyValue{
					Type:  bacnet.PropertyValueType(prop.ApplicationTagNumber),
					Value: value,
				}, bacnet.PropertyIdentifier{
					Type: bacnet.PresentValue,
				})
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (bac *Bacnet) Event(info *dm.DeviceInfo, event *dm.Event) error {
	switch event.Type {
	case dm.TypeReportEvent:
		return bac.PropertyGet(info, nil)
	default:
		return errors.New("event type not supported yet")
	}
}

func (bac *Bacnet) PropertyGet(info *dm.DeviceInfo, _ []string) error {
	w, ok := bac.ws[info.Name]
	if !ok {
		bac.log.Warn("worker not exist according to device", "device", info.Name)
		return ErrWorkerNotExist
	}
	if err := w.Execute(false); err != nil {
		return err
	}
	return nil
}

func (bac *Bacnet) working(w interface{}) {
	worker, ok := w.(*Worker)
	if !ok {
		return
	}
	ticker := time.NewTicker(worker.job.Interval)
	bac.tickers = append(bac.tickers, ticker)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			err := worker.Execute(true)
			if err != nil {
				bac.log.Error("failed to execute job", "error", err)
			}
		}
	}
}

func writeValue(c *bacip.Client, device bacnet.Device, object bacnet.ObjectID, propertyValue bacnet.PropertyValue,
	property bacnet.PropertyIdentifier) error {
	wp := bacip.WriteProperty{
		ObjectID:      object,
		Property:      property,
		PropertyValue: propertyValue,
		Priority:      bacnet.ManualLifeSafety1,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := c.WriteProperty(ctx, device, wp)
	if err != nil {
		fmt.Printf("%v\t", err)
		return err
	}
	return nil
}
