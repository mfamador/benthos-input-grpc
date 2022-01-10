// Package services contains the business logic for Posting messages
package services //nolint
import (
	"github.com/Jeffail/benthos/v3/public/service"
	"github.com/rs/zerolog/log"
)

// Server exposes the interface to do operations on the Packets entity
type Server interface {
	Post(message string) error
}

type server struct {
	messageChan chan *service.Message
}

// NewService creates a Packets service
func NewService(messageChan chan *service.Message) Server {
	return &server{messageChan}
}

// Post posts message to repository benthos
func (s *server) Post(message string) error {
	// SEND TO BENTHOS
	log.Info().Msg(message)
	msg := service.NewMessage([]byte(message))
	s.messageChan <- msg
	return nil
}
