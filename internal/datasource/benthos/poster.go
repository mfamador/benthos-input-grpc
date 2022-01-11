// Package benthos contains the benthos repository implementation
package benthos

import (
	"github.com/Jeffail/benthos/v3/public/service"
	"github.com/mfamador/benthos-input-grpc/internal/repository"
)

type server struct {
	messageChan chan *service.Message
}

// NewPoster handles post messages to benthos
func NewPoster(messageChan chan *service.Message) repository.Poster {
	return &server{messageChan: messageChan}
}

func (s *server) Post(message string) error {
	msg := service.NewMessage([]byte(message))
	s.messageChan <- msg
	return nil
}
