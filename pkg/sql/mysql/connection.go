// Go MySQL Driver - A MySQL-Driver for Go's database/sql package
//
// Copyright 2012 The Go-MySQL-Driver Authors. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package mysql

import (
	"context"
	"encoding/json"
	"io"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/u2takey/mysqlgate/pkg/sql"
	"github.com/u2takey/mysqlgate/pkg/sql/driver"
)

type MysqlConn struct {
	buf              buffer
	netConn          net.Conn
	rawConn          net.Conn // underlying connection when netConn is TLS connection.
	affectedRows     uint64
	insertId         uint64
	cfg              *Config
	maxAllowedPacket int
	maxWriteSize     int
	writeTimeout     time.Duration
	flags            ClientFlag
	status           StatusFlag
	sequence         uint8
	parseTime        bool
	reset            bool // set when the Go SQL package calls ResetSession
	serverVersion    string

	// for context support (Go 1.8+)
	watching bool
	watcher  chan<- context.Context
	closech  chan struct{}
	finished chan<- struct{}
	canceled atomicError // set non-nil if conn is canceled
	closed   atomicBool  // set when conn is closed, before closech is closed
}

// Handles parameters set in DSN after the connection is established
func (mc *MysqlConn) handleParams() (err error) {
	var cmdSet strings.Builder
	for param, val := range mc.cfg.Params {
		switch param {
		// Charset: character_set_connection, character_set_client, character_set_results
		case "charset":
			charsets := strings.Split(val, ",")
			for i := range charsets {
				// ignore errors here - a charset may not exist
				err = mc.exec("SET NAMES " + charsets[i])
				if err == nil {
					break
				}
			}
			if err != nil {
				return
			}

		// Other system vars accumulated in a single SET command
		default:
			if cmdSet.Len() == 0 {
				// Heuristic: 29 chars for each other key=value to reduce reallocations
				cmdSet.Grow(4 + len(param) + 1 + len(val) + 30*(len(mc.cfg.Params)-1))
				cmdSet.WriteString("SET ")
			} else {
				cmdSet.WriteByte(',')
			}
			cmdSet.WriteString(param)
			cmdSet.WriteByte('=')
			cmdSet.WriteString(val)
		}
	}

	if cmdSet.Len() > 0 {
		err = mc.exec(cmdSet.String())
		if err != nil {
			return
		}
	}

	return
}

func (mc *MysqlConn) markBadConn(err error) error {
	if mc == nil {
		return err
	}
	if err != errBadConnNoWrite {
		return err
	}
	return driver.ErrBadConn
}

func (mc *MysqlConn) Begin() (driver.Tx, error) {
	return mc.begin(false)
}

func (mc *MysqlConn) begin(readOnly bool) (driver.Tx, error) {
	if mc.closed.IsSet() {
		errLog.Print(ErrInvalidConn)
		return nil, driver.ErrBadConn
	}
	var q string
	if readOnly {
		q = "START TRANSACTION READ ONLY"
	} else {
		q = "START TRANSACTION"
	}
	err := mc.exec(q)
	if err == nil {
		return &mysqlTx{mc}, err
	}
	return nil, mc.markBadConn(err)
}

func (mc *MysqlConn) Close() (err error) {
	// Makes Close idempotent
	if !mc.closed.IsSet() {
		err = mc.writeCommandPacket(ComQuit)
	}

	mc.cleanup()

	return
}

// Closes the network connection and unsets internal variables. Do not call this
// function after successfully authentication, call Close instead. This function
// is called before auth or on auth failure because MySQL will have already
// closed the network connection.
func (mc *MysqlConn) cleanup() {
	if !mc.closed.TrySet(true) {
		return
	}

	// Makes cleanup idempotent
	close(mc.closech)
	if mc.netConn == nil {
		return
	}
	if err := mc.netConn.Close(); err != nil {
		errLog.Print(err)
	}
}

func (mc *MysqlConn) error() error {
	if mc.closed.IsSet() {
		if err := mc.canceled.Value(); err != nil {
			return err
		}
		return ErrInvalidConn
	}
	return nil
}

