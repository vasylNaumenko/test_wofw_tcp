package app

import (
	"log/slog"
	"os"
)

// initLogger initializes the logger.
func (a *App) initLogger() {
	a.logger = slog.New(
		slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{
				Level: slog.LevelDebug,
			},
		),
	)
}
