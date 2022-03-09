package utils

import (
	"context"
	"database/sql"
	"testing"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"

	"github.com/google/uuid"
	"github.com/hizzuu/plate-backend/ent"
	"github.com/hizzuu/plate-backend/ent/enttest"
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

func NewDBClient(t *testing.T) (*client, func()) {
	conn, err := sql.Open("txdb", uuid.New().String())
	if err != nil {
		panic(err)
	}

	c := enttest.NewClient(t, enttest.WithOptions(ent.Driver(entsql.OpenDB(dialect.MySQL, conn))))

	return &client{client: c}, func() {
		conn.Close()
		c.Close()
	}
}

func (c *client) DoInTx(ctx context.Context, f func(ctx context.Context) (interface{}, error)) (interface{}, error) {
	tx, err := c.client.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, txCtxKey, tx)
	v, err := f(ctx)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return nil, err
		}

		return nil, err
	}

	if err := tx.Commit(); err != nil {
		if err := tx.Rollback(); err != nil {
			return nil, err
		}

		return nil, err
	}

	return v, nil
}

func (c *client) Client(ctx context.Context) *ent.Client {
	if tx, ok := ctx.Value(txCtxKey).(*ent.Tx); ok {
		return tx.Client()
	}

	return c.client
}
