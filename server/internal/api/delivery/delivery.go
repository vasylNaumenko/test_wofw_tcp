package delivery

import "test_server/proto/cite"

type (
	CiteGRPCHandler interface {
		cite.CiteServiceServer
	}
)
