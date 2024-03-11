package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"test_server/internal/app"
	"test_server/internal/config"
)

const (
	DefaultConfigPath = "./config.yaml"
)

func main() {
	// Load a configuration.
	cfg := config.MustLoad(DefaultConfigPath)

	// Create a new application.
	app.NewApp(cfg).Run(registerGracefulHandle())
}

// registerGracefulHandle registers a graceful shutdown handler for the application.
func registerGracefulHandle() context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		cancel()
	}()

	return ctx
}
