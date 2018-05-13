package config

// Config : A type that handles the configuration of the app
type Config struct {
	Clair struct {
		Endpoint string
	}

  Log struct {
		Level  string
		Format string
	}
}

var config Config

func init() {
	config = Config{}
}

func Get() (*Config) {
	return &config
}
