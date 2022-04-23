// @program:     enet
// @file:        response.go
// @author:      edte
// @create:      2022-04-21 21:29
// @description:
package http

type Response struct {
	raw []byte
}

func NewResponse() *Response {
	return &Response{}
}

func (r *Response) Write(bytes []byte) (int, error) {
	r.raw = append(r.raw, bytes...)
	return 0, nil
}
