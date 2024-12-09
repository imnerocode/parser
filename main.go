package main

import (
	"fmt"
	"log"
	"net"

	"github.com/imnerocode/parser/test_grpc"
	pb "github.com/imnerocode/parser/test_grpc/test_grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)
	pb.RegisterParserTestServiceServer(grpcServer, &test_grpc.ParserTestServiceImpl{})

	fmt.Println("Server is running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
