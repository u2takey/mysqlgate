package server

import (
	"fmt"
	"net"
	"runtime"

	"github.com/u2takey/mysqlgate/pkg/log"
	"github.com/u2takey/mysqlgate/pkg/sql"
	_ "github.com/u2takey/mysqlgate/pkg/sql/mysql"
)

var mLog = log.ModuleLogger("server")

type Server struct {
	listenAddr    string
	defaultDbAddr string
	db            *sql.DB

	listener net.Listener
}

func NewServer(addr, defaultDbAddr string) (*Server, error) {
	var err error
	s := &Server{
		listenAddr:    addr,
		defaultDbAddr: defaultDbAddr,
	}
	s.db, err = sql.Open("mysql", defaultDbAddr)
	if err != nil {
		return nil, err
	}
	s.listener, err = net.Listen("tcp", addr)
	return s, err
}

func (s *Server) Run() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			mLog.Error("method", "Run", "msg", "accept failed", "err", err.Error())
			continue
		}
		fmt.Println(conn)
	}
}

func (s *Server) onConn(c net.Conn) {
	conn := s.newClientConn(c) //新建一个conn

	defer func() {
		err := recover()
		if err != nil {
			const size = 4096
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)] //获得当前goroutine的stacktrace
			golog.Error("server", "onConn", "error", 0,
				"remoteAddr", c.RemoteAddr().String(),
				"stack", string(buf),
			)
		}

		conn.Close()
	}()

	if allowConnect := conn.IsAllowConnect(); allowConnect == false {
		err := mysql.NewError(mysql.ER_ACCESS_DENIED_ERROR, "ip address access denied by kingshard.")
		conn.writeError(err)
		conn.Close()
		return
	}
	if err := conn.Handshake(); err != nil {
		golog.Error("server", "onConn", err.Error(), 0)
		conn.writeError(err)
		conn.Close()
		return
	}

	conn.schema = s.GetSchema(conn.user)

	conn.Run()
}