func (mc *MysqlConn) Prepare(query string) (driver.Stmt, error) {
	if mc.closed.IsSet() {
		errLog.Print(ErrInvalidConn)
		return nil, driver.ErrBadConn
	}
	// Send command
	err := mc.writeCommandPacketStr(ComStmtPrepare, query)
	if err != nil {
		// STMT_PREPARE is safe to retry.  So we can return ErrBadConn here.
		errLog.Print(err)
		return nil, driver.ErrBadConn
	}

	stmt := &MysqlStmt{
		mc: mc,
	}

	// Read Result
	columnCount, err := stmt.readPrepareResultPacket()
	if err == nil {
		if stmt.paramCount > 0 {
			if err = mc.readUntilEOF(); err != nil {
				return nil, err
			}
		}

		if columnCount > 0 {
			err = mc.readUntilEOF()
		}
	}

	return stmt, err
}

func (mc *MysqlConn) interpolateParams(query string, args []driver.Value) (string, error) {
	// Number of ? should be same to len(args)
	if strings.Count(query, "?") != len(args) {
		return "", driver.ErrSkip
	}

	buf, err := mc.buf.takeCompleteBuffer()
	if err != nil {
		// can not take the buffer. Something must be wrong with the connection
		errLog.Print(err)
		return "", ErrInvalidConn
	}
	buf = buf[:0]
	argPos := 0

	for i := 0; i < len(query); i++ {
		q := strings.IndexByte(query[i:], '?')
		if q == -1 {
			buf = append(buf, query[i:]...)
			break
		}
		buf = append(buf, query[i:i+q]...)
		i += q

		arg := args[argPos]
		argPos++

		if arg == nil {
			buf = append(buf, "NULL"...)
			continue
		}

		switch v := arg.(type) {
		case int64:
			buf = strconv.AppendInt(buf, v, 10)
		case uint64:
			// Handle uint64 explicitly because our custom ConvertValue emits unsigned values
			buf = strconv.AppendUint(buf, v, 10)
		case float64:
			buf = strconv.AppendFloat(buf, v, 'g', -1, 64)
		case bool:
			if v {
				buf = append(buf, '1')
			} else {
				buf = append(buf, '0')
			}
		case time.Time:
			if v.IsZero() {
				buf = append(buf, "'0000-00-00'"...)
			} else {
				buf = append(buf, '\'')
				buf, err = appendDateTime(buf, v.In(mc.cfg.Loc))
				if err != nil {
					return "", err
				}
				buf = append(buf, '\'')
			}
		case json.RawMessage:
			buf = append(buf, '\'')
			if mc.status&StatusNoBackslashEscapes == 0 {
				buf = escapeBytesBackslash(buf, v)
			} else {
				buf = escapeBytesQuotes(buf, v)
			}
			buf = append(buf, '\'')
		case []byte:
			if v == nil {
				buf = append(buf, "NULL"...)
			} else {
				buf = append(buf, "_binary'"...)
				if mc.status&StatusNoBackslashEscapes == 0 {
					buf = escapeBytesBackslash(buf, v)
				} else {
					buf = escapeBytesQuotes(buf, v)
				}
				buf = append(buf, '\'')
			}
		case string:
			buf = append(buf, '\'')
			if mc.status&StatusNoBackslashEscapes == 0 {
				buf = escapeStringBackslash(buf, v)
			} else {
				buf = escapeStringQuotes(buf, v)
			}
			buf = append(buf, '\'')
		default:
			return "", driver.ErrSkip
		}

		if len(buf)+4 > mc.maxAllowedPacket {
			return "", driver.ErrSkip
		}
	}
	if argPos != len(args) {
		return "", driver.ErrSkip
	}
	return string(buf), nil
}

func (mc *MysqlConn) Exec(query string, args []driver.Value) (driver.Result, error) {
	if mc.closed.IsSet() {
		errLog.Print(ErrInvalidConn)
		return nil, driver.ErrBadConn
	}
	if len(args) != 0 {
		if !mc.cfg.InterpolateParams {
			return nil, driver.ErrSkip
		}
		// try to interpolate the parameters to save extra roundtrips for preparing and closing a statement
		prepared, err := mc.interpolateParams(query, args)
		if err != nil {
			return nil, err
		}
		query = prepared
	}
	mc.affectedRows = 0
	mc.insertId = 0

	err := mc.exec(query)
	if err == nil {
		return &MysqlResult{
			AffectedRows: int64(mc.affectedRows),
			InsertId:     int64(mc.insertId),
		}, err
	}
	return nil, mc.markBadConn(err)
}

