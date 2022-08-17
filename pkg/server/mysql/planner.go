package mysql

import (
	"context"

	"github.com/u2takey/mysqlgate/pkg/sql"
	"github.com/u2takey/mysqlgate/pkg/sql/driver"
	parser "github.com/u2takey/sqlparser"
	"github.com/u2takey/sqlparser/ast"
	_ "github.com/u2takey/sqlparser/test_driver"
)

type QueryContext struct {
	context.Context
	mc   *MysqlConn
	db   *sql.DB
	data string
	cmd  byte

	stmts     []ast.StmtNode
	sqlParsed uint8

	aborted bool
	lastErr error
}

func NewQueryContext(ctx context.Context, db *sql.DB) *QueryContext {
	return &QueryContext{Context: ctx, db: db}
}

func (q *QueryContext) WithCmdData(cmd byte, data string) *QueryContext {
	q.cmd, q.data = cmd, data
	return q
}

func (q *QueryContext) WithConn(mc *MysqlConn) *QueryContext {
	q.mc = mc
	return q
}

func (q *QueryContext) Abort() {
	q.aborted = true
}

type QueryPlan interface {
	Query(ctx *QueryContext) error
	InitDB(ctx *QueryContext) error
}

type aggregatedQueryPlan struct {
	plans []QueryPlan
}

func NewQueryPlan() QueryPlan {
	return &aggregatedQueryPlan{
		plans: []QueryPlan{
			&parserPlan{},
			&defaultQueryPlan{},
		},
	}
}

func (q *aggregatedQueryPlan) Query(ctx *QueryContext) error {
	for _, p := range q.plans {
		if ctx.aborted {
			return nil
		}
		if err := p.Query(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (q *aggregatedQueryPlan) InitDB(ctx *QueryContext) error {
	for _, p := range q.plans {
		if ctx.aborted {
			return nil
		}
		if err := p.InitDB(ctx); err != nil {
			return err
		}
	}
	return nil
}

type defaultQueryPlan struct {
}

func (q *defaultQueryPlan) InitDB(ctx *QueryContext) error {
	conn, err := ctx.db.Conn(ctx)
	if err != nil {
		return err
	}
	err = conn.Raw(func(driverConn interface{}) error {
		if c, ok := driverConn.(driver.ConnExtend); !ok {
			return NewCustomError(ErUnknownError, "init db not supported on backend driver")
		} else {
			if err := c.UseDb(ctx, ctx.data); err != nil {
				return err
			}
			ctx.mc.database = ctx.data
			return nil
		}
	})
	if err != nil {
		return err
	}
	return ctx.mc.writeOK(nil)
}

func (q *defaultQueryPlan) Query(ctx *QueryContext) error {
	conn, err := ctx.db.Conn(ctx)
	if err != nil {
		return err
	}
	rows, err := conn.QueryContextExtend(ctx, ctx.data)
	if err != nil {
		return err
	}
	if col, err := rows.Columns(); err == nil && len(col) == 0 {
		return ctx.mc.writeOK(&MysqlResult{
			Status:       StatusFlag(rows.Status),
			AffectedRows: rows.AffectedRows,
			InsertId:     rows.InsertId,
		})
	}
	return ctx.mc.writeResultSet(rows)
}

type parserPlan struct {
}

func (q *parserPlan) InitDB(ctx *QueryContext) error {
	return nil
}

func (q *parserPlan) Query(ctx *QueryContext) error {
	sqlParser := parser.New()
	stmts, _, err := sqlParser.Parse(ctx.data, "", "")
	if err != nil {
		return err
	}
	ctx.sqlParsed += 1
	ctx.stmts = stmts
	return nil
}
