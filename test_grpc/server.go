package test_grpc

import (
	"context"

	pb "github.com/imnerocode/parser/test_grpc/test_grpc"
)

type ParserTestServiceImpl struct {
	pb.UnimplementedParserTestServiceServer
}

func (si *ParserTestServiceImpl) Parser(ctx context.Context, rq *pb.ParserTestRequest) (*pb.ParserTestResponse, error) {
	response := &pb.ParserTestResponse{
		FilePath:   rq.GetFilePath(),
		FileFormat: rq.GetFileFormat(),
	}
	return response, nil
}
