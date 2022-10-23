package json

import (
	"encoding/json"
	"io"
	"os"
)

type Parser[T interface{}] struct {
}

func NewParser[T interface{}]() *Parser[T] {
	return &Parser[T]{}
}

func (p *Parser[T]) Parse(filePath string) (*T, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, _ := io.ReadAll(file)

	var appConfig T
	err = json.Unmarshal(bytes, &appConfig)
	if err != nil {
		return nil, err
	}

	return &appConfig, nil
}
