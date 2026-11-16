package api

import "github.com/kelseyhightower/envconfig"

type Config struct {
	AppName string `default:"bookmark_service" envconfig:"APP_NAME"`
	AppPort int    `default:"8081" envconfig:"APP_PORT"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process("", cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
