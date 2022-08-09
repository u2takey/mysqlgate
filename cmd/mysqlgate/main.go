package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/u2takey/mysqlgate/pkg/log"
	"github.com/u2takey/mysqlgate/pkg/server"
	"github.com/u2takey/mysqlgate/version"
)

var (
	showVersion = flag.Bool("version", false, "show version of MysqlGate")
	logLevel    = flag.String("log", "info", "set log level with debug|info|warn|error|fatal")
	listenAddr  = flag.String("addr", "0.0.0.0:3316", "proxy listen address")
	defaultDb   = flag.String("db", "root:root@tcp(127.0.0.1:3306)/?charset=utf8&parseTime=True", "default db connection string")
)

func main() {
	flag.Parse()
	if *showVersion {
		fmt.Println(version.Version())
		return
	}
	log.SetLogLevel(*logLevel)

	svr, err := server.NewServer(*listenAddr, *defaultDb)
	if err != nil {
		log.Error("msg", "init server failed", "err", err)
		os.Exit(1)
	}
	svr.Run()
}
