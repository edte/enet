// @program:     enet
// @file:        coon.go
// @author:      edte
// @create:      2022-04-10 17:35
// @description:
package net

import "io"

// Conn 连接处理
type Conn interface {
	io.Reader
	io.Writer
	io.Closer
	LocalAddr() Addr
	RemoteAddr() Addr
}
