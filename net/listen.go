// @program:     webserver
// @file:        listen.go
// @author:      edte
// @create:      2022-04-10 17:16
// @description:
package net

import "net"

type Addr interface {
	Network() string // name of the network (for example, "tcp", "udp")
	String() string  // string form of address (for example, "192.0.2.1:25", "[2001:db8::1]:80")
}

type IP []byte

type TCPAddr struct {
	IP   IP
	Port int
	Zone string // IPv6 scoped addressing zone
}

type Listener interface {
	Accept() (Conn, error)

	Close() error

	Addr() Addr
}

type TCPListener struct {
	network string
	address string
}

func NewTCPListener(network string, address string) *TCPListener {
	return &TCPListener{network: network, address: address}
}

func (T TCPListener) Accept() (Conn, error) {
	//TODO implement me
	panic("implement me")
}

func (T TCPListener) Close() error {
	//TODO implement me
	panic("implement me")
}

func (T TCPListener) Addr() Addr {
	//TODO implement me
	panic("implement me")
}

// Listen case: tcp,127.0.0.1:8080
func Listen(network, address string) (Listener, error) {
	var l Listener

	if err := check(network, address); err != nil {
		return nil, err
	}

	switch network {
	case "tcp":
		l = NewTCPListener(network, address)
	case "udp":
		l = nil
	}

	return l, nil
}

func check(network, address string) error {
	net.Listen()
}
