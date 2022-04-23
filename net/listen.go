// @program:     enet
// @file:        listen.go
// @author:      edte
// @create:      2022-04-10 17:16
// @description:
package net

import (
	"strconv"
)

var (
	// 支持的协议
	validProtocol = map[string]bool{
		"tcp": true, "TCP": true, "UDP": true, "udp": true,
	}
)

type Listener interface {
	Accept() (Conn, error)

	Close() error

	Addr() Addr

	init() error
}

// Listen case: tcp,127.0.0.1:8080
func Listen(protocol, address string) (l Listener, err error) {
	host, port, err := addrHandle(protocol, address)
	if err != nil {
		return
	}

	switch protocol {
	case "tcp":
		t, err := strconv.Atoi(port)
		if err != nil {
			return nil, err
		}
		l = NewTCPListener(NewTCPAddr("tcp", t, host))
	case "udp":
		l = nil
	case "unix":
		l = nil
	}

	if err = l.init(); err != nil {
		return
	}

	return
}
