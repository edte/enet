package main

import (
	"fmt"
	"webserver/net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		panic(err)
	}
	fmt.Println(conn.LocalAddr())
	fmt.Println(conn.RemoteAddr())
}
