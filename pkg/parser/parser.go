package parser

import "github.com/ysw505/kirk/pkg/book"


// Parser book
type Parser struct {
}

// NewParser function
func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) ParseBook(yamlText string) *Book.Book {
	return &Book.Book{}
}
