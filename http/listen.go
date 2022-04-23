// @program:     enet
// @file:        listen.go
// @author:      edte
// @create:      2022-04-21 21:50
// @description:
package http

func Listen(addr string, handler Handler) error {
	return NewServer(addr, handler).Listen()
}

func HandleFunc(r string, h RouteHandle) {
	defaultHandler[r] = h
}
