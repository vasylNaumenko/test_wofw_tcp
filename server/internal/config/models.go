package config

import "time"

type (
	// Config defines the properties of the application configuration.
	Config struct {
		Delivery Delivery `yaml:"delivery"`
		Cites    Cites    `yaml:"cites"`
	}

	// Delivery defines API server configuration.
	Delivery struct {
		GRPCServer     GRPCServer     `yaml:"grpc-server"`
		DDOSProtection DDOSProtection `yaml:"ddos-protection"`
	}

	// GRPCServer defines gRPC server configuration.
	GRPCServer struct {
		Port    string        `yaml:"port"`
		Timeout time.Duration `yaml:"timeout"`
	}

	// Cites configures the CITES API.
	Cites struct {
		BaseURL string `yaml:"base-url"`
	}

	// DDOSProtection defines the configuration for DDOS protection.
	DDOSProtection struct {
		Enabled           bool    `yaml:"enabled"`
		PowComplexity     uint32  `yaml:"pow-complexity"`
		PowSalt           string  `yaml:"pow-salt"`
		RequestsPerSecond float64 `yaml:"requests-per-second"`
	}
)
