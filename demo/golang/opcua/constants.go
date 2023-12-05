package opcua

import (
	"github.com/baetyl/baetyl-go/v2/errors"
	"github.com/gopcua/opcua/ua"
)

var (
	ErrWorkerNotExist        = errors.New("worker not exist")
	ErrDriverNameNotExist    = errors.New("failed to get driverName in msg")
	ErrDevNameNotExist       = errors.New("failed to get deviceName in msg")
	ErrEventTypeNotSupported = errors.New("event type not supported yet")
)

const (
	DeviceOffline = 0
	DeviceOnline  = 1

	DefaultAntsPoolSize = 1000

	OpcuaIdTypeI = "i"
	OpcuaIdTypeS = "s"
	OpcuaIdTypeG = "g"
	OpcuaIdTypeB = "b"

	Bool    = "bool"
	Int16   = "int16"
	UInt16  = "uint16"
	Int32   = "int32"
	UInt32  = "uint32"
	Int64   = "int64"
	UInt64  = "uint64"
	Float32 = "float32"
	Float64 = "float64"
	String  = "string"
)

var types = map[string]ua.TypeID{
	Bool:    ua.TypeIDBoolean,
	Int16:   ua.TypeIDInt16,
	UInt16:  ua.TypeIDUint16,
	Int32:   ua.TypeIDInt32,
	UInt32:  ua.TypeIDUint32,
	Int64:   ua.TypeIDInt64,
	UInt64:  ua.TypeIDUint64,
	Float32: ua.TypeIDFloat,
	Float64: ua.TypeIDDouble,
	String:  ua.TypeIDString,
}

func variant2Value(fieldType string, val *ua.Variant) (interface{}, error) {
	if types[fieldType] != val.Type() {
		return nil, errors.Errorf("property type error")
	}
	switch val.Type() {
	case ua.TypeIDBoolean:
		return val.Bool(), nil
	case ua.TypeIDFloat:
		return val.Float(), nil
	case ua.TypeIDDouble:
		return val.Float(), nil
	case ua.TypeIDInt16, ua.TypeIDInt32, ua.TypeIDInt64:
		return val.Int(), nil
	case ua.TypeIDUint16, ua.TypeIDUint32, ua.TypeIDUint64:
		return val.Uint(), nil
	case ua.TypeIDString:
		return val.String(), nil
	default:
		return nil, errors.Errorf("unsupported type")
	}
}
