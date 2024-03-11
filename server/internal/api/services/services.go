package services

type (
	// CiteService describes the Cite service.
	CiteService interface {
		GetCite() (string, error)
	}

	// POWService describes the POW service.
	POWService interface {
		GetPow(nonce string) string
		ValidatePOW(riddle, proof string) bool
		GetNonce() string
	}
)
