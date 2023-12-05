package modbus

import "time"

type Config struct {
	DriverName string `yaml:"drivername" json:"drivername"`
	// Slaves slave list
	Slaves []SlaveConfig `yaml:"slaves" json:"slaves"`
	// Jobs job list
	Jobs []Job `yaml:"jobs" json:"jobs" validate:"validjobs"`
}

type Job struct {
	Device string `yaml:"device" json:"device"`
	// Interval the interval between task execution
	Interval time.Duration `yaml:"interval" json:"interval" default:"5s"`
	// Maps definition of data points
	Maps []MapConfig `yaml:"maps" json:"maps"`
}

// SlaveConfig modbus slave device configuration
type SlaveConfig struct {
	Device string `yaml:"device" json:"device"`
	// Id slave id
	Id byte `yaml:"id" json:"id"`
	// Mode mode of connecting
	Mode string `yaml:"mode" json:"mode" default:"rtu" validate:"regexp=^(tcp|rtu)?$"`
	// Address Device path (/dev/ttyS0)
	Address string `yaml:"address" json:"address" default:"/dev/ttyS0"`
	// Timeout Read (Write) timeout.
	Timeout time.Duration `yaml:"timeout" json:"timeout" default:"10s"`
	// IdleTimeout Idle timeout to close the connection
	IdleTimeout time.Duration `yaml:"idletimeout" json:"idletimeout" default:"1m"`
	//// RTU only
	// BaudRate (default 19200)
	BaudRate int `yaml:"baudrate" json:"baudrate" default:"19200"`
	// DataBits: 5, 6, 7 or 8 (default 8)
	DataBits int `yaml:"databits" json:"databits" default:"8" validate:"min=5, max=8"`
	// StopBits: 1 or 2 (default 1)
	StopBits int `yaml:"stopbits" json:"stopbits" default:"1" validate:"min=1, max=2"`
	// Parity: N - None, E - Even, O - Odd (default E)
	// (The use of no parity requires 2 stop bits.)
	Parity string `yaml:"parity" json:"parity" default:"E" validate:"regexp=^(E|N|O)?$"`
}

// MapConfig map point configuration
type MapConfig struct {
	Id string `yaml:"id" json:"id"`
	// Name name of map config
	Name string `yaml:"name" json:"name"`
	// Type type of map type
	Type string `yaml:"type" json:"type"`
	// Function
	Function byte `yaml:"function" json:"function" validate:"min=1, max=4" validate:"nonzero"`
	// Address
	Address uint16 `yaml:"address" json:"address"`
	// Quantity
	Quantity uint16 `yaml:"quantity" json:"quantity"`
	// SwapByte whether swap byte, meaning using big endian or little endian
	SwapByte bool `yaml:"swapByte" json:"swapByte"`
	// SwapRegister whether swap high and low register
	SwapRegister bool `yaml:"swapRegister" json:"swapRegister"`
}
