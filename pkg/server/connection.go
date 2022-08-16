package server

import "net"

type ClientConn struct {
}

func NewClientConn(conn net.Conn) *ClientConn {
	return &ClientConn{}
}

func (c *ClientConn) Close() {

}
