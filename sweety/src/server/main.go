package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	bp "sweety/proto/business"
)

type server struct {
	bp.UnimplementedHelloHTTPServer
}

func (s *server) SayHello(cxt context.Context, req *bp.HelloHTTPRequest) (*bp.HelloHTTPResponse, error) {
	return &bp.HelloHTTPResponse{Message: "Hello, " + req.Name}, nil
}

func RegisterHelloHTTPServer(s grpc.ServiceRegistrar, srv bp.HelloHTTPServer) {
	s.RegisterService(&bp.HelloHTTP_ServiceDesc, srv)
}

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	RegisterHelloHTTPServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
