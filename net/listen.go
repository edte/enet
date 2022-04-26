// @program:     enet
// @file:        listen.go
// @author:      edte
// @create:      2022-04-10 17:16
// @description:
package net

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
}

// Listen case: tcp,127.0.0.1:8080
func Listen(protocol, address string) (l Listener, err error) {
	addr, err := addrHandle(protocol, address)
	if err != nil {
		return
	}

	switch protocol {
	case "tcp":
		l, err = NewTCPListener(addr)
	case "udp":
		l = nil
	case "unix":
		l = nil
	}

	return
}
