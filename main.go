package main

import (
	"fmt"

	"github.com/imnerocode/parser/parser"
)

func main() {
	/*
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
	*/

	meshAttr, err := parser.ParserToOBJ("utils/models/Hello_Node.obj")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Vertices: %+v\n", meshAttr.Vertices)
	fmt.Printf("UVs: %+v\n", meshAttr.UV)
	fmt.Printf("Normals: %+v", meshAttr.Normals)

}
