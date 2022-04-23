// @program:     enet
// @file:        transport.go
// @author:      edte
// @create:      2022-04-22 09:40
// @description:
package http

import (
	"github.com/edte/enet/net"
)

// 复用 tcp 连接池

var (
	transport = NewTransport(65535)
)

type Transport struct {
	conn    map[string]net.Conn
	maxConn int
}

func NewTransport(maxConn int) *Transport {
	return &Transport{
		conn:    make(map[string]net.Conn),
		maxConn: maxConn,
	}
}

func (t *Transport) get(r *Request) (net.Conn, error) {
	n, ok := t.conn[r.RemoteAddr]
	if ok {
		return n, nil
	}

	dial, err := net.Dial("tcp", r.RemoteAddr)
	if err != nil {
		return nil, err
	}

	t.conn[r.RemoteAddr] = dial

	return dial, nil
}
