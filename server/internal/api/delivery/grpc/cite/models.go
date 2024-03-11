package cite

import (
	"errors"

	"test_server/proto/cite"
)

// Error variables.
var (
	ErrorInvalidPOW = errors.New("invalid POW")
	ErrorRateLimit  = errors.New("rate limit exceeded")
)

// eof is a string for EOF.
const eof = "EOF"

// request is a struct for request.
type request struct {
	req *cite.POWRequest
	err error
}
