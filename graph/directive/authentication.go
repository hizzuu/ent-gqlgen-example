package directive

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/hizzuu/plate-backend/graph/model"
	"github.com/hizzuu/plate-backend/internal/domain"
)

func (d *directive) Authentication(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	err, ok := ctx.Value(model.AuthErrorCtxKey).(error)
	if ok {
		return nil, domain.NewUnauthorizedError(err)
	}

	return next(ctx)
}
