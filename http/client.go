// @program:     enet
// @file:        client.go
// @author:      edte
// @create:      2022-04-21 21:29
// @description:
package http

import (
	"bytes"
	"io"
)

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) send(req *Request) (resp *Response, err error) {
	conn, err := transport.get(req)
	if err != nil {
		return nil, err
	}
	_, err = conn.Write([]byte(req.String()))
	if err != nil {
		return nil, err
	}

	b := bytes.Buffer{}
	_, err = b.ReadFrom(conn)
	if err != nil {
		return nil, err
	}

	return NewResponse(), nil
}

func (c *Client) Do(req *Request) (resp *Response, err error) {
	for {
		send, err := c.send(req)
		return send, err
	}
}

func (c *Client) Get(url string) (resp *Response, err error) {
	r, _ := NewRequest("get", url, nil)
	r.RemoteAddr = "127.0.0.1:1234"
	r.Host = "/"
	return c.Do(r)
}

func (c *Client) Post(url, contentType string, body io.Reader) (resp *Response, err error) {
	r, _ := NewRequest("post", url, body)
	return c.Do(r)
}

var (
	defaultClient = NewClient()
)

func Get(url string) (resp *Response, err error) {
	return defaultClient.Get(url)
}

func Post(url, contentType string, body io.Reader) (resp *Response, err error) {
	return defaultClient.Post(url, contentType, body)
}
