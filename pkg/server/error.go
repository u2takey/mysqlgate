package server

import "errors"

var (
	ErrBadConn                   = errors.New("connection was bad")
	MysqlErrorCodeUnknown uint16 = 1105
)
