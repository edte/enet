package main

import (
	"fmt"
	"net"
)

var (
	addr = "127.0.0.1:9091"
)

func main() {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("dial err", err)
		return
	}
	conn.Write([]byte("hello world"))
}
