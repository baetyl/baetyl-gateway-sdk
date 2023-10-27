package golang

type DriverConfig struct {
	Name       string `yaml:"name" json:"name"`
	BinPath    string `yaml:"binPath" json:"binPath"`
	ConfigPath string `yaml:"configPath" json:"configPath"`
}

type Entry struct {
	Entry string `yaml:"entry" json:"entry" validate:"nonzero"`
}
