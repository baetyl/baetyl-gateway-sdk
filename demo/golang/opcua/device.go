package opcua

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"

	plugin "github.com/baetyl/baetyl-gateway-sdk/sdk/golang"
	dm "github.com/baetyl/baetyl-go/v2/dmcontext"
	"github.com/baetyl/baetyl-go/v2/errors"
	"github.com/baetyl/baetyl-go/v2/spec/v1"
	"github.com/gopcua/opcua"
	"github.com/gopcua/opcua/ua"
)

type Device struct {
	info        *dm.DeviceInfo
	opcuaClient *opcua.Client
	cfg         DeviceConfig
	driver      *Driver
	status      int
	fail        int
}

func NewDevice(d *Driver, info *dm.DeviceInfo, cfg DeviceConfig) (*Device, error) {
	opts := []opcua.Option{
		opcua.RequestTimeout(cfg.Timeout),
		opcua.SecurityPolicy(cfg.Security.Policy),
		opcua.SecurityModeString(cfg.Security.Mode),
	}

	cli := opcua.NewClient(cfg.Endpoint)
	var ctx, cancel = context.WithTimeout(context.Background(), cfg.Timeout)
	defer cancel()
	if err := cli.Dial(ctx); err != nil {
		return nil, err
	}
	defer cli.Close()
	var res, err = cli.GetEndpoints()
	if err != nil {
		return nil, err
	}
	var ep = opcua.SelectEndpoint(res.Endpoints, cfg.Security.Policy, ua.MessageSecurityModeFromString(cfg.Security.Mode))
	if ep == nil {
		return nil, nil
	}

	if cfg.Auth != nil && cfg.Auth.Username != "" && cfg.Auth.Password != "" {
		opts = append(opts,
			opcua.AuthUsername(cfg.Auth.Username, cfg.Auth.Password),
			opcua.SecurityFromEndpoint(ep, ua.UserTokenTypeUserName),
		)
	} else if cfg.Certificate != nil && cfg.Certificate.Cert != "" && cfg.Certificate.Key != "" {
		cert, err := decodeCert([]byte(cfg.Certificate.Cert))
		if err != nil {
			return nil, err
		}
		key, err := decodeKey([]byte(cfg.Certificate.Key))
		if err != nil {
			return nil, err
		}
		opts = append(opts, opcua.AuthCertificate(cert),
			opcua.Certificate(cert),
			opcua.PrivateKey(key),
			opcua.SecurityFromEndpoint(ep, ua.UserTokenTypeCertificate))
	} else {
		opts = append(opts,
			opcua.AuthAnonymous(),
			opcua.SecurityFromEndpoint(ep, ua.UserTokenTypeAnonymous),
		)
	}

	// optimize timeout
	ctx, cancel = context.WithTimeout(context.Background(), cfg.Timeout)
	defer cancel()
	var client = opcua.NewClient(cfg.Endpoint, opts...)
	if err = client.Connect(ctx); err != nil {
		return nil, errors.Trace(err)
	}
	return &Device{driver: d, info: info, cfg: cfg, opcuaClient: client}, nil
}

func decodeCert(certPEM []byte) ([]byte, error) {
	block, _ := pem.Decode(certPEM)
	if block == nil || block.Type != "CERTIFICATE" {
		return nil, errors.Errorf("failed to decode cert")
	}
	return block.Bytes, nil
}

func decodeKey(keyPEM []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(keyPEM)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, errors.Errorf("failed to decode key")
	}
	var pk, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, errors.Errorf("failed to parse key")
	}
	return pk, nil
}

func (d *Device) UpdateStatus(status int) error {
	if status == d.status {
		return nil
	}
	if status == DeviceOffline {
		d.fail++
		if d.fail == 3 {
			err := d.Offline()
			if err != nil {
				return err
			}
			d.status = DeviceOffline
			d.fail = 0
		}
	} else if status == DeviceOnline {
		err := d.Online()
		if err != nil {
			return err
		}
		d.status = DeviceOnline
	}
	return nil
}

func (d *Device) Online() error {
	msg := &v1.Message{}
	msg.Kind = v1.MessageDeviceLifecycleReport
	msg.Metadata = make(map[string]string)
	msg.Metadata[dm.KeyDriverName] = d.driver.driverName
	msg.Metadata[dm.KeyDeviceName] = d.info.Name
	msg.Content = v1.LazyValue{Value: true}

	msgData, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	_, err = d.driver.report.State(&plugin.Request{Req: string(msgData)})
	if err != nil {
		return err
	}
	return nil
}

func (d *Device) Offline() error {
	msg := &v1.Message{}
	msg.Kind = v1.MessageDeviceLifecycleReport
	msg.Metadata = make(map[string]string)
	msg.Metadata[dm.KeyDriverName] = d.driver.driverName
	msg.Metadata[dm.KeyDeviceName] = d.info.Name
	msg.Content = v1.LazyValue{Value: false}

	msgData, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	_, err = d.driver.report.State(&plugin.Request{Req: string(msgData)})
	if err != nil {
		return err
	}
	return nil
}
