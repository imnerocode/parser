package parser

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ConvertToFloat32(arrayString []string) ([]float32, error) {
	var arrayFloat32 []float32

	for _, v := range arrayString {
		// Limpiar espacios alrededor del string
		trimmed := strings.TrimSpace(v)

		// Ignorar valores vacíos
		if trimmed == "" {
			continue
		}

		// Intentar convertir a float32
		arrayConverted, err := strconv.ParseFloat(trimmed, 32)
		if err != nil {
			return nil, fmt.Errorf("error converting '%s' to float32: %v", v, err)
		}
		arrayFloat32 = append(arrayFloat32, float32(arrayConverted))
	}
	return arrayFloat32, nil
}

func ParserToOBJ(obj_file string) (*MeshAttributes, error) {

	content, err := os.ReadFile(obj_file)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")

	var vertices []string
	var cleanedVertices []string
	var uvs []string
	var cleanedUVs []string
	var normals []string
	var cleanedNormals []string

	for _, line := range lines {
		if strings.HasPrefix(line, "v ") {
			parts := strings.Fields(line)
			if len(parts) == 4 {
				vertices = append(vertices, parts...)
			} else {
				log.Println("Error line of vertices, number of parts incorrect", line)
			}
		}
		if strings.HasPrefix(line, "vt ") {
			parts := strings.Fields(line)
			if len(parts) == 3 {
				uvs = append(uvs, parts...)
			} else {
				log.Println("Error line of textures, number of parts incorrect", line)
			}
		}
		if strings.HasPrefix(line, "vn ") {
			parts := strings.Fields(line)
			if len(parts) == 4 {
				normals = append(normals, parts...)
			} else {
				log.Println("Error line of textures, number of parts incorrect", line)
			}
		}

	}
	for _, v := range vertices {
		cleaned := strings.ReplaceAll(v, "v", "")
		cleanedVertices = append(cleanedVertices, cleaned)
	}
	for _, u := range uvs {
		cleaned := strings.ReplaceAll(u, "vt", "")
		cleanedUVs = append(cleanedUVs, cleaned)
	}
	for _, n := range normals {
		cleaned := strings.ReplaceAll(n, "vn", "")
		cleanedNormals = append(cleanedNormals, cleaned)
	}

	arrayVertices, err := ConvertToFloat32(cleanedVertices)
	arrayUVs, _ := ConvertToFloat32(cleanedUVs)
	arrayNormals, _ := ConvertToFloat32(cleanedNormals)
	if err != nil {
		panic(err)
	}
	var meshAttr MeshAttributes

	meshAttr.Vertices = arrayVertices
	meshAttr.UV = arrayUVs
	meshAttr.Normals = arrayNormals

	return &meshAttr, nil
}
