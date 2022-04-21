package main

import (
	"fmt"
	"webserver/net"
)

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		panic(err)
	}
	for {
		ac, err := l.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println(ac.LocalAddr())
		fmt.Println(ac.RemoteAddr())
	}
}
