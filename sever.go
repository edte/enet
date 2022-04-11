package main

import (
	"fmt"
	"net"
	"strconv"
	"syscall"
)

func main() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		panic(err)
	}

	addr := syscall.SockaddrInet4{
		Port: 1234,
	}
	copy(addr.Addr[:], net.ParseIP("127.0.0.1"))

	err = syscall.Bind(fd, &addr)
	if err != nil {
		panic(err)
	}

	err = syscall.Listen(fd, 5)
	if err != nil {
		panic(err)
	}

	for {
		f, _, err := syscall.Accept(fd)

		if err != nil {
			panic(err)
		}

		name, err := syscall.Getpeername(f)
		if err != nil {
			panic(err)
		}
		fmt.Println(toAddr(name))

		_, err = syscall.Write(f, []byte("hello world"))
		if err != nil {
			panic(err)
		}

	}
}

type Addr interface {
	Network() string // name of the network (for example, "tcp", "udp")
	String() string  // string form of address (for example, "192.0.2.1:25", "[2001:db8::1]:80")
}

type TCPAddr struct {
	IP   IP
	Port int
	Zone string // IPv6 scoped addressing zone
}

type IP []byte

func (T TCPAddr) Network() string {
	return string(T.IP)
}

func (T TCPAddr) String() string {
	addr := net.IPAddr{
		IP: net.IP(T.IP),
	}
	return addr.String() + strconv.Itoa(T.Port)
}

func toAddr(sa syscall.Sockaddr) Addr {
	i := sa.(*syscall.SockaddrInet4)
	return &TCPAddr{IP: i.Addr[0:], Port: i.Port}
}
