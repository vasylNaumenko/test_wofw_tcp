package delivery

import "test_server/proto/cite"

type (
	// CiteGRPCHandler describes the Cite gRPC handler.
	CiteGRPCHandler interface {
		cite.CiteServiceServer
	}
)
