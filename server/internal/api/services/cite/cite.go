package cite

import (
	"log/slog"

	"test_server/internal/api/repository"
	"test_server/internal/api/services"
)

var _ services.CiteService = &Service{}

type (
	// Service defines the main service struct.
	Service struct {
		citeRepo repository.Cite
		log      *slog.Logger
	}
)

// NewService creates a new service.
func NewService(citeRepo repository.Cite, logger *slog.Logger) *Service {
	return &Service{
		citeRepo: citeRepo,
		log:      logger.With("service", "cite"),
	}
}

// GetCite returns a new cite.
func (s Service) GetCite() (string, error) {
	cite, err := s.citeRepo.GetCite()
	if err != nil {
		s.log.Error("failed to get cite", err)
		return "", err
	}

	return cite, nil
}
