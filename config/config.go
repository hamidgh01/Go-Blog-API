package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConf
}

type ServerConf struct {
	Host string
	Port int
}

var cfg *Config

func Load() (*Config, error) {

	viper.SetConfigFile("./config.yml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file. origin: %w", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to decode configurations to `Config` struct. origin: %w", err)
	}

	return cfg, nil
}
