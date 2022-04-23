// @program:     enet
// @file:        handler.go
// @author:      edte
// @create:      2022-04-21 21:49
// @description:
package http

type Handler interface {
	Handle(req *Request, resp *Response)
}

type ResponseWriter interface {
	Write([]byte) (int, error)
}

type RouteHandle func(r *Request, w ResponseWriter)

var (
	NotFountHandle = RouteHandle(func(r *Request, w ResponseWriter) {
		_, _ = w.Write([]byte("404 not fount"))
	})

	IndexHandle = RouteHandle(func(r *Request, w ResponseWriter) {
		_, _ = w.Write([]byte("404 not fount"))
	})
)

var (
	defaultHandler = map[string]RouteHandle{
		"404": NotFountHandle,
		"/":   IndexHandle,
	}
)
