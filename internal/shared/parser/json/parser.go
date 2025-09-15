package json

import (
	"encoding/json"
	"io"
	"os"
)

type Parser[T any] struct {
}

func NewParser[T any]() *Parser[T] {
	return &Parser[T]{}
}

func (p *Parser[T]) Parse(filePath string) (*T, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = file.Close()
	}()

	bytes, _ := io.ReadAll(file)

	var appConfig T
	err = json.Unmarshal(bytes, &appConfig)
	if err != nil {
		return nil, err
	}

	return &appConfig, nil
}
