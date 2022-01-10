// Package services contains the business logic for Posting messages
package services //nolint
import "github.com/rs/zerolog/log"

// Server exposes the interface to do operations on the Packets entity
type Server interface {
	Post(message string) error
}

type server struct{}

// NewService creates a Packets service
func NewService() Server {
	return &server{}
}

// Post posts message to repository benthos
func (s *server) Post(message string) error {
	// SEND TO BENTHOS
	log.Info().Msg(message)
	return nil
}
