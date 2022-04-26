package main

import (
	"fmt"
	"github.com/edte/enet/net"
)

func main() {
	socket, err := net.NewSocket(net.WithSocketLocalAddr(net.NewAddr("tcp", "127.0.0.1", 1234)))
	if err != nil {
		panic(err)
	}
	if err = socket.Listen(); err != nil {
		panic(err)
	}
	for {
		s, err := socket.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println(s.RemoteAddr())
		fmt.Println(s.LocalAddr())
	}
}
