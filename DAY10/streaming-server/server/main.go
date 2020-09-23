package main

import (
	"log"
	"net"

	pb "streaming-server/gen/grpc-gateway/gen"

	"google.golang.org/grpc"
)

const (
	port = ":50052"
)

// SayHello implements helloworld.GreeterServer.SayHello
func (s *Server) SayHello(req *pb.HelloRequest, srv pb.Greeter_SayHelloServer) error {
	log.Printf("Received: %v", req.GetName())
	for i := 0; i < 10; i++ {
		srv.Send(&pb.HelloReply{Message: req.GetName()})
	}
	return nil
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
