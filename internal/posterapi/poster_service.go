// Package posterapi handles the message posts
package posterapi

import (
	"context"
	"fmt"

	"github.com/mfamador/benthos-input-grpc/internal/datasource/benthos"

	benthosSvc "github.com/Jeffail/benthos/v3/public/service"
	"github.com/mfamador/benthos-input-grpc/internal/service"
	"github.com/mfamador/benthos-input-grpc/pkg/posterv1"
)

// PosterService defines the interface for a Poster Service
type PosterService interface {
	Post(context.Context, *posterv1.PostRequest) (*posterv1.PostReply, error)
}

type posterService struct {
	serverService service.Poster
}

// NewPosterService instantiates a new poster service
func NewPosterService(messageChan chan *benthosSvc.Message) PosterService {
	benthosPoster := service.NewPoster(benthos.NewPoster(messageChan))
	return &posterService{
		serverService: benthosPoster,
	}
}

// Post posts a message to the chosen repository implementation
func (d *posterService) Post(_ context.Context, request *posterv1.PostRequest) (*posterv1.PostReply, error) {
	if err := d.serverService.Post(request.Message); err != nil {
		return nil, fmt.Errorf("error posting message: %v", err)
	}
	return new(posterv1.PostReply), nil
}
