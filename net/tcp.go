// @program:     enet
// @file:        tcp.go
// @author:      edte
// @create:      2022-04-21 18:49
// @description:
package net

// TCPListener TCP 监听器
type TCPListener struct {
	socket    *Socket
	localAddr Addr
}

func NewTCPListener(a *addr) (Listener, error) {
	l := &TCPListener{
		localAddr: a,
	}

	socket, err := NewSocket(WithSocketLocalAddr(a))
	if err != nil {
		return nil, err
	}

	if err = socket.Listen(); err != nil {
		return nil, err
	}

	l.socket = socket

	return l, nil
}

func (T *TCPListener) Accept() (Conn, error) {
	socket, err := T.socket.Accept()
	if err != nil {
		return nil, err
	}

	return NewTCPCoon(socket), nil
}

func (T *TCPListener) Close() error {
	return T.socket.Close()
}

func (T *TCPListener) Addr() Addr {
	return T.localAddr
}

// TCPCoon tcp 连接
type TCPCoon struct {
	socket *Socket
}

func NewTCPCoon(s *Socket) *TCPCoon {
	return &TCPCoon{socket: s}
}

func (T *TCPCoon) Read(b []byte) (n int, err error) {
	return T.socket.Read(b)
}

func (T *TCPCoon) Write(b []byte) (n int, err error) {
	return T.socket.Write(b)
}

func (T *TCPCoon) Close() error {
	return T.socket.Close()
}

func (T *TCPCoon) LocalAddr() Addr {
	return T.socket.LocalAddr()
}

func (T *TCPCoon) RemoteAddr() Addr {
	return T.socket.RemoteAddr()
}

// TCPDial tcp 拨号器
type TCPDial struct {
	socket     *Socket
	remoteAddr *addr
}

func NewTCPDial(remoteAddr *addr) *TCPDial {
	return &TCPDial{remoteAddr: remoteAddr}
}

func (T *TCPDial) Dail() (conn Conn, err error) {
	s, err := NewSocket(WithSocketRemoteAddr(T.remoteAddr))
	if err != nil {
		return nil, err
	}
	T.socket = s

	if err = s.Dial(); err != nil {
		return nil, err
	}

	return NewTCPCoon(s), nil
}
