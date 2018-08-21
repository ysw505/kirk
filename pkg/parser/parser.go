package parser

import (
	"github.com/ysw505/kirk/pkg/book"
	"gopkg.in/yaml.v2"
)


// Parser book
type Parser struct {
}

// NewParser function
func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) ParseRequest(yamlText string) (*Book.Request, error) {

	var req Book.Request

	err := yaml.Unmarshal([]byte(yamlText), &req)

	if err != nil {
		return nil, err
	}

	return &req, nil
}
