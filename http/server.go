// @program:     enet
// @file:        server.go
// @author:      edte
// @create:      2022-04-21 21:55
// @description:
package http

import (
	"github.com/edte/enet/net"
)

type Server struct {
	Addr    string
	handler Handler
}

func NewServer(addr string, handler Handler) *Server {
	return &Server{Addr: addr, handler: handler}
}

func (s *Server) Listen() error {
	listen, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}

	for {
		ac, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		c := newConn(s, ac)
		go c.Start()
	}
}
