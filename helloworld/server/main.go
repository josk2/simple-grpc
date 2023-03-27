package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "simplegrpc/protoc/helloworld"

	"google.golang.org/grpc"
)

var port = flag.Int("port", 50051, "the server port")

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	fmt.Printf("Server is running at %v ", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
