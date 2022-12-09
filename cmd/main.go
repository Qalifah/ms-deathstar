package main

import (
	"context"
	"deathstar/database"
	"deathstar/handler"
	"deathstar/lib/pubsub/pulsar"
	"deathstar/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	pubSubSrv, err := pulsar.New("pulsar://localhost:6650", "targets.acquired")
	if err != nil {
		log.Fatalln(err)
	}
	ctx := context.Background()
	db, err := database.New(ctx, "mongodb://localhost:27017", "service")
	if err != nil {
		log.Fatalln(err)
	}

	ctrl := handler.New(db, pubSubSrv)
	lis, err := net.Listen("tcp", "127.0.0.1:50050")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterDeathstarServer(s, ctrl)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
