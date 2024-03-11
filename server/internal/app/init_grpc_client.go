package app

import (
	"google.golang.org/grpc"

	citehandler "test_server/internal/api/delivery/grpc/cite"
	"test_server/proto/cite"
)

// initGRPCClient initializes gRPC client.
func (a *App) initGRPCClient() {
	a.clientGRPC = grpc.NewServer()
	a.registerGRPCHandlers()
}

// registerGRPCHandlers registers gRPC handlers.
func (a *App) registerGRPCHandlers() {
	a.citeGRPCHandler = citehandler.NewHandler(
		a.logger,
		a.config.Delivery.GRPCServer.Timeout,
		a.config.Delivery.DDOSProtection.Enabled,
		a.config.Delivery.DDOSProtection.RequestsPerSecond,
		a.powService,
		a.citeService,
	)
	cite.RegisterCiteServiceServer(a.clientGRPC, a.citeGRPCHandler)
}
