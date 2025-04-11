package main

import (
	"context"
	"fmt"
	example "grpc-demo/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	example.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *example.HelloRequest) (*example.HelloResponse, error) {
	return &example.HelloResponse{
		Message: fmt.Sprintf("Hello, %s!", req.GetName()),
	}, nil
}

func startServer(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	example.RegisterGreeterServer(s, &server{})
	log.Printf("Server listening at %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	go startServer(":50051")
	go startServer(":50052")

	// 阻止主程序退出
	select {}
}
