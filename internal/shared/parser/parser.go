package parser

type Parser[T interface{}] interface {
	Parse(filePath string) (*T, error)
}
