package plugin

import (
	"context"

	"github.com/u2takey/mysqlgate/pkg/sql"
)

type DefaultPlugin struct {
	db *sql.DB
}

func (p *DefaultPlugin) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return p.db.QueryContext(ctx, query, args...)
}

func (p *DefaultPlugin) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	p.db.Begin()
	return p.db.ExecContext(ctx, query, args...)
}
