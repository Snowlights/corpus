package main

import (
	"context"
	"github.com/Snowlights/corpus/cache"
	"github.com/Snowlights/corpus/model"
	"github.com/Snowlights/corpus/model/daoimpl"
	"github.com/Snowlights/corpus/router"
	corpus "github.com/Snowlights/pub/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)

const (
	port = ":50051"
)

func InitFunc() (err error) {
	fun := "main.InitFunc -->"
	ctx := context.Background()

	// You can put more preparations here
	model.Prepare(ctx)
	daoimpl.PrePare(ctx)
	cache.Prepare(ctx)

	log.Println("%s usccess %s",ctx,fun)
	return nil
}

func main() {
	err := InitFunc()
	if err != nil{
		log.Fatalf("failed to init: %v", err)
	}
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	corpus.RegisterCorpusServiceServer(s, new(router.CorpusServiceTmp))
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	time.Sleep(time.Hour)
}