package app

import "test_server/internal/api/repository/cite"

// registerRepositories registers repositories.
func (a *App) registerRepositories() {
	a.citeRepository = cite.NewRepository(a.config.Cites.BaseURL, a.logger)
}
