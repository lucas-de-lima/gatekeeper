package config

import (
	"github.com/spf13/viper"
)

// Config is the application configuration root.
type Config struct {
	Version     string `mapstructure:"VERSION"`
	Environment string `mapstructure:"ENVIRONMENT"`
	Server      ServerConfig
	Log         LogConfig
	Database    DatabaseConfig
}

// ServerConfig application server settings.
type ServerConfig struct {
	Port int `mapstructure:"SERVER_PORT"`
}

// LogConfig logging settings.
type LogConfig struct {
	Level string `mapstructure:"LOG_LEVEL"`
}

// DatabaseConfig database settings.
type DatabaseConfig struct {
	URL string `mapstructure:"DATABASE_URL"`
}

// Load loads configuration from environment variables and .env file.
func Load() (*Config, error) {
	v := viper.New()
	v.SetConfigFile(".env")
	v.AutomaticEnv()

	v.SetDefault("VERSION", "0.1.0")
	v.SetDefault("ENVIRONMENT", "development")
	v.SetDefault("SERVER_PORT", 8080)
	v.SetDefault("LOG_LEVEL", "info")
	v.SetDefault("DATABASE_URL", "postgres://gatekeeper:gatekeeper@localhost:5432/gatekeeper?sslmode=disable")

	_ = v.ReadInConfig() // ignore if no file

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
