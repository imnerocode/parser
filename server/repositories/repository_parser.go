package repositories

import (
	"context"

	pb "github.com/imnerocode/parser/parser"
	"gorm.io/gorm"
)

type ParserRepository struct {
	Db *gorm.DB
	pb.UnimplementedParserServiceServer
}

func NewParserRepository(db *gorm.DB) *ParserRepository {
	return &ParserRepository{Db: db}
}

func (r *ParserRepository) ParserModel(ctx context.Context, rq *pb.ParserRequest) (*pb.ParserResponse, error) {
	response := &pb.ParserResponse{}

	return response, nil
}
