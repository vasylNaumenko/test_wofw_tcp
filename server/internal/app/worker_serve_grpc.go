package app

import (
	"context"
	"net"
)

// serveGRPC serves gRPC server, registers gRPC services,
// and listens for the context to stop the server.
func serveGRPC(ctx context.Context, app *App) {
	logger := app.logger.With("worker", "serveGRPC")

	logger.Info("starting GRPC server",
		"port", app.config.Delivery.GRPCServer.Port,
	)

	listener, err := net.Listen("tcp", ":"+app.config.Delivery.GRPCServer.Port)
	if err != nil {
		logger.Error("failed to listen: %v", err)
		return
	}

	go func() {
		<-ctx.Done()
		app.clientGRPC.GracefulStop()
		logger.Info("GRPC is stopped")
	}()

	if err = app.clientGRPC.Serve(listener); err != nil {
		logger.Error("failed to serve", err)
	}
}