// Internal function to execute commands
func (mc *MysqlConn) exec(query string) error {
	// Send command
	if err := mc.writeCommandPacketStr(ComQuery, query); err != nil {
		return mc.markBadConn(err)
	}

	// Read Result
	resLen, err := mc.readResultSetHeaderPacket()
	if err != nil {
		return err
	}

	if resLen > 0 {
		// columns
		if err := mc.readUntilEOF(); err != nil {
			return err
		}

		// rows
		if err := mc.readUntilEOF(); err != nil {
			return err
		}
	}

	return mc.discardResults()
}

func (mc *MysqlConn) Query(query string, args []driver.Value) (driver.Rows, error) {
	return mc.query(query, args)
}

func (mc *MysqlConn) query(query string, args []driver.Value) (*textRows, error) {
	if mc.closed.IsSet() {
		errLog.Print(ErrInvalidConn)
		return nil, driver.ErrBadConn
	}
	if len(args) != 0 {
		if !mc.cfg.InterpolateParams {
			return nil, driver.ErrSkip
		}
		// try client-side prepare to reduce roundtrip
		prepared, err := mc.interpolateParams(query, args)
		if err != nil {
			return nil, err
		}
		query = prepared
	}
	// Send command
	err := mc.writeCommandPacketStr(ComQuery, query)
	if err == nil {
		// Read Result
		var resLen int
		resLen, err = mc.readResultSetHeaderPacket()
		if err == nil {
			rows := new(textRows)
			rows.mc = mc

			if resLen == 0 {
				rows.rs.done = true

				switch err := rows.NextResultSet(); err {
				case nil, io.EOF:
					return rows, nil
				default:
					return nil, err
				}
			}

			// Columns
			rows.rs.columns, rows.rs.rawColumns, err = mc.readColumns(resLen)
			return rows, err
		}
	}
	return nil, mc.markBadConn(err)
}

// Gets the value of the given MySQL System Variable
// The returned byte slice is only valid until the next read
func (mc *MysqlConn) getSystemVar(name string) ([]byte, error) {
	// Send command
	if err := mc.writeCommandPacketStr(ComQuery, "SELECT @@"+name); err != nil {
		return nil, err
	}

	// Read Result
	resLen, err := mc.readResultSetHeaderPacket()
	if err == nil {
		rows := new(textRows)
		rows.mc = mc
		rows.rs.columns = []mysqlField{{fieldType: fieldTypeVarChar}}

		if resLen > 0 {
			// Columns
			if err := mc.readUntilEOF(); err != nil {
				return nil, err
			}
		}

		dest := make([]driver.Value, resLen)
		if err = rows.readRow(dest); err == nil {
			return dest[0].([]byte), mc.readUntilEOF()
		}
	}
	return nil, err
}

// finish is called when the query has canceled.
func (mc *MysqlConn) cancel(err error) {
	mc.canceled.Set(err)
	mc.cleanup()
}

// finish is called when the query has succeeded.
func (mc *MysqlConn) finish() {
	if !mc.watching || mc.finished == nil {
		return
	}
	select {
	case mc.finished <- struct{}{}:
		mc.watching = false
	case <-mc.closech:
	}
}

// Ping implements driver.Pinger interface
func (mc *MysqlConn) Ping(ctx context.Context) (err error) {
	if mc.closed.IsSet() {
		errLog.Print(ErrInvalidConn)
		return driver.ErrBadConn
	}

	if err = mc.watchCancel(ctx); err != nil {
		return
	}
	defer mc.finish()

	if err = mc.writeCommandPacket(ComPing); err != nil {
		return mc.markBadConn(err)
	}

	return mc.readResultOK()
}

// BeginTx implements driver.ConnBeginTx interface
func (mc *MysqlConn) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	if mc.closed.IsSet() {
		return nil, driver.ErrBadConn
	}

	if err := mc.watchCancel(ctx); err != nil {
		return nil, err
	}
	defer mc.finish()

	if sql.IsolationLevel(opts.Isolation) != sql.LevelDefault {
		level, err := mapIsolationLevel(opts.Isolation)
		if err != nil {
			return nil, err
		}
		err = mc.exec("SET TRANSACTION ISOLATION LEVEL " + level)
		if err != nil {
			return nil, err
		}
	}

	return mc.begin(opts.ReadOnly)
}

