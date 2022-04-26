package main

import (
	"fmt"
	"github.com/edte/enet/net"
	"syscall"
)

func main() {
	//l, err := net.Listen("tcp", "127.0.0.1:1234")
	//if err != nil {
	//	panic(err)
	//}
	//for {
	//	ac, err := l.Accept()
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(ac.LocalAddr())
	//	fmt.Println(ac.RemoteAddr())
	//}

	socket, err := net.NewSocket("tcp", "127.0.0.1", 1234)
	if err != nil {
		panic(err)
	}
	if err = socket.Listen(); err != nil {
		panic(err)
	}
	for {
		fd, err := socket.Accept()
		if err != nil {
			panic(err)
		}
		data := make([]byte, 10000)
		_, err = syscall.Read(fd, data)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(data))
	}

}
