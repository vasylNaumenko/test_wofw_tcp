package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// MustLoad reads the configuration and returns it.
func MustLoad(cfgPath string) *Config {
	cfg, err := loadConfigFromFile(cfgPath)
	if err != nil {
		panic(fmt.Sprintf("config validation failed: %s", err))
	}

	return cfg
}

// loadConfigFromFile reads the configuration from the file.
func loadConfigFromFile(filepath string) (*Config, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return &cfg, nil
}
