package directive

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/hizzuu/plate-backend/internal/interface/controller"
)

type directive struct {
	userCtrl controller.User
}

type Directive interface {
	Authentication(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error)
	Constraint(ctx context.Context, obj interface{}, next graphql.Resolver, label string, notEmpty *bool, notBlank *bool, pattern *string, min *int, max *int) (interface{}, error)
	CurrentUser(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error)
}

func New(userCtrl controller.User) *directive {
	return &directive{userCtrl: userCtrl}
}
