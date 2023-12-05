package bacnet

import (
	"encoding/json"
	"time"

	plugin "github.com/baetyl/baetyl-gateway-sdk/sdk/golang"
	dm "github.com/baetyl/baetyl-go/v2/dmcontext"
	"github.com/baetyl/baetyl-go/v2/errors"
	"github.com/baetyl/baetyl-go/v2/spec/v1"
	"github.com/chen-Leo/bacnet"
	"github.com/chen-Leo/bacnet/bacip"
)

var bacnetClient *bacip.Client
var devices []bacnet.Device

type Slave struct {
	info         *dm.DeviceInfo
	bacnetClient *bacip.Client
	device       bacnet.Device
	cfg          SlaveConfig
	driver       *Driver
	fail         int
	status       int
}

func NewSlave(d *Driver, info *dm.DeviceInfo, cfg SlaveConfig) (*Slave, error) {
	if bacnetClient == nil {
		c, err := bacip.NewClientByIp(cfg.Address, cfg.Port)
		if err != nil {
			return nil, errors.Trace(err)
		}
		bacnetClient = c
	}
	if devices != nil && len(devices) > 0 {
		slave, err := generateSlave(d, info, cfg)
		if err == nil {
			return slave, nil
		}
	}
	var devs []bacnet.Device
	for {
		d.log.Info("Search devices on", "address", cfg.Address)
		var searchErr error
		devs, searchErr = bacnetClient.WhoIs(bacip.WhoIs{}, 5*time.Second)
		if searchErr != nil {
			return nil, errors.Trace(searchErr)
		}
		if devs != nil && len(devs) != 0 {
			break
		}
	}
	d.log.Info("Find devices:", "devices", devs, "cfg", cfg)
	devices = devs
	return generateSlave(d, info, cfg)
}

func generateSlave(d *Driver, info *dm.DeviceInfo, cfg SlaveConfig) (*Slave, error) {
	slave := &Slave{
		info:         info,
		cfg:          cfg,
		bacnetClient: bacnetClient,
		driver:       d,
	}
	for _, device := range devices {
		if uint32(device.ID.Instance) == cfg.DeviceID {
			slave.device = device
			return slave, nil
		}
	}
	return nil, errors.New("device not find")
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
