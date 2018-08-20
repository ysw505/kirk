package parser

import "testing"

func TestParse(t *testing.T) {

	sample := `
    url: www.naver.com
    method: GET
  `

	parser := NewParser()

	result := parser.ParseBook(sample)

	if result.URL != "www.naver.com" {
		t.Error("Fail to parse")
	}

}
