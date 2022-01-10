// Package serverapi handles the messages posts
package serverapi

import (
	"context"
	"fmt"
	"github.com/mfamador/benthos-input-grpc/internal/services"
	"github.com/mfamador/benthos-input-grpc/pkg/serverv1"
)

// ServerService exposes the interface
type ServerService interface {
	Post(context.Context, *serverv1.PostRequest) (*serverv1.PostReply, error)
}

type serverService struct {
	serverService services.Server
}

// NewServerService instantiates a new server service
func NewServerService() ServerService {
	ss := services.NewService()

	return &serverService{
		serverService: ss,
	}
}

// Post post a message
func (d *serverService) Post(_ context.Context, request *serverv1.PostRequest) (*serverv1.PostReply, error) {
	if err := d.serverService.Post(request.Message); err != nil {
		return nil, fmt.Errorf("error posting message: %v", err)
	}
	return new(serverv1.PostReply), nil
}
