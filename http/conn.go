// @program:     enet
// @file:        conn.go
// @author:      edte
// @create:      2022-04-21 22:06
// @description:
package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/edte/enet/net"
)

// http 连接
type conn struct {
	server *Server
	netc   net.Conn

	readeBuffer bytes.Buffer
	writeBuffer bytes.Buffer
}

func newConn(server *Server, netc net.Conn) *conn {
	return &conn{
		server:      server,
		netc:        netc,
		readeBuffer: bytes.Buffer{},
		writeBuffer: bytes.Buffer{},
	}
}

func (c *conn) Start() {
	//go func() {
	//	for {
	//		_, err := c.readeBuffer.ReadFrom(c.netc)
	//		if err != nil {
	//			panic(err)
	//		}
	//	}
	//}()

	for {
		b := c.readeBuffer.Bytes()
		if len(b) < 0 {
			continue
		}

		r := &Request{}
		if err := json.NewDecoder(c.netc).Decode(&r); err != nil {
			fmt.Println(string(b))
			panic(err)
		}

		fmt.Println(r)
		c.readeBuffer.Reset()

		//req, err := c.readRequest()
		//if err != nil {
		//	panic(err)
		//}
		//resp, err := c.handleRequest(req)
		//if err != nil {
		//	panic(err)
		//}
		//if err = c.writeResponse(resp); err != nil {
		//	panic(err)
		//}
	}
}

func (c *conn) readRequest() (req *Request, err error) {
	r := &Request{}
	if err = json.Unmarshal(c.readeBuffer.Bytes(), r); err != nil {
		return nil, err
	}

	c.readeBuffer.Reset()

	//r, err := NewRequest("1", "2", nil)
	//if err != nil {
	//	return nil, err
	//}
	//r.Host = c.netc.RemoteAddr().Host()
	return r, err
}

func (c *conn) handleRequest(req *Request) (resp *Response, err error) {
	defaultHandler[req.Host](req, resp)
	return
}

func (c *conn) writeResponse(resp *Response) (err error) {
	//c.netc.Write()
	return nil
}
