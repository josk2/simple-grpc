package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"sync"
	"time"

	"simplegrpc/helloworld/protoc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const defaultName = "world"

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect %v", err)
	}
	defer conn.Close()

	c := protoc.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	timearr := []int{2, 4, 1, 3, 2}
	wg := sync.WaitGroup{}

	for _, i := range timearr {
		wg.Add(1)
		go doSayHello(ctx, c, &wg, i)
	}

	wg.Wait()
	fmt.Println("Done")

}

func doSayHello(ctx context.Context, c protoc.GreeterClient, wg *sync.WaitGroup, num int) {
	r, err := c.SayHello(ctx, &protoc.HelloRequest{Name: fmt.Sprintf("#%d", num)})
	if err != nil {
		log.Fatalf("could not gree %v", err)
	}
	log.Printf("Get response: %s", r.GetMessage())
	wg.Done()
}
