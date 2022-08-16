package mysql

import (
	"context"
	"net"
)

type Connector interface {
	// OnConnect handle Client conn with init process
	OnConnect(context.Context, net.Conn) (Conn, error)
}

type Conn interface {
	Run(ctx context.Context, plan QueryPlan) error
}

type Tx interface {
}
