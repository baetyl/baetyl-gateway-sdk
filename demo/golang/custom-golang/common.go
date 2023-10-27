package custom

import (
	"errors"

	"github.com/hashicorp/go-hclog"
)

// point 模拟器有如下点位
// Temperature 温度，只读，Index 0
// Humidity    湿度，只读，Index 1
// Pressure    压力，读写，Index 2
const (
	Temperature = 0
	Humidity    = 1
	Pressure    = 2
)

var (
	ErrWorkerNotExist          = errors.New("worker not exist")
	ErrDriverNameNotExist      = errors.New("failed to get driverName in msg")
	ErrDevNameNotExist         = errors.New("failed to get deviceName in msg")
	ErrEventTypeNotSupported   = errors.New("event type not supported yet")
	ErrMessageTypeNotSupported = errors.New("message type not supported yet")
)

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
