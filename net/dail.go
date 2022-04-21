// @program:     webserver
// @file:        dail.go
// @author:      edte
// @create:      2022-04-10 17:16
// @description:
package net

import (
	"strconv"
)

type dailer interface {
	Dail() (conn Conn, err error)
}

func Dial(protocol, address string) (conn Conn, err error) {
	host, port, err := addrHandle(protocol, address)
	if err != nil {
		return
	}

	var d dailer

	switch protocol {
	case "tcp":
		t, err := strconv.Atoi(port)
		if err != nil {
			return nil, err
		}
		d = NewTCPDial(NewTCPAddr("tcp", t, host))
	case "udp":
		d = nil
	case "unix":
		d = nil
	}

	return d.Dail()
}
