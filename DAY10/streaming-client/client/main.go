package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "streaming-client/gen/grpc-gateway/gen"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50052"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := c.SayHello(ctx)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	for i := 0; i <= 10; i++ {
		fmt.Printf("sending %v into the stream\n", i)
		stream.Send(&pb.HelloRequest{Name: name})
		time.Sleep(100 * time.Millisecond)
	}
}
