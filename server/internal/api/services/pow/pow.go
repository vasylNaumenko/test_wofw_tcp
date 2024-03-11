package pow

import (
	"github.com/bwesterb/go-pow"
	"github.com/google/uuid"

	"test_server/internal/api/services"
)

var _ services.POWService = &Service{}

const powMessage = "pow"

type (
	// Service defines the main service struct.
	Service struct {
		complexity uint32
	}
)

// NewService creates a new service.
func NewService(complexity uint32) *Service {
	return &Service{
		complexity: complexity,
	}
}

// ValidatePOW validates proof of work.
func (s Service) ValidatePOW(riddle, proof string) bool {
	ok, _ := pow.Check(riddle, proof, []byte(powMessage))
	return ok
}

// GetNonce returns a new nonce.
func (s Service) GetNonce() string {
	return uuid.New().String()
}

// GetPow returns a new proof of work.
func (s Service) GetPow(nonce string) string {
	return pow.NewRequest(s.complexity, []byte(nonce))
}
