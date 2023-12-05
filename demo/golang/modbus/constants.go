package modbus

import "github.com/baetyl/baetyl-go/v2/errors"

var (
	ErrClientInvalid        = errors.New("device client is invalid")
	ErrUnsupportedFieldType = errors.New("unsupported field type")

	ErrWorkerNotExist     = errors.New("worker not exist")
	ErrDriverNameNotExist = errors.New("failed to get driverName in msg")
	ErrDevNameNotExist    = errors.New("failed to get deviceName in msg")
	ErrGetDeviceProps     = errors.New("failed to get device props in msg")
)

const (
	SlaveOffline = 0
	SlaveOnline  = 1

	DefaultAntsPoolSize = 1000

	ModeTCP Mode = "tcp"
	ModeRTU Mode = "rtu"

	Coil            = 1
	DiscreteInput   = 2
	HoldingRegister = 3
	InputRegister   = 4
	SlaveId         = "slaveid"
	SysTime         = "time"

	Bool    = "bool"
	Int16   = "int16"
	UInt16  = "uint16"
	Int32   = "int32"
	UInt32  = "uint32"
	Int64   = "int64"
	UInt64  = "uint64"
	Float32 = "float32"
	Float64 = "float64"
)

var SysType = map[string]struct{}{
	Bool:    {},
	Int16:   {},
	UInt16:  {},
	Int32:   {},
	UInt32:  {},
	Int64:   {},
	UInt64:  {},
	Float32: {},
	Float64: {},
}
var SysName = map[string]struct{}{
	SysTime: {},
	SlaveId: {},
}
