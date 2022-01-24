// Package input has the input related definition
package input

import (
	"context"

	"github.com/Jeffail/benthos/v3/public/service"
	"github.com/mfamador/benthos-input-grpc/internal/config"
	"github.com/mfamador/benthos-input-grpc/internal/server"
	"github.com/rs/zerolog/log"
)

const (
	fieldMaxInFlight   = "max_in_flight"
	defaultMaxInFlight = 10
)

var gRPCInputConfigSpec = service.NewConfigSpec().
	Summary("Creates an input that receives msgs from a gRPC server.").
	Field(service.NewIntField(fieldMaxInFlight).Default(defaultMaxInFlight))

func newGRPCInput(conf *service.ParsedConfig) (service.Input, error) {
	mf, err := conf.FieldInt(fieldMaxInFlight)
	if err != nil {
		log.Panic().Msgf("could not get max in flight value: %v", err)
	}
	input := gRPCInput{
		messageChan: make(chan *service.Message, mf),
		errorChan:   make(chan error),
	}
	go func() {
		if err := server.RunApp(config.Config.Server, input.messageChan, input.errorChan); err != nil {
			log.Panic().Msgf("failed to run app: %v", err)
		}
	}()
	return &input, nil
}

func init() {
	err := service.RegisterInput(
		"grpc_server", gRPCInputConfigSpec,
		func(conf *service.ParsedConfig, mgr *service.Resources) (service.Input, error) {
			return newGRPCInput(conf)
		})
	if err != nil {
		panic(err)
	}
}

//------------------------------------------------------------------------------

type gRPCInput struct {
	messageChan chan *service.Message
	errorChan   chan error
}

func (rts *gRPCInput) Connect(ctx context.Context) error {
	return nil
}

func (rts *gRPCInput) Read(ctx context.Context) (*service.Message, service.AckFunc, error) {
	record := <-rts.messageChan
	return record, func(ctx context.Context, err error) error {
		rts.errorChan <- err
		return err
	}, nil
}

func (rts *gRPCInput) Close(ctx context.Context) error {
	return nil
}
