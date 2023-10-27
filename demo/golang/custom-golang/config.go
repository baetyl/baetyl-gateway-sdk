package custom

import "time"

type Config struct {
	DriverName string         `yaml:"drivername" json:"drivername"`
	Devices    []DeviceConfig `yaml:"devices" json:"devices"`
	Jobs       []Job          `yaml:"jobs" json:"jobs"`
}

// DeviceConfig 连接一个设备所需的具体配置，一般包含设备标识、接入地址、证书、超时时间等，用于生成一个连接设备的 client
// 此处为一个模拟设备，仅包含:
// * Device : 标识符信息
type DeviceConfig struct {
	Device string `yaml:"device" json:"device"`
}

// Property 待采集点位信息，包含点位标识符、数据类型、点位采集的必要请求信息等
// * Name : 点位名称
// * Type : 点位数据类型
// * Index : 单位在设备中的唯一标识
type Property struct {
	Name  string `yaml:"name" json:"name"`
	Type  string `yaml:"type" json:"type"`
	Index int    `yaml:"index" json:"index"`
}

// Job 针对一个设备的采集配置，包含
// * Device : 设备标识，要与某个 DeviceConfig.Device 对应，否则会获取不到设备连接信息，进而无法完成采集
// * Interval : 采集频率
// * Properties : 待采集点位列表
type Job struct {
	Device     string        `yaml:"device" json:"device"`
	Interval   time.Duration `yaml:"interval" json:"interval" default:"20s"`
	Properties []Property    `yaml:"properties" json:"properties"`
}
