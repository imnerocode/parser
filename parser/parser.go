package parser

import (
	"fmt"
	"os"
)

func ParserToOBJ(obj_file string) (*MeshAttributes, error) {
	file, err := os.Open(obj_file)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var content []byte

	_, err = file.Read(content)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(content))
	return nil, nil
}
