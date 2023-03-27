package main

import (
	"context"
	"fmt"
	"log"

	"simplegrpc/helloworld/protoc"
)

type server struct {
	protoc.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *protoc.HelloRequest) (*protoc.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())

	rep := protoc.HelloReply{Message: fmt.Sprintf("Hello %s", in.GetName())}
	return &rep, nil
}

func (s server) SayHelloAgain(ctx context.Context, in *protoc.HelloRequest) (*protoc.HelloReply, error) {
	log.Printf("Received Again: %v", in.GetName())

	rep := protoc.HelloReply{Message: fmt.Sprintf("Hello again %s", in.GetName())}
	return &rep, nil
}
