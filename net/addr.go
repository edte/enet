// @program:     enet
// @file:        addr.go
// @author:      edte
// @create:      2022-04-21 19:34
// @description:
package net

import (
	"errors"
	"net"
	"strconv"
	"strings"
	"syscall"
)

// Addr 地址处理
type Addr interface {
	Protocol() string // tcp、udp
	Host() string     // ip
	Port() int        // :8080
	String() string
}

type addr struct {
	host     string
	port     int
	protocol string
}

func NewAddr(protocol string, host string, port int) *addr {
	return &addr{protocol: protocol, port: port, host: host}
}

func (a *addr) Protocol() string {
	return a.protocol
}

func (a *addr) Host() string {
	return a.host
}

func (a *addr) Port() int {
	return a.port
}

func (a *addr) String() string {
	return a.host + strconv.Itoa(a.port)
}

func addrHandle(protocol, address string) (addr *addr, err error) {
	if !validProtocol[protocol] {
		return nil, errors.New("not support protocol")
	}

	s := strings.Split(address, ":")

	t, err := strconv.Atoi(s[1])
	if err != nil {
		return
	}
	addr = NewAddr(protocol, s[0], t)

	return
}

func toAddr(sa syscall.Sockaddr) Addr {
	i := sa.(*syscall.SockaddrInet4)
	n := net.IPAddr{
		IP: i.Addr[:],
	}
	return NewAddr("tcp", n.String(), i.Port)
}
