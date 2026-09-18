package api

import "github.com/kelseyhightower/envconfig"

type Config struct {
	AppName    string `default:"bookmark_service" envconfig:"APP_NAME"`
	AppPort    string `default:"8080" envconfig:"APP_PORT"`
	InstanceID string ``
	LogLevel   string `default:"info" envconfig:"LOG_LEVEL"`
}

func NewConfig(envPrefix string) (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process(envPrefix, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
