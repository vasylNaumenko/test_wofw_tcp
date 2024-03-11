package app

import (
	"context"
	"log/slog"

	"google.golang.org/grpc"

	"test_server/internal/api/delivery"
	"test_server/internal/api/repository"
	"test_server/internal/api/services"
	"test_server/internal/config"
)

// App defines main application struct.
type (
	App struct {
		config *config.Config
		logger *slog.Logger

		clientGRPC     *grpc.Server
		citeRepository repository.Cite

		powService  services.POWService
		citeService services.CiteService

		citeGRPCHandler delivery.CiteGRPCHandler
	}
	worker func(ctx context.Context, a *App)
)

// NewApp creates a new application.
func NewApp(cfg *config.Config) *App {
	return &App{
		config: cfg,
	}
}

// Run runs the application. It uses the context to stop the application.
func (a *App) Run(ctx context.Context) {
	a.initLogger()

	a.registerRepositories()
	a.registerServices()

	a.initGRPCClient()
	a.runWorkers(ctx)
}
