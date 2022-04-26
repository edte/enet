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

// http://tianyu-code.top/Linux%E7%BD%91%E7%BB%9C%E7%BC%96%E7%A8%8B/Linux%E5%9F%BA%E6%9C%AC%E5%A5%97%E6%8E%A5%E5%AD%97%E7%BC%96%E7%A8%8B/

type SocketOption func(socket *Socket) error

type Socket struct {
	sfd     int // socket fd
	backlog int // 全连接队列长度

	fd          int // conn fd
	localAdder  Addr
	remoteAdder Addr
}

func NewSocket(opts ...SocketOption) (s *Socket, err error) {
	s = &Socket{}

	fd, err := createSocket("tcp")
	if err != nil {
		return nil, err
	}
	for _, opt := range opts {
		if err = opt(s); err != nil {
			return nil, err
		}
	}

	s.sfd = fd
	s.fd = fd

	return s, nil
}

func WithSocketLocalAddr(addr *addr) SocketOption {
	return func(s *Socket) error {
		s.localAdder = addr
		return nil
	}
}

func WithSocketRemoteAddr(addr *addr) SocketOption {
	return func(s *Socket) error {
		s.remoteAdder = addr
		return nil
	}
}

func WithSocketBlock(flag bool) SocketOption {
	return func(s *Socket) error {
		return syscall.SetNonblock(s.sfd, flag)
	}
}

func WithSocketReUseAddr() SocketOption {
	return func(s *Socket) error {
		return setSocket(s.sfd, syscall.SO_REUSEADDR, 4)
	}
}

func WithSocketReUsePort() SocketOption {
	return func(s *Socket) error {
		return setSocket(s.sfd, 0xf, 4)
	}
}

func WithSocketBacklog(b int) SocketOption {
	return func(s *Socket) error {
		s.backlog = b
		return nil
	}
}

func (s *Socket) bind() error {
	return syscall.Bind(s.sfd, s.addr(s.localAdder))
}

func (s *Socket) addr(a Addr) *syscall.SockaddrInet4 {
	addr := syscall.SockaddrInet4{
		Port: a.Port(),
	}
	copy(addr.Addr[:], net.ParseIP(a.Host()))
	return &addr
}

func (s *Socket) listen() error {
	return syscall.Listen(s.sfd, s.backlog)
}

func (s *Socket) connect() error {
	return syscall.Connect(s.sfd, s.addr(s.remoteAdder))
}

func (s *Socket) accept() (fd int, err error) {
	fd, _, err = syscall.Accept(s.sfd)
	return
}

func (s *Socket) Close() error {
	return syscall.Close(s.sfd)
}

func (s *Socket) LocalAddr() Addr {
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

func (s *Socket) Accept() (socket *Socket, err error) {
	fd, err := s.accept()
	if err != nil {
		return nil, err
	}

	addr, err := syscall.Getpeername(fd)
	if err != nil {
		return nil, err
	}

	res := &Socket{
		sfd:         s.sfd,
		backlog:     s.backlog,
		fd:          fd,
		localAdder:  s.localAdder,
		remoteAdder: toAddr(addr),
	}

	return res, nil
}

func (s *Socket) Dial() error {
	if err := s.connect(); err != nil {
		return err
	}

	a, err := syscall.Getsockname(s.sfd)
	if err != nil {
		return err
	}

	s.localAdder = toAddr(a)

	return nil
}

func (s *Socket) FD() int {
	return s.sfd
}

func (s *Socket) Read(b []byte) (n int, err error) {
	return syscall.Read(s.fd, b)
}

func (s *Socket) Write(b []byte) (n int, err error) {
	return syscall.Write(s.fd, b)
}

// new a socket
func createSocket(pro string) (fd int, err error) {
	if pro == "tcp" || pro == "TCP" {
		return syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	}

	if pro == "udp" || pro == "UDP" {
		return syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_UDP)
	}

	return 0, errors.New("not supported protocol")
}

func setSocket(fd int, opt int, val int) error {
	return syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, opt, val)
}