func (mc *MysqlConn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	dargs, err := namedValueToValue(args)
	if err != nil {
		return nil, err
	}

	if err := mc.watchCancel(ctx); err != nil {
		return nil, err
	}

	rows, err := mc.query(query, dargs)
	if err != nil {
		mc.finish()
		return nil, err
	}
	rows.finish = mc.finish
	return rows, err
}

func (mc *MysqlConn) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	dargs, err := namedValueToValue(args)
	if err != nil {
		return nil, err
	}

	if err := mc.watchCancel(ctx); err != nil {
		return nil, err
	}
	defer mc.finish()

	return mc.Exec(query, dargs)
}

func (mc *MysqlConn) PrepareContext(ctx context.Context, query string) (driver.Stmt, error) {
	if err := mc.watchCancel(ctx); err != nil {
		return nil, err
	}

	stmt, err := mc.Prepare(query)
	mc.finish()
	if err != nil {
		return nil, err
	}

	select {
	default:
	case <-ctx.Done():
		stmt.Close()
		return nil, ctx.Err()
	}
	return stmt, nil
}

func (stmt *MysqlStmt) QueryContext(ctx context.Context, args []driver.NamedValue) (driver.Rows, error) {
	dargs, err := namedValueToValue(args)
	if err != nil {
		return nil, err
	}

	if err := stmt.mc.watchCancel(ctx); err != nil {
		return nil, err
	}

	rows, err := stmt.query(dargs)
	if err != nil {
		stmt.mc.finish()
		return nil, err
	}
	rows.finish = stmt.mc.finish
	return rows, err
}

func (stmt *MysqlStmt) ExecContext(ctx context.Context, args []driver.NamedValue) (driver.Result, error) {
	dargs, err := namedValueToValue(args)
	if err != nil {
		return nil, err
	}

	if err := stmt.mc.watchCancel(ctx); err != nil {
		return nil, err
	}
	defer stmt.mc.finish()

	return stmt.Exec(dargs)
}

func (mc *MysqlConn) watchCancel(ctx context.Context) error {
	if mc.watching {
		// Reach here if canceled,
		// so the connection is already invalid
		mc.cleanup()
		return nil
	}
	// When ctx is already cancelled, don't watch it.
	if err := ctx.Err(); err != nil {
		return err
	}
	// When ctx is not cancellable, don't watch it.
	if ctx.Done() == nil {
		return nil
	}
	// When watcher is not alive, can't watch it.
	if mc.watcher == nil {
		return nil
	}

	mc.watching = true
	mc.watcher <- ctx
	return nil
}

func (mc *MysqlConn) startWatcher() {
	watcher := make(chan context.Context, 1)
	mc.watcher = watcher
	finished := make(chan struct{})
	mc.finished = finished
	go func() {
		for {
			var ctx context.Context
			select {
			case ctx = <-watcher:
			case <-mc.closech:
				return
			}

			select {
			case <-ctx.Done():
				mc.cancel(ctx.Err())
			case <-finished:
			case <-mc.closech:
				return
			}
		}
	}()
}

func (mc *MysqlConn) CheckNamedValue(nv *driver.NamedValue) (err error) {
	nv.Value, err = converter{}.ConvertValue(nv.Value)
	return
}

// ResetSession implements driver.SessionResetter.
// (From Go 1.10)
func (mc *MysqlConn) ResetSession(ctx context.Context) error {
	if mc.closed.IsSet() {
		return driver.ErrBadConn
	}
	mc.reset = true
	return nil
}

// IsValid implements driver.Validator interface
// (From Go 1.15)
func (mc *MysqlConn) IsValid() bool {
	return !mc.closed.IsSet()
}

func (mc *MysqlConn) ServerVersion() string {
	return mc.serverVersion
}

func (mc *MysqlConn) LastInsertId() uint64 {
	return mc.insertId
}

func (mc *MysqlConn) RowsAffected() uint64 {
	return mc.affectedRows
}

func (mc *MysqlConn) Status() uint16 {
	return uint16(mc.status)
}

func (mc *MysqlConn) UseDb(ctx context.Context, dbName string) error {
	if mc.cfg.DBName == dbName || len(dbName) == 0 {
		return nil
	}
	if err := mc.writeCommandPacketStr(ComInitDB, dbName); err != nil {
		return err
	}
	if err := mc.readResultOK(); err != nil {
		return err
	}
	mc.cfg.DBName = dbName
	return nil
}
