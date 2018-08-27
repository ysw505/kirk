package httpclient

import (
	"testing"
	"github.com/ysw505/kirk/pkg/book"
)

func TestClient_SetRequest(t *testing.T) {

	c := NewClient()

	testBook := Book.Request{
		URL: "http://www.naver.com",
		Method: "GET",
	}

	c.SetRequest(&testBook)

	c.Do()
}