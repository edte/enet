package main

import (
	"fmt"
	"github.com/edte/enet/net"
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

		b := make([]byte, 10000)

		_, err = ac.Read(b)
		fmt.Println(string(b))
	}

}
