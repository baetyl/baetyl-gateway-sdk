package modbus

import (
	"encoding/json"

	plugin "github.com/baetyl/baetyl-gateway-sdk/sdk/golang"
	dm "github.com/baetyl/baetyl-go/v2/dmcontext"
	"github.com/baetyl/baetyl-go/v2/spec/v1"
)

type Slave struct {
	client *mbClient
	dev    *dm.DeviceInfo
	ctx    dm.Context
	cfg    SlaveConfig
	driver *Driver
	fail   int
	status int
}

func NewSlave(ctx dm.Context, d *Driver, dev *dm.DeviceInfo, cfg SlaveConfig, client *mbClient) *Slave {
	return &Slave{
		status: SlaveOffline,
		ctx:    ctx,
		dev:    dev,
		client: client,
		cfg:    cfg,
		driver: d,
	}
}

func (s *Slave) UpdateStatus(status int) error {
	if status == s.status {
		return nil
	}
	if status == SlaveOffline {
		s.fail++
		if s.fail == 3 {
			err := s.Offline()
			if err != nil {
				return err
			}
			s.status = SlaveOffline
			s.fail = 0
		}
	} else if status == SlaveOnline {
		err := s.Online()
		if err != nil {
			return err
		}
		s.status = SlaveOnline
	}
	return nil
}

func (s *Slave) Online() error {
	msg := &v1.Message{}
	msg.Kind = v1.MessageDeviceLifecycleReport
	msg.Metadata = make(map[string]string)
	msg.Metadata[dm.KeyDriverName] = s.driver.driverName
	msg.Metadata[dm.KeyDeviceName] = s.dev.Name
	msg.Content = v1.LazyValue{Value: true}

	msgData, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	_, err = s.driver.report.State(&plugin.Request{Req: string(msgData)})
	if err != nil {
		return err
	}
	return nil
}

func (s *Slave) Offline() error {
	msg := &v1.Message{}
	msg.Kind = v1.MessageDeviceLifecycleReport
	msg.Metadata = make(map[string]string)
	msg.Metadata[dm.KeyDriverName] = s.driver.driverName
	msg.Metadata[dm.KeyDeviceName] = s.dev.Name
	msg.Content = v1.LazyValue{Value: false}

	msgData, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	_, err = s.driver.report.State(&plugin.Request{Req: string(msgData)})
	if err != nil {
		return err
	}
	return nil
}
