// Package server contains the gRPC instantiation
package server

import (
	"fmt"
	"net"

	"github.com/Jeffail/benthos/v3/public/service"
	"github.com/mfamador/benthos-input-grpc/internal/posterapi"
	"github.com/mfamador/benthos-input-grpc/pkg/posterv1"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

// Config defines the handler configuration
type Config struct {
	GrpcPort int `yaml:"grpcPort"`
}

// Start runs the gRPC server
func Start(grpcServer *grpc.Server, list net.Listener) error {
	log.Info().Msg("Test API gRPC server")
	return grpcServer.Serve(list)
}

// GetGRPCServer returns the gRPC server
func GetGRPCServer(conf Config, messageChann chan *service.Message) (*grpc.Server, net.Listener, error) {
	grpcServer := grpc.NewServer()
	posterv1.RegisterPosterServer(grpcServer, posterapi.NewPosterService(messageChann))
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.GrpcPort))
	if err != nil {
		return nil, nil, fmt.Errorf("error creating listener: %v", err)
	}
	return grpcServer, listener, nil
}

// RunApp runs the API
func RunApp(sConf Config, messageChann chan *service.Message) error {
	grpcServer, lis, err := GetGRPCServer(sConf, messageChann)
	if err != nil {
		return fmt.Errorf("failed to build grpcServer: %v", err)
	}
	if e := Start(grpcServer, lis); e != nil {
		log.Fatal().Msgf("Failed to start the gRPC server")
	}
	return nil
}
