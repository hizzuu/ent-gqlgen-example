package repository

import (
	"context"

	"github.com/hizzuu/plate-backend/ent"
)

type dbClient interface {
	Client(ctx context.Context) *ent.Client
}
