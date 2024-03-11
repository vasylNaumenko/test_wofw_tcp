package pow

import (
	"github.com/bwesterb/go-pow"
	"github.com/google/uuid"

	"test_server/internal/api/services"
)

var _ services.POWService = &Service{}

type (
	// Service defines the main service struct.
	Service struct {
		complexity uint32
		salt       string
	}
)

// NewService creates a new service.
func NewService(complexity uint32, salt string) *Service {
	return &Service{
		complexity: complexity,
		salt:       salt,
	}
}

// ValidatePOW validates proof of work.
func (s Service) ValidatePOW(riddle, proof string) bool {
	ok, _ := pow.Check(riddle, proof, []byte(s.salt))
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
