package httpin

import (
	"github.com/baetyl/baetyl-gateway-sdk/sdk/golang"
)

type Engine struct {
	cfg    *Config
	ss     []*Server
	report plugin.Report
}

func NewEngine(cfg *Config, report plugin.Report) (*Engine, error) {
	e := &Engine{
		cfg:    cfg,
		report: report,
	}
	L().Debug("NewEngine function")

	var ss []*Server
	for _, item := range cfg.Servers {
		s, err := NewServer(&item)
		if err != nil {
			L().Error("engine new Server error", err.Error())
			continue
		}
		a := NewAPI(cfg.DriverName, report)
		s.SetAPI(a)
		s.InitRoute()
		ss = append(ss, s)
	}

	e.ss = ss
	return e, nil
}

// Start 启动对每个设备的周期性采集上报
func (e *Engine) Start() {
	L().Debug("Engine Start function")
	for _, s := range e.ss {
		go s.Run()
	}
}

func (e *Engine) Restart() {
	L().Debug("Engine Restart function")
	e.Stop()
	e.Start()
}

func (e *Engine) Stop() {
	L().Debug("Engine Stop function")
	if e != nil && e.ss != nil {
		for _, s := range e.ss {
			if s != nil {
				s.Close()
			}
		}
	}
}
