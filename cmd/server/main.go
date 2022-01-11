// Package main starts the gRPC server and benthos pip
package main

import (
	"github.com/Jeffail/benthos/v3/lib/service"
	_ "github.com/mfamador/benthos-input-grpc/input"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Timestamp().Msg("Poster gRPC API Server")

	service.Run()
}
