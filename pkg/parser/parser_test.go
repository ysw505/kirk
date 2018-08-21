package parser

import "testing"

func TestParse(t *testing.T) {

	sample := `
    url: www.naver.com
    method: GET
  `
	expected := "www.naver.com"

	parser := NewParser()

	result, err := parser.ParseRequest(sample)

	if err != nil {
		t.Error(err.Error())
	}

	if result.URL != "www.naver.com" {
		t.Error("Fail to parse")
	}

	t.Logf("parse sample , expected %s == result %s", expected, result.URL)
}
