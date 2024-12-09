package repositories

import (
	pb "github.com/imnerocode/parser/parser"
	"gorm.io/gorm"
)

type ParserRepository struct {
	Db *gorm.DB
	pb.UnimplementedParserServiceServer
}
