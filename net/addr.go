// @program:     enet
// @file:        addr.go
// @author:      edte
// @create:      2022-04-21 19:34
// @description:
package net

import (
	"errors"
	"net"
	"strings"
	"syscall"
)

// Addr 地址处理
type Addr interface {
	Protocol() string
	Host() string
	Port() int
	String() string
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
