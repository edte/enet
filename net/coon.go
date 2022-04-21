// @program:     webserver
// @file:        coon.go
// @author:      edte
// @create:      2022-04-10 17:35
// @description:
package net

// Conn 连接处理
type Conn interface {
	Read(b []byte) (n int, err error)
	Write(b []byte) (n int, err error)
	Close() error
	LocalAddr() Addr
	RemoteAddr() Addr
}
