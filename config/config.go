package config

type Config struct {
	ApiVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	MetaData   MetaData `yaml:"metadata"`
	Spec       Spec     `yaml:"spec"`
}

type MetaData struct {
	Name string `yaml:"name"`
}

type Spec struct {
	KeenTune KeenTune `yaml:"keentune"`
}

type KeenTune struct {
	Url string `yaml:"Url"`
}

var (
	Build   string
	Version string
)

func New() *Config {
	return &Config{}
}
