package cite

import (
	"context"
	"log/slog"
	"time"

	"golang.org/x/time/rate"

	"test_server/internal/api/delivery"
	"test_server/internal/api/services"
	"test_server/proto/cite"
)

var _ delivery.CiteGRPCHandler = &Handler{}

// Handler - struct for handling Cite requests.
type Handler struct {
	cite.UnimplementedCiteServiceServer
	log               *slog.Logger
	powService        services.POWService
	ddosEnabled       bool
	citeService       services.CiteService
	timeout           time.Duration
	requestsPerSecond float64
}

// NewHandler - constructor.
func NewHandler(
	log *slog.Logger,
	timeout time.Duration,
	powEnabled bool,
	requestsPerSecond float64,
	powService services.POWService,
	citeService services.CiteService,
) *Handler {
	return &Handler{
		log:               log.With("handler", "cite"),
		timeout:           timeout,
		ddosEnabled:       powEnabled,
		requestsPerSecond: requestsPerSecond,
		powService:        powService,
		citeService:       citeService,
	}
}

// GetCite - method for handling Cite requests.
// It checks the POW and returns the Cite.
func (h Handler) GetCite(stream cite.CiteService_GetCiteServer) error {
	ctx, cancel := context.WithTimeout(stream.Context(), h.timeout)
	defer cancel()

	requestChan := make(chan request)
	nonce := h.powService.GetNonce()
	limiter := rate.NewLimiter(rate.Limit(h.requestsPerSecond), int(h.requestsPerSecond))

	for {
		// Start a goroutine that calls stream.Recv() and sends the result to requestChan
		go func() {
			req, err := stream.Recv()
			requestChan <- request{req: req, err: err}
		}()

		select {
		case <-ctx.Done():
			return ctx.Err()
		case data := <-requestChan:
			if h.ddosEnabled && !limiter.Allow() {
				return ErrorRateLimit
			}

			closeConnection, err := h.handleRequestData(stream, data, nonce)
			if err != nil {
				return err
			}

			if closeConnection {
				return nil
			}
		}
	}
}
