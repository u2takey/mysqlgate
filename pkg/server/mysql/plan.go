package mysql

import (
	"context"

	"github.com/u2takey/mysqlgate/pkg/sql"
)

type QueryContext struct {
	context.Context
	mc      *MysqlConn
	aborted bool
}

func NewQueryContext(ctx context.Context, conn *MysqlConn) *QueryContext {
	return &QueryContext{Context: ctx, mc: conn}
}

type QueryPlan interface {
	Query(ctx *QueryContext, sql string) error
}

type aggregatedQueryPlan struct {
	plans []QueryPlan
}

func NewQueryPlan(db *sql.DB) QueryPlan {
	return &aggregatedQueryPlan{
		plans: []QueryPlan{&defaultQueryPlan{db: db}},
	}
}

func (q *aggregatedQueryPlan) Query(ctx *QueryContext, sql string) error {
	for _, p := range q.plans {
		if ctx.aborted {
			return nil
		}
		if err := p.Query(ctx, sql); err != nil {
			return err
		}
	}
	return nil
}

type defaultQueryPlan struct {
	db *sql.DB
}

func (q *defaultQueryPlan) Query(ctx *QueryContext, sql string) error {
	conn, err := q.db.Conn(ctx)
	if err != nil {
		return err
	}
	rows, err := conn.QueryContextExtend(ctx, sql)
	if err != nil {
		return err
	}
	return ctx.mc.writeResultSet(rows)
}
