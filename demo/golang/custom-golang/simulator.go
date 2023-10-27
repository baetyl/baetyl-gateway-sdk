package custom

import (
	"encoding/json"
	"math/rand"
	"time"
)

const (
	SimulatorInterval = 10 * time.Second
)

// Simulator 自定义设备模拟器
// 每隔 10s 数据变化一次
type Simulator struct {
	name string
	point
}

func NewSimulator(name string) *Simulator {
	s := &Simulator{name: name}
	s.point = point{
		Temperature: float32(rand.Intn(100)),
		Humidity:    float32(rand.Intn(100)),
		Pressure:    float32(rand.Intn(100)),
	}
	go func() {
		t := time.NewTicker(SimulatorInterval)
		defer t.Stop()
		for {
			select {
			case <-t.C:
				s.Temperature = float32(rand.Intn(100))
				s.Humidity = float32(rand.Intn(100))
				s.Pressure = float32(rand.Intn(100))
			}
		}
	}()
	return s
}

func (s *Simulator) Get(name string) (any, error) {
	dt, err := json.Marshal(&s.point)
	if err != nil {
		return nil, err
	}
	res := make(map[string]any)
	err = json.Unmarshal(dt, &res)
	if err != nil {
		return nil, err
	}
	return res[name], nil
}

func (s *Simulator) Set(val float32, index int) {
	switch index {
	case Temperature:
	case Humidity:
	case Pressure:
		s.Pressure = val
	}
}

// point 模拟器有如下点位
// Temperature 温度，只读，Index 0
// Humidity    湿度，只读，Index 1
// Pressure    压力，读写，Index 2
type point struct {
	Temperature float32 `yaml:"temperature" json:"temperature"`
	Humidity    float32 `yaml:"humidity" json:"humidity"`
	Pressure    float32 `yaml:"pressure" json:"pressure"`
}
