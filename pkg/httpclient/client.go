package httpclient

import (
	"net/http"

	"context"
	"github.com/ysw505/kirk/pkg/book"
	"net/http/httptrace"
	"fmt"
)

type Client struct {
	c *http.Client
	req *http.Request
}

func NewClient() *Client {

	client := Client{}
/*
	tr := &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
*/
	httpClient := &http.Client{
		Transport: http.DefaultTransport,
	}

	client.c = httpClient

	return &client
}

func (c *Client) SetRequest(req *Book.Request) {

	newReq, _ := http.NewRequest(req.Method, req.URL, nil)

	trace := &httptrace.ClientTrace{
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			fmt.Printf("DNS Info: %+v\n", dnsInfo)
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("Got Conn: %+v\n", connInfo)
		},
		ConnectStart: func(_, _ string) {

			fmt.Println("STart connect")
		},
	}

	for k, v := range req.Header {
		newReq.Header.Set(k,v)
	}

	fmt.Println("test")
	newReq = newReq.WithContext(httptrace.WithClientTrace(context.Background(), trace))

	//c.c.Transport.RoundTrip(newReq)

	c.req = newReq
}

func (c *Client) Do() http.Response {
	c.c.Do(c.req)
}