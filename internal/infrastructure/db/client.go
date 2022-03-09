package db

import (
	"context"
	"database/sql"

	entsql "entgo.io/ent/dialect/sql"

	"github.com/hizzuu/plate-backend/conf"
	"github.com/hizzuu/plate-backend/ent"
	"github.com/hizzuu/plate-backend/internal/domain"
)

const txCtxKey contextKey = "tx"

type contextKey string

type client struct {
	client *ent.Client
}

type Client interface {
	Client(ctx context.Context) *ent.Client
	DoInTx(ctx context.Context, f func(ctx context.Context) (interface{}, error)) (interface{}, error)
}

func NewClient(conn *sql.DB) *client {
	c := ent.NewClient(ent.Driver(entsql.OpenDB(conf.C.DB.Dbms, conn)))
	if conf.C.App.Debug {
		c = c.Debug()
	}

	return &client{client: c}
}

func (c *client) DoInTx(ctx context.Context, f func(ctx context.Context) (interface{}, error)) (interface{}, error) {
	tx, err := c.client.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, domain.NewDBError(err)
	}

	ctx = context.WithValue(ctx, txCtxKey, tx)
	v, err := f(ctx)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return nil, domain.NewDBError(err)
		}

		return nil, domain.NewDBError(err)
	}

	if err := tx.Commit(); err != nil {
		if err := tx.Rollback(); err != nil {
			return nil, domain.NewDBError(err)
		}

		return nil, domain.NewDBError(err)
	}

	return v, nil
}

func (c *client) Client(ctx context.Context) *ent.Client {
	if tx, ok := ctx.Value(txCtxKey).(*ent.Tx); ok {
		return tx.Client()
	}

	return c.client
}
