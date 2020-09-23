package main

import (
	"fmt"
	"io"
	"log"
	"net"

	pb "streaming-client/gen/grpc-gateway/gen"

	"google.golang.org/grpc"
)

const (
	port = ":50052"
)

// SayHello implements helloworld.GreeterServer.SayHello
func (c *Server) SayHello(srv pb.Greeter_SayHelloServer) error {
	for {
		msg, err := srv.Recv()
		if err == io.EOF {
			return srv.SendAndClose(&pb.HelloReply{Message: "hello"})
		}

		if err != nil {
			log.Fatalf("could not recieve srv: %v", err)
		}
		fmt.Println(msg.GetName())
	}
}

// Server ...
type Server struct {
	pb.UnimplementedGreeterServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
