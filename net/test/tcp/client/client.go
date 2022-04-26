package main

import (
	"github.com/edte/enet/net"
	"syscall"
)

func main() {
	//conn, err := net.Dial("tcp", "127.0.0.1:1234")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(conn.LocalAddr())
	//fmt.Println(conn.RemoteAddr())

	socket, err := net.NewSocket("tcp", "127.0.0.1", 1234)
	if err != nil {
		panic(err)
	}
	if err = socket.Dial(); err != nil {
		panic(err)
	}

	_, err = syscall.Write(socket.FD(), []byte("hello world"))
	if err != nil {
		panic(err)
	}

}
