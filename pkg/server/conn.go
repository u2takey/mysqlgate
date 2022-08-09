package server

import (
	"bytes"
	"context"
	"encoding/binary"
	"net"
	"runtime"
	"sync"

	"github.com/u2takey/mysqlgate/pkg/plugin"
	"github.com/u2takey/mysqlgate/pkg/sql/mysql"
)

// client <-> proxy
type ClientConn struct {
	sync.Mutex

	c     net.Conn
	proxy *Server

	capability   mysql.ClientFlag
	status       mysql.StatusFlag
	connectionId uint32
	pkg          *PacketIO

	charset string
	user    string
	db      string
	plugin  *plugin.DefaultPlugin

	salt []byte

	closed       bool
	lastInsertId int64
	affectedRows int64
	stmtId       uint32
}

var DefaultCapacity = mysql.ClientLongPassword | mysql.ClientLongFlag |
	mysql.ClientConnectWithDB | mysql.ClientProtocol41 |
	mysql.ClientTransactions | mysql.ClientSecureConn

var baseConnId uint32 = 10000

func (c *ClientConn) Handshake() error {
	if err := c.writeInitialHandshake(); err != nil {
		mLog.Error("method", "Handshake", "connectionId", c.connectionId, "msg", "send initial handshake error", "err", err.Error())
		return err
	}

	if err := c.readHandshakeResponse(); err != nil {
		mLog.Error("method", "Handshake", "connectionId", c.connectionId, "msg", "read Handshake Response error", "err", err.Error())
		return err
	}

	if err := c.writeOK(nil); err != nil {
		mLog.Error("method", "Handshake", "connectionId", c.connectionId, "msg", "write ok fail", "error", err.Error())
		return err
	}

	c.pkg.Sequence = 0
	return nil
}

func (c *ClientConn) Close() error {
	if c.closed {
		return nil
	}

	err := c.c.Close()
	c.closed = true
	return err
}

func (c *ClientConn) writeInitialHandshake() error {
	data := make([]byte, 4, 128)
	data = append(data, 10) //min version 10

	//server version[00]
	data = append(data, "5.6"...)
	data = append(data, 0)

	//connection id
	data = append(data, byte(c.connectionId), byte(c.connectionId>>8), byte(c.connectionId>>16), byte(c.connectionId>>24))
	data = append(data, c.salt[0:8]...) //auth-plugin-data-part-1
	data = append(data, 0)              //filter [00]

	//capability flag lower 2 bytes, using default capability here
	data = append(data, byte(DefaultCapacity), byte(DefaultCapacity>>8))
	data = append(data, uint8(45))                         //charset, utf-8 default
	data = append(data, byte(c.status), byte(c.status>>8)) //status

	//below 13 byte may not be used
	//capability flag upper 2 bytes, using default capability here
	data = append(data, byte(DefaultCapacity>>16), byte(DefaultCapacity>>24))
	data = append(data, 0x15)                         //filter [0x15], for wireshark dump, value is 0x15
	data = append(data, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0) //reserved 10 [00]
	data = append(data, c.salt[8:]...)                //auth-plugin-data-part-2
	data = append(data, 0)                            //filter [00]

	return c.writePacket(data)
}

func (c *ClientConn) readPacket() ([]byte, error) {
	return c.pkg.ReadPacket()
}

func (c *ClientConn) writePacket(data []byte) error {
	return c.pkg.WritePacket(data)
}

func (c *ClientConn) writePacketBatch(total, data []byte, direct bool) ([]byte, error) {
	return c.pkg.WritePacketBatch(total, data, direct)
}

