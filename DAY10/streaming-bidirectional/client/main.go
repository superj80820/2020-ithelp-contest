package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	pb "streaming-bidirectional/gen/grpc-gateway/gen"

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
	stream, err := c.SayHello(context.Background())
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// make blocking channel
	waitc := make(chan struct{})

	// send feeds to the stream ( go routine )
	go func() {
		for i := 1; i <= 5; i++ {
			if err := stream.Send(&pb.HelloRequest{Name: name}); err != nil {
				log.Fatalf("error while sending feed: %v", err)
			}
			time.Sleep(time.Second)
		}
		if err := stream.CloseSend(); err != nil {
			log.Fatalf("failed to close stream: %v", err)
		}
	}()

	// recieve feeds frrom the stream ( go routine )
	go func() {
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("failed to recieve: %v", err)
				close(waitc)
				return
			}

			fmt.Println("New feed recieved : ", msg.GetMessage())
		}

	}()

	<-waitc
}
