package server

import (
	"context"
	"crypto/rand"
	"net"

	"github.com/u2takey/mysqlgate/pkg/log"
	"github.com/u2takey/mysqlgate/pkg/server/mysql"
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

func (s *Server) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			conn, err := s.listener.Accept()
			if err != nil {
				mLog.Error("method", "Run", "msg", "accept failed", "err", err.Error())
				continue
			}
			go s.onConn(conn)
		}
	}
}

func (s *Server) onConn(c net.Conn) {
	cfg := mysql.NewConfig()
	cfg.User, cfg.Passwd = "root", "root"
	cfg.Salt = make([]byte, 20)
	_, _ = rand.Read(cfg.Salt)
	connector := mysql.NewConnector(cfg)

	//defer func() {
	//	err := recover()
	//	if err != nil {
	//		const size = 4096
	//		buf := make([]byte, size)
	//		buf = buf[:runtime.Stack(buf, false)] //获得当前goroutine的stacktrace
	//		mLog.Error("server", "onConn", "error", 0,
	//			"remoteAddr", c.RemoteAddr().String(),
	//			"stack", string(buf),
	//		)
	//	}
	//}()

	conn, err := connector.OnConnect(context.Background(), c)
	if err != nil {
		mLog.Error("method", "onConn", "msg", "on connect failed", "err", err.Error())
		return
	}
	mLog.Debug("method", "onConn", "msg", "connect success")
	err = conn.Run(mysql.NewQueryContext(context.Background(), s.db))
	if err != nil {
		mLog.Error("method", "onConn", "err", err.Error(), "msg", "conn break")
	}
}
