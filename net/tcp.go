// @program:     enet
// @file:        tcp.go
// @author:      edte
// @create:      2022-04-21 18:49
// @description:
package net

import (
	"strconv"
	"syscall"
)

var (
	// 默认 TCP 全连接队列长度
	defaultTCPBacklog = 100
)

// TCPAddr 表示 tcp 地址
type TCPAddr struct {
	protocol string
	port     int
	host     string
}

func NewTCPAddr(protocol string, port int, host string) *TCPAddr {
	return &TCPAddr{protocol: protocol, port: port, host: host}
}

func (T *TCPAddr) Protocol() string {
	return T.protocol
}

func (T *TCPAddr) Host() string {
	return T.host
}

func (T *TCPAddr) Port() int {
	return T.port
}

func (T *TCPAddr) String() string {
	return T.host + strconv.Itoa(T.port)
}

// TCPListener TCP 监听器
type TCPListener struct {
	localAddr  *TCPAddr
	remoteAddr *TCPAddr
	backlog    int
	fd         int
}

func NewTCPListener(localAddr *TCPAddr) *TCPListener {
	return &TCPListener{
		localAddr: localAddr,
		backlog:   defaultTCPBacklog,
	}
}

func (T *TCPListener) init() error {
	fd, err := createSocket(T.localAddr.protocol, T.localAddr.host, T.localAddr.port, T.backlog)
	if err != nil {
		return err
	}
	T.fd = fd
	return nil
}

func (T *TCPListener) Accept() (Conn, error) {
	fd, addr, err := syscall.Accept(T.fd)
	if err != nil {
		return nil, err
	}
	T.remoteAddr = toAddr(addr).(*TCPAddr)
	return NewTCPCoon(T.localAddr, T.remoteAddr, fd), nil
}

func (T *TCPListener) Close() error {
	if err := syscall.Close(T.fd); err != nil {
		return err
	}
	return nil
}

func (T *TCPListener) Addr() Addr {
	return T.localAddr
}

// TCPCoon tcp 连接
type TCPCoon struct {
	localAddr  Addr
	remoteAddr Addr
	fd         int
}

func NewTCPCoon(localAddr Addr, remoteAddr Addr, fd int) *TCPCoon {
	return &TCPCoon{localAddr: localAddr, remoteAddr: remoteAddr, fd: fd}
}

func (T *TCPCoon) Read(b []byte) (n int, err error) {
	return syscall.Read(T.fd, b)
}

func (T *TCPCoon) Write(b []byte) (n int, err error) {
	return syscall.Write(T.fd, b)
}

func (T *TCPCoon) Close() error {
	return syscall.Close(T.fd)
}

func (T *TCPCoon) LocalAddr() Addr {
	return T.localAddr
}

func (T *TCPCoon) RemoteAddr() Addr {
	return T.remoteAddr
}

// TCPDial tcp 拨号器
type TCPDial struct {
	localAddr  *TCPAddr
	remoteAddr *TCPAddr
	fd         int
}

func NewTCPDial(remoteAddr *TCPAddr) *TCPDial {
	return &TCPDial{remoteAddr: remoteAddr}
}

func (T *TCPDial) Dail() (conn Conn, err error) {
	fd, err := dialSocket(T.remoteAddr.protocol, T.remoteAddr.host, T.remoteAddr.port)
	if err != nil {
		return nil, err
	}
	T.fd = fd
	// 获取本地地址
	name, err := syscall.Getsockname(fd)
	if err != nil {
		return nil, err
	}
	T.localAddr = toAddr(name).(*TCPAddr)
	return NewTCPCoon(T.localAddr, T.remoteAddr, T.fd), nil
}
