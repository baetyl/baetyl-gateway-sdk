package iec104

import "time"

type Config struct {
	DriverName string        `yaml:"drivername" json:"drivername"`
	Slaves     []SlaveConfig `yaml:"slaves" json:"slaves"`
	Jobs       []Job         `yaml:"jobs" json:"jobs"`
}

type SlaveConfig struct {
	Device   string        `yaml:"device" json:"device"`
	ID       byte          `yaml:"id" json:"id"`
	Interval time.Duration `yaml:"interval,omitempty" json:"interval,omitempty"`
	Endpoint string        `yaml:"endpoint" json:"endpoint"`
	AIOffset uint16        `yaml:"aiOffset" json:"aiOffset"`
	AOOffset uint16        `yaml:"aoOffset" json:"aoOffset"`
	DIOffset uint16        `yaml:"diOffset" json:"diOffset"`
	DOOffset uint16        `yaml:"doOffset" json:"doOffset"`
}

type Job struct {
	Device   string        `yaml:"device" json:"device"`
	Interval time.Duration `yaml:"interval" json:"interval" default:"900s"`
	Points   []Point       `yaml:"points" json:"points"`
}

type Point struct {
	Name      string `yaml:"name" json:"name"`
	PointNum  uint   `yaml:"pointNum" json:"pointNum"`
	PointType string `yaml:"pointType" json:"pointType"`
	Type      string `yaml:"type" json:"type"`
}

type Publish struct {
	QOS   uint32 `yaml:"qos" json:"qos" validate:"min=0, max=1"`
	Topic string `yaml:"topic" json:"topic" default:"timer" validate:"nonzero"`
}
