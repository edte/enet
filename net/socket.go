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

type SocketOption func(socket *Socket) error

type Socket struct {
	fd      int
	backlog int // 全连接队列长度

	localAdder  Addr
	remoteAdder Addr
}

func NewSocket(protocol string, ip string, port int, opts ...SocketOption) (s *Socket, err error) {
	s = &Socket{
		fd: 0,
		remoteAdder: &addr{
			host:     ip,
			port:     port,
			protocol: protocol,
		},
	}

	fd, err := s.createSocket()
	if err != nil {
		return nil, err
	}

	s.fd = fd

	for _, opt := range opts {
		if err = opt(s); err != nil {
			return nil, err
		}
	}

	return s, nil
}

func WithSocketBlock(flag bool) SocketOption {
	return func(s *Socket) error {
		return syscall.SetNonblock(s.fd, flag)
	}
}

func WithSocketReUseAddr() SocketOption {
	return func(s *Socket) error {
		return s.setSocket(syscall.SO_REUSEADDR, 4)
	}
}

func WithSocketReUsePort() SocketOption {
	return func(s *Socket) error {
		return s.setSocket(0xf, 4)
	}
}

func WithSocketBacklog(b int) SocketOption {
	return func(s *Socket) error {
		s.backlog = b
		return nil
	}
}

func (s *Socket) bind() error {
	return syscall.Bind(s.fd, s.addr())
}

func (s *Socket) addr() *syscall.SockaddrInet4 {
	addr := syscall.SockaddrInet4{
		Port: s.remoteAdder.Port(),
	}
	copy(addr.Addr[:], net.ParseIP(s.remoteAdder.Host()))
	return &addr
}

func (s *Socket) listen() error {
	return syscall.Listen(s.fd, s.backlog)
}

func (s *Socket) connect() error {
	return syscall.Connect(s.fd, s.addr())
}

func (s *Socket) accept() (fd int, err error) {
	fd, _, err = syscall.Accept(s.fd)
	return
}

func (s *Socket) Close() error {
	return syscall.Close(s.fd)
}

func (s *Socket) createSocket() (fd int, err error) {
	if s.remoteAdder.Protocol() == "tcp" || s.remoteAdder.Protocol() == "TCP" {
		return syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	}

	if s.remoteAdder.Protocol() == "udp" || s.remoteAdder.Protocol() == "UDP" {
		return syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_UDP)
	}

	return 0, errors.New("not supported protocol")
}

func (s *Socket) setSocket(opt int, val int) error {
	return syscall.SetsockoptInt(s.fd, syscall.SOL_SOCKET, opt, val)
}

func (s *Socket) LocalAddr() Addr {
	//syscall.Getsockname()
	//syscall.Getpeername()

	return s.localAdder
}

func (s *Socket) RemoteAddr() Addr {
	return s.remoteAdder
}

func (s *Socket) Listen() (err error) {
	if err = s.bind(); err != nil {
		return
	}
	if err = s.listen(); err != nil {
		return
	}
	return nil
}

func (s *Socket) Accept() (fd int, err error) {
	return s.accept()
}

func (s *Socket) Dial() error {
	return s.connect()
}

func (s *Socket) FD() int {
	return s.fd
}
