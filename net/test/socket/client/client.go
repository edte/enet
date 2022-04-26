package main

import (
	"fmt"
	"github.com/edte/enet/net"
)

func main() {
	socket, err := net.NewSocket(net.WithSocketRemoteAddr(net.NewAddr("tcp", "127.0.0.1", 1234)))
	if err != nil {
		panic(err)
	}
	if err = socket.Dial(); err != nil {
		panic(err)
	}

	fmt.Println(socket.RemoteAddr())
	fmt.Println(socket.LocalAddr())
}
