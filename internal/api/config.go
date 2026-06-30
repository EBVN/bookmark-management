package api

import "github.com/kelseyhightower/envconfig"

// Config struct for app config
type Config struct {
	AppName string `envconfig:"APP_NAME" default:"bookmark_service"`
	AppPort string `envconfig:"APP_PORT" default:"8080"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process("", cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
