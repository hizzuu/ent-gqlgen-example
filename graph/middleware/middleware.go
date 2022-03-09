package middleware

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/99designs/gqlgen/graphql"
	"github.com/hizzuu/plate-backend/internal/domain"
	"github.com/hizzuu/plate-backend/internal/infrastructure/logger"
)

type middleware struct {
	logger logger.Logger
}

type Middleware interface {
	ErrorPresenter() graphql.ErrorPresenterFunc
	Operation() graphql.OperationMiddleware
	Recover() graphql.RecoverFunc
}

func New(logger logger.Logger) *middleware {
	return &middleware{logger: logger}
}

// ErrorPresenter
func (m *middleware) ErrorPresenter() graphql.ErrorPresenterFunc {
	return graphql.ErrorPresenterFunc(func(ctx context.Context, e error) *gqlerror.Error {
		var extendedError interface{ Extensions() map[string]interface{} }
		err := graphql.DefaultErrorPresenter(ctx, e)
		for e != nil {
			u, ok := e.(interface {
				Unwrap() error
			})
			if !ok {
				break
			}

			if domain.IsStackTrace(e) {
				e = u.Unwrap()
				continue
			}

			if !domain.IsError(e) {
				e = u.Unwrap()
				continue
			}

			err = &gqlerror.Error{
				Path:    graphql.GetPath(ctx),
				Message: e.Error(),
			}
			if errors.As(e, &extendedError) {
				err.Extensions = extendedError.Extensions()
			}

			e = u.Unwrap()
		}

		return err
	})
}

// Operation
func (m *middleware) Operation() graphql.OperationMiddleware {
	return graphql.OperationMiddleware(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		oc := graphql.GetOperationContext(ctx)
		m.logger.InfofWithCtx(ctx, "operation_name: %s, variables: %v", oc.OperationName, oc.Variables)
		return next(ctx)
	})
}

// Recover
func (m *middleware) Recover() graphql.RecoverFunc {
	return graphql.RecoverFunc(func(ctx context.Context, err interface{}) error {
		e := errors.Errorf("%v", err)
		m.logger.InfofWithCtx(ctx, e.Error())
		return domain.NewInternalServerError(e)
	})
}
