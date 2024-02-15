// Package benthos contains the benthos repository implementation
package benthos

import (
	"github.com/benthosdev/benthos/v4/public/service"
	"github.com/mfamador/benthos-input-grpc/internal/repository"
)

type server struct {
	messageChan chan *service.Message
	errorChan   chan error
}

// NewPoster handles post messages to benthos
func NewPoster(messageChan chan *service.Message, errorChan chan error) repository.Poster {
	return &server{messageChan: messageChan, errorChan: errorChan}
}

func (s *server) Post(message string) error {
	msg := service.NewMessage([]byte(message))
	s.messageChan <- msg

	if err := <-s.errorChan; err != nil {
		return err
	}
	return nil
}
