// @program:     enet
// @file:        dail.go
// @author:      edte
// @create:      2022-04-10 17:16
// @description:
package net

type dailer interface {
	Dail() (conn Conn, err error)
}

func Dial(protocol, address string) (conn Conn, err error) {
	addr, err := addrHandle(protocol, address)
	if err != nil {
		return
	}

	var d dailer

	switch protocol {
	case "tcp":
		d = NewTCPDial(addr)
	case "udp":
		d = nil
	case "unix":
		d = nil
	}

	return d.Dail()
}
