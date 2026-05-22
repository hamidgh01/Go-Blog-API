package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConf
	Postgres PostgresConf
}

type ServerConf struct {
	Host string
	Port int
}
type PostgresConf struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	// SSLMode         string
	// MaxIdleConns    int
	// MaxOpenConns    int
	// ConnMaxLifetime int
	// TimeZone        string
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
