package mysql

import (
	"context"
	"net"

	"github.com/u2takey/mysqlgate/pkg/log"
)

var mLog = log.ModuleLogger("server.mysql")

type MysqlConnector struct {
	cfg *Config
}

func NewConnector(cfg *Config) *MysqlConnector {
	return &MysqlConnector{cfg: cfg}
}

func (c *MysqlConnector) OnConnect(ctx context.Context, conn net.Conn) (Conn, error) {
	m := &MysqlConn{
		netConn:          conn,
		maxAllowedPacket: maxPacketSize,
		maxWriteSize:     maxPacketSize - 1,
		cfg:              c.cfg,
		buf:              newBuffer(conn),
	}
	return m, m.handshake(ctx)
}
