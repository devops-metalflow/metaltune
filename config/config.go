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
	Bench     string
	Brain     string
	KeenTuned string
	Target    string
}

var (
	Build   string
	Version string
)

func New() *Config {
	return &Config{}
}
