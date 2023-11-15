package httpin

import (
	"time"

	"github.com/baetyl/baetyl-go/v2/utils"
	"github.com/hashicorp/go-hclog"
)

type Config struct {
	DriverName string         `yaml:"drivername" json:"drivername"`
	Servers    []ServerConfig `yaml:"devices" json:"devices"`
}

type ServerConfig struct {
	Port         string            `yaml:"port" json:"port"`
	ReadTimeout  time.Duration     `yaml:"readTimeout" json:"readTimeout" default:"30s"`
	WriteTimeout time.Duration     `yaml:"writeTimeout" json:"writeTimeout" default:"30s"`
	ShutdownTime time.Duration     `yaml:"shutdownTime" json:"shutdownTime" default:"3s"`
	Certificate  utils.Certificate `yaml:",inline" json:",inline"`
}

var (
	_log hclog.Logger
)

func InitL(l hclog.Logger) {
	_log = l
}

func L() hclog.Logger {
	if _log == nil {
		return hclog.L()
	}
	return _log
}
