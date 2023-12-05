package bacnet

import (
	"github.com/baetyl/baetyl-go/v2/errors"
)

var (
	ErrWorkerNotExist     = errors.New("worker not exist")
	ErrDriverNameNotExist = errors.New("failed to get driverName in msg")
	ErrDevNameNotExist    = errors.New("failed to get deviceName in msg")
	ErrGetDeviceProps     = errors.New("failed to get device props in msg")
)

const (
	SlaveOffline        = 0
	SlaveOnline         = 1
	DefaultAntsPoolSize = 1000
)
