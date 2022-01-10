// Package services contains the business logic for Posting messages
package services

import (
	"github.com/mfamador/benthos-input-grpc/internal/repository"
	"github.com/rs/zerolog/log"
)

// Server exposes the interface to do operations on the Packets entity
type Server interface {
	Post(message string) error
}

type server struct {
	repo repository.Server
}

// NewServer creates a Packets service
func NewServer(repo repository.Server) Server {
	return &server{repo}
}

// Post posts message to repository benthos
func (s *server) Post(message string) error {
	log.Info().Msg(message)
	return s.repo.Post(message)
}