func (c *ClientConn) readHandshakeResponse() error {
	data, err := c.readPacket()

	if err != nil {
		return err
	}

	pos := 0

	//capability
	c.capability = mysql.ClientFlag(binary.LittleEndian.Uint32(data[:4]))
	pos += 4

	//skip max packet size
	pos += 4

	//charset, skip, if you want to use another charset, use set names
	//c.collation = CollationId(data[pos])
	pos++

	//skip reserved 23[00]
	pos += 23

	// user name
	c.user = string(data[pos : pos+bytes.IndexByte(data[pos:], 0)])

	pos += len(c.user) + 1

	//auth length and auth
	authLen := int(data[pos])
	pos++
	//auth := data[pos : pos+authLen]

	//check user
	//if _, ok := c.proxy.users[c.user]; !ok {
	//	golog.Error("ClientConn", "readHandshakeResponse", "error", 0,
	//		"auth", auth,
	//		"client_user", c.user,
	//		"config_set_user", c.user,
	//		"password", c.proxy.users[c.user])
	//	return mysql.NewDefaultError(mysql.ER_ACCESS_DENIED_ERROR, c.user, c.c.RemoteAddr().String(), "Yes")
	//}

	//check password
	//checkAuth := mysql.CalcPassword(c.salt, []byte(c.proxy.users[c.user]))
	//if !bytes.Equal(auth, checkAuth) {
	//	golog.Error("ClientConn", "readHandshakeResponse", "error", 0,
	//		"auth", auth,
	//		"checkAuth", checkAuth,
	//		"client_user", c.user,
	//		"config_set_user", c.user,
	//		"password", c.proxy.users[c.user])
	//	return mysql.NewDefaultError(mysql.ER_ACCESS_DENIED_ERROR, c.user, c.c.RemoteAddr().String(), "Yes")
	//}

	pos += authLen

	var db string
	if c.capability&mysql.ClientConnectWithDB > 0 {
		if len(data[pos:]) == 0 {
			return nil
		}

		db = string(data[pos : pos+bytes.IndexByte(data[pos:], 0)])
		pos += len(c.db) + 1

	}
	c.db = db

	return nil
}

func (c *ClientConn) clean() {

}

func (c *ClientConn) Run() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			const size = 4096
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]

			mLog.Error("method", "Run", "stack", string(buf), "err", err.Error())
		}

		c.Close()
	}()
	defer c.clean()
	for {
		data, err := c.readPacket()
		if err != nil {
			return
		}

		if err := c.dispatch(data); err != nil {
			mLog.Error("method", "Run", "connectionId", c.connectionId, "err", err.Error())
			c.writeError(err)
			if err == ErrBadConn {
				c.Close()
			}
		}

		if c.closed {
			return
		}

		c.pkg.Sequence = 0
	}
}

func (c *ClientConn) dispatch(data []byte) error {
	cmd := data[0]
	data = data[1:]

	switch cmd {
	case mysql.ComQuit:
		c.Close()
		return nil
	default:
		c.plugin.ExecContext(context.Background(), string(data))
	}

	return nil
}

func (c *ClientConn) writeOK(r *mysql.MysqlResult) error {
	if r == nil {
		r = &mysql.MysqlResult{}
	}
	data := make([]byte, 4, 32)

	data = append(data, mysql.IOK)

	data = append(data, PutLengthEncodedInt(uint64(r.AffectedRows))...)
	data = append(data, PutLengthEncodedInt(uint64(r.InsertId))...)

	if c.capability&mysql.ClientProtocol41 > 0 {
		data = append(data, byte(c.status), byte(c.status>>8))
		data = append(data, 0, 0)
	}

	return c.writePacket(data)
}

func (c *ClientConn) writeError(e error) error {
	var m *mysql.MySQLError
	var ok bool
	if m, ok = e.(*mysql.MySQLError); !ok {
		m = &mysql.MySQLError{Number: MysqlErrorCodeUnknown, Message: e.Error()}
	}

	data := make([]byte, 4, 16+len(m.Message))

	data = append(data, mysql.IERR)
	data = append(data, byte(m.Number), byte(m.Number>>8))

	if c.capability&mysql.ClientProtocol41 > 0 {
		data = append(data, '#')
		//data = append(data, m.State...)
	}

	data = append(data, m.Message...)
	return c.writePacket(data)
}

func (c *ClientConn) writeEOF(status uint16) error {
	data := make([]byte, 4, 9)

	data = append(data, mysql.IEOF)
	if c.capability&mysql.ClientProtocol41 > 0 {
		data = append(data, 0, 0)
		data = append(data, byte(status), byte(status>>8))
	}

	return c.writePacket(data)
}

func (c *ClientConn) writeEOFBatch(total []byte, status uint16, direct bool) ([]byte, error) {
	data := make([]byte, 4, 9)

	data = append(data, mysql.IEOF)
	if c.capability&mysql.ClientProtocol41 > 0 {
		data = append(data, 0, 0)
		data = append(data, byte(status), byte(status>>8))
	}

	return c.writePacketBatch(total, data, direct)
}
