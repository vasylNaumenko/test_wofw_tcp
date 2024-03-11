package repository

type (
	Cite interface {
		GetCite() (string, error)
	}
)
