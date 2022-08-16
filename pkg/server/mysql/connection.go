package mysql

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/go-sql-driver/mysql"
)

type MysqlConn struct {
	buf          buffer
	netConn      net.Conn // underlying connection
	connectionId uint32
	status       StatusFlag
	sequence     uint8
	capability   ClientFlag
	database     string

	// config
	cfg              *Config
	maxAllowedPacket int
	maxWriteSize     int
	writeTimeout     time.Duration

	// backend
	plan QueryPlan
}

func (mc *MysqlConn) handshake(ctx context.Context) error {
	if err := mc.writeHandshakePacket(); err != nil {
		return err
	}
	if err := mc.readHandshakeResponse(); err != nil {
		mc.writeError(err)
		return err
	}
	if err := mc.writeOK(nil); err != nil {
		mc.writeError(err)
		return err
	}
	mc.sequence = 0
	return nil
}

func (mc *MysqlConn) Run(ctx context.Context, plan QueryPlan) error {
	//defer func() {
	//	r := recover()
	//	if err, ok := r.(error); ok {
	//		const size = 4096
	//		buf := make([]byte, size)
	//		buf = buf[:runtime.Stack(buf, false)]
	//
	//		mLog.Error("method", "Run", "err", err.Error(), "stack", string(buf))
	//	}
	//	mc.cleanup()
	//}()

	mc.plan = plan

	for {
		select {
		case <-ctx.Done():
			mc.cleanup()
			return ctx.Err()
		default:
			data, err := mc.readPacket()
			if err != nil {
				return err
			}
			cmd := data[0]
			data = data[1:]

			switch cmd {
			case ComQuit:
				// todo
				err = mc.writeOK(nil)
			case ComQuery:
				err = mc.plan.Query(NewQueryContext(ctx, mc), string(data))
			case ComPing:
				err = mc.writeOK(nil)
			case ComSetOption:
				err = mc.writeEOF(0)
			case ComInitDB:
				// todo pick db and call use db
				mc.database = string(data)
				err = mc.writeOK(nil)
			case ComFieldList:
				fallthrough
			case ComStmtPrepare:
				fallthrough
			case ComStmtExecute:
				fallthrough
			case ComStmtClose:
				fallthrough
			case ComStmtSendLongData:
				fallthrough
			case ComStmtReset:
				fallthrough
			default:
				msg := fmt.Sprintf("command %d not supported now", cmd)
				mLog.Error("method", "Run", "msg", msg)
				err = NewCustomError(ErUnknownError, msg)
			}
			if err != nil {
				mc.writeError(err)
			}
			if err == mysql.ErrInvalidConn {
				mc.cleanup()
				return err
			}
			mc.sequence = 0
		}
	}
}

func (mc *MysqlConn) cleanup() {

}
