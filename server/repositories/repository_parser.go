package repositories

import (
	"context"

	pb "github.com/imnerocode/parser/parser"
)

type ParserRepository struct {
	pb.UnimplementedParserServiceServer
}

func (r *ParserRepository) ParserModel(ctx context.Context, rq *pb.ParserRequest) (*pb.ParserResponse, error) {
	// Llamamos a la función que procesa el archivo .obj
	meshAttr, err := pb.ParserToOBJ(rq.GetFilePath())
	if err != nil {
		return nil, err
	}
	response := &pb.ParserResponse{
		MeshAttr: &pb.MeshAttributes{
			Vertices: meshAttr.Vertices,
			Indices:  meshAttr.Indices,
			Normals:  meshAttr.Normals,
			UV:       meshAttr.UV,
		},
	}

	return response, nil
}
