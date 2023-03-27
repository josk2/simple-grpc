package main

import (
	"context"
	"fmt"
	"log"

	pb "simplegrpc/protoc/helloworld"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())

	rep := pb.HelloReply{Message: fmt.Sprintf("Hello %s", in.GetName())}
	return &rep, nil
}
