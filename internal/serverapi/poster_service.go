// Package serverapi handles the messages posts
package serverapi

import (
	"context"
	"fmt"

	"github.com/mfamador/benthos-input-grpc/internal/datasource/benthos"

	benthosSvc "github.com/Jeffail/benthos/v3/public/service"
	"github.com/mfamador/benthos-input-grpc/internal/service"
	"github.com/mfamador/benthos-input-grpc/pkg/posterv1"
)

// PosterService exposes the interface
type PosterService interface {
	Post(context.Context, *posterv1.PostRequest) (*posterv1.PostReply, error)
}

type posterService struct {
	serverService service.Poster
}

// NewPosterService instantiates a new server service
func NewPosterService(messageChann chan *benthosSvc.Message) PosterService {
	ss := service.NewPoster(benthos.NewPoster(messageChann))
	return &posterService{
		serverService: ss,
	}
}

// Post posts a message
func (d *posterService) Post(_ context.Context, request *posterv1.PostRequest) (*posterv1.PostReply, error) {
	if err := d.serverService.Post(request.Message); err != nil {
		return nil, fmt.Errorf("error posting message: %v", err)
	}
	return new(posterv1.PostReply), nil
}
