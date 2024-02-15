// Package main starts the gRPC server and benthos pip
package main

import (
	"context"

	_ "github.com/benthosdev/benthos/v4/public/components/all"
	"github.com/benthosdev/benthos/v4/public/service"
	_ "github.com/mfamador/benthos-input-grpc/internal/input"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Timestamp().Msg("Poster gRPC API Server")

	service.RunCLI(context.Background())
}
