// Package main starts the Data API
package main

import (
	"github.com/mfamador/benthos-input-grpc/internal/config"
	"github.com/mfamador/benthos-input-grpc/internal/server"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Timestamp().Msg("Test Server gRPC API")
	if err := server.RunApp(config.Config.Server); err != nil {
		log.Panic().Msgf("failed to run app: %v", err)
	}
}
