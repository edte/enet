// @program:     enet
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

	// 设置非阻塞 socket
	// todo: tcp 默认为阻塞还是非阻塞
	//if err = syscall.SetNonblock(fd, true); err != nil {
	//	return 0, err
	//}

	// client close 后，server 会等待 2msl，调试不方便，故先直接设置 REUSE_ADDR，可以立刻开新连接
	if err = syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, -1); err != nil {
		return 0, nil
	}

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
