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

func (a addr) Protocol() string {
	return a.protocol
}

func (a addr) Host() string {
	return a.host
}

func (a addr) Port() int {
	return a.port
}

func (a addr) String() string {
	return a.host + strconv.Itoa(a.port)
}

func toAddr(sa syscall.Sockaddr) Addr {
	i := sa.(*syscall.SockaddrInet4)
	n := net.IPAddr{
		IP: i.Addr[:],
	}
	return NewTCPAddr("tcp", i.Port, n.String())
}

func addrHandle(protocol, address string) (host, port string, err error) {
	if !validProtocol[protocol] {
		return "", "", errors.New("not support protocol")
	}
	s := strings.Split(address, ":")
	return s[0], s[1], nil
}
