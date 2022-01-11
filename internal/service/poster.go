// Package services contains the business logic for Posting messages
package service

import (
	"github.com/mfamador/benthos-input-grpc/internal/repository"
)

// Poster interface for poster services
type Poster interface {
	Post(message string) error
}

type poster struct {
	repo repository.Poster
}

// NewPoster creates a Packets service
func NewPoster(repo repository.Poster) Poster {
	return &poster{repo}
}

// Post posts a message to a repository
func (s *poster) Post(message string) error {
	return s.repo.Post(message)
}
