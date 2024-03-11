package config

import (
	"errors"
	"strings"
)

// Validate validates Config struct
func (c Config) Validate() error {
	var errs []string

	if e := c.Delivery.Validate(); len(e) > 0 {
		errs = append(errs, e...)
	}
	if e := c.Cites.Validate(); len(e) > 0 {
		errs = append(errs, e...)
	}

	if len(errs) > 0 {
		return errors.New(strings.Join(errs, ", "))
	}

	return nil
}

// Validate validates Cites struct
func (c Cites) Validate() []string {
	var errs []string

	if len(c.BaseURL) == 0 {
		errs = append(errs, "base-url is required")
	}

	return errs
}

// Validate validates Delivery struct
func (c Delivery) Validate() []string {
	var errs []string

	if len(c.GRPCServer.Port) == 0 {
		errs = append(errs, "port is required")
	}

	if c.GRPCServer.Timeout == 0 {
		errs = append(errs, "timeout is required")
	}

	return errs
}

// Validate validates DDOSProtection struct
func (d DDOSProtection) Validate() []string {
	var errs []string

	if !d.Enabled {
		return errs
	}
	if d.PowComplexity == 0 {
		errs = append(errs, "pow-complexity is required")
	}
	if d.RequestsPerSecond == 0 {
		errs = append(errs, "requests-per-second is required")
	}

	return errs
}
