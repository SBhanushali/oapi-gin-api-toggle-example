package api

import (
	"context"

	"github.com/SBhanushali/oapi-gin-api-toggle-example/server"
)

type svc struct{}

var _ server.StrictServerInterface = (*svc)(nil)

func New() server.StrictServerInterface {
	return &svc{}
}

func (s *svc) GetGreeting(ctx context.Context, params server.GetGreetingRequestObject) (server.GetGreetingResponseObject, error) {
	message := "Hello, World!"
	return server.GetGreeting200JSONResponse{
		Message: &message,
	}, nil
}
