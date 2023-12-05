package iec104

import (
	"encoding/json"

	plugin "github.com/baetyl/baetyl-gateway-sdk/sdk/golang"
	dm "github.com/baetyl/baetyl-go/v2/dmcontext"
	"github.com/baetyl/baetyl-go/v2/spec/v1"
	"github.com/thinkgos/go-iecp5/cs104"
)

type Slave struct {
	info    *dm.DeviceInfo
	client  *cs104.Client
	cfg     SlaveConfig
	handler *Handler
	driver  *Driver
	fail    int
	status  int
}

func NewSlave(d *Driver, info *dm.DeviceInfo, cfg SlaveConfig) (*Slave, error) {
	// TODO 自动重连、重连时间间隔等配置目前为默认，后续可以按需抽取到配置文件中
	option := cs104.NewOption()
	if err := option.AddRemoteServer(cfg.Endpoint); err != nil {
		return nil, err
	}
	handler := &Handler{}
	client := cs104.NewClient(handler, option)
	// 异步打开tcp连接，监听通道apdu消息。
	err := client.Start()
	if err != nil {
		return nil, err
	}
	return &Slave{
		client:  client,
		cfg:     cfg,
		info:    info,
		handler: handler,
		driver:  d,
	}, nil
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
	msg.Metadata[dm.KeyDeviceName] = s.info.Name
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
	msg.Metadata[dm.KeyDeviceName] = s.info.Name
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
