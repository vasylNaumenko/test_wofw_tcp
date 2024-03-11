package app

import (
	"test_server/internal/api/services/cite"
	"test_server/internal/api/services/pow"
)

// registerServices register services in app struct.
func (a *App) registerServices() {
	a.powService = pow.NewService(a.config.Delivery.DDOSProtection.PowComplexity)
	a.citeService = cite.NewService(a.citeRepository, a.logger)
}
