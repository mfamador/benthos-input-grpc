// Package input has the input related definition
package input

import (
	"context"
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/storage"
	"github.com/Jeffail/benthos/v3/public/service"
	"github.com/mfamador/benthos-input-grpc/internal/config"
	"github.com/mfamador/benthos-input-grpc/internal/server"
	"github.com/rs/zerolog/log"
)

var gRPCInputConfigSpec = service.NewConfigSpec().
	Summary("Creates an input that receives msgs from a gRPC server.")

func newGRPCInput(conf *service.ParsedConfig) (service.Input, error) {
	const maxchanns = 50
	input := gRPCInput{
		messageChan: make(chan *service.Message, maxchanns),
	}
	go func() {
		if err := server.RunApp(config.Config.Server, input.messageChan); err != nil {
			log.Panic().Msgf("failed to run app: %v", err)
		}
	}()

	return service.AutoRetryNacks(&input), nil
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
}

func (rts *gRPCInput) Connect(ctx context.Context) error {
	return nil
}

func (rts *gRPCInput) writeMessages(resp *storage.EntityQueryResult, msgType string) {
	for _, entity := range resp.Entities {
		record, _ := json.Marshal(entity.Properties)
		msg := service.NewMessage(record)
		msg.MetaSet("message_type", msgType)
		rts.messageChan <- msg
	}
}

func (rts *gRPCInput) Read(ctx context.Context) (*service.Message, service.AckFunc, error) {
	record := <-rts.messageChan
	return record, func(ctx context.Context, err error) error {
		return nil
	}, nil
}

func (rts *gRPCInput) Close(ctx context.Context) error {
	return nil
}
