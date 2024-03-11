package cite

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"test_server/internal/api/repository"
)

var _ repository.Cite = &Repository{}

// Repository implements repository.TraderService.
type Repository struct {
	baseURL string
	log     *slog.Logger
}

type response struct {
	Fortune string `json:"fortune"`
}

// NewRepository constructor Repository.
func NewRepository(baseURL string, log *slog.Logger) *Repository {
	return &Repository{
		baseURL: baseURL,
		log:     log.With("repository", "cite"),
	}
}

// GetCite returns a random quote from the external API.
func (r Repository) GetCite() (string, error) {
	resp, err := http.Get(r.baseURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var res response
	if err = json.Unmarshal(body, &res); err != nil {
		return "", err
	}

	return res.Fortune, nil
}
