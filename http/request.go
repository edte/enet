// @program:     enet
// @file:        request.go
// @author:      edte
// @create:      2022-04-21 21:29
// @description:
package http

import (
	"encoding/json"
	"io"
)

type Request struct {
	Method           string
	URL              string
	Version          string
	Header           map[string][]string
	Body             io.Reader
	ContentLength    int64
	TransferEncoding []string
	Close            bool
	Host             string
	Form             map[string][]string
	PostForm         map[string][]string
	RemoteAddr       string

	Response *Response `json:"-"`
}

func (r *Request) String() string {
	d, _ := json.Marshal(r)
	return string(d)
}

func NewRequest(method, url string, body io.Reader) (req *Request, err error) {
	return &Request{
		Method: method,
		URL:    url,
		Body:   body,
	}, nil
}
