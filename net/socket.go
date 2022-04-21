// @program:     webserver
// @file:        socket.go
// @author:      edte
// @create:      2022-04-10 17:22
// @description:
package net

import (
	"errors"
	"net"
	"syscall"
)

func createSocket(pro string, host string, port int, backlog int) (fd int, err error) {
	if pro == "tcp" || pro == "TCP" {
		fd, err = syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	} else {
		return 0, errors.New("not supported protocol")
	}
	if err != nil {
		return 0, err
	}
	addr := syscall.SockaddrInet4{
		Port: port,
	}
	copy(addr.Addr[:], net.ParseIP(host))

	if err = syscall.Bind(fd, &addr); err != nil {
		return 0, err
	}

	if err = syscall.Listen(fd, backlog); err != nil {
		return 0, err
	}

	return fd, nil
}

func dialSocket(pro string, host string, port int) (fd int, err error) {
	if pro == "tcp" || pro == "TCP" {
		fd, err = syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	} else {
		return 0, errors.New("not supported protocol")
	}
	if err != nil {
		return
	}
	addr := syscall.SockaddrInet4{
		Port: port,
	}
	copy(addr.Addr[:], net.ParseIP(host))

	if err := syscall.Connect(fd, &addr); err != nil {
		return 0, err
	}

	return fd, nil
}
