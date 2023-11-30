package httpin

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/baetyl/baetyl-gateway-sdk/sdk/golang"
	dm "github.com/baetyl/baetyl-go/v2/dmcontext"
	v1 "github.com/baetyl/baetyl-go/v2/spec/v1"
)

type Engine struct {
	cfg    *Config
	server *Server
	report plugin.Report
}

func NewEngine(cfg *Config, report plugin.Report) (*Engine, error) {
	e := &Engine{
		cfg:    cfg,
		report: report,
	}
	L().Debug("NewEngine function")

	var s *Server
	s, err := NewServer(cfg)
	if err != nil {
		L().Error("engine new Server error", err.Error())
		return nil, err
	}
	a := NewAPI(cfg.DriverName, report)
	s.SetAPI(a)
	s.InitRoute()
	e.server = s
	return e, nil
}

// Start 启动对每个设备的周期性采集上报
func (e *Engine) Start() {
	L().Debug("Engine Start function")
	go e.server.Run()
}

func (e *Engine) Restart() {
	L().Debug("Engine Restart function")
	e.Stop()
	e.Start()
}

func (e *Engine) Stop() {
	L().Debug("Engine Stop function")
	if e != nil && e.server != nil {
		e.server.Close()
	}
}

func (e *Engine) Event(info *dm.DeviceInfo, event *dm.Event) error {
	return ErrEventTypeNotSupported
}

func (e *Engine) PropertyGet(info *dm.DeviceInfo, _ []string) error {
	res, err := http.Get(e.cfg.InitAddress + "/v1/devices/" + info.Name + "/propertyget")
	if err != nil {
		return err
	}
	if res.StatusCode >= 300 || res.StatusCode < 200 {
		return errors.New(res.Status)
	}
	var props map[string]any
	err = json.NewDecoder(res.Body).Decode(&props)
	if err != nil {
		return err
	}
	msg := v1.Message{
		Kind: v1.MessageDeviceDesire,
		Metadata: map[string]string{
			dm.KeyDriverName: e.cfg.DriverName,
			dm.KeyDeviceName: info.Name,
		},
		Content: v1.LazyValue{Value: props},
	}

	dt, err := json.Marshal(msg)
	if err != nil {
		return E(ErrRunning, F("error", err.Error()))
	}

	_, err = e.report.Post(&plugin.Request{Req: string(dt)})
	if err != nil {
		return E(ErrRunning, F("error", err.Error()))
	}
	return nil
}
