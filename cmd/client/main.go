// Package main starts the Data API
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"
	"math/rand"
	"time"

	"github.com/mfamador/benthos-input-grpc/internal/config"
	"github.com/mfamador/benthos-input-grpc/pkg/posterv1"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func main() {
	log.Info().Timestamp().Msg("Client gRPC Poster API")
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", config.Config.Server.GrpcPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error().Timestamp().Msg(err.Error())
	}
	client := posterv1.NewPosterClient(conn)
	for {
		const size = 10
		request := posterv1.PostRequest{Message: fmt.Sprintf(`{"foo":%q,"bar":%q}`, randSeq(size), randSeq(size))}
		reply, err := client.Post(context.Background(), &request)
		if err != nil {
			log.Error().Timestamp().Msg(err.Error())
		} else {
			log.Info().Timestamp().Msgf("msg sent: %v", reply)
		}
		time.Sleep(time.Second)
	}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
