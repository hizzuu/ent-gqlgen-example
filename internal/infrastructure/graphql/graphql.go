package graphql

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/hizzuu/plate-backend/graph/directive"
	"github.com/hizzuu/plate-backend/graph/generated"
	"github.com/hizzuu/plate-backend/graph/middleware"
	"github.com/hizzuu/plate-backend/internal/infrastructure/logger"
	"github.com/hizzuu/plate-backend/internal/interface/controller"
	"github.com/hizzuu/plate-backend/internal/interface/resolver"
)

func New(
	logger logger.Logger,
	ctrl controller.Controller,
) *handler.Server {
	d := directive.New(ctrl.User)
	r := resolver.New(ctrl)
	mw := middleware.New(logger)

	conf := generated.Config{Resolvers: r}
	conf.Directives.Constraint = d.Constraint
	conf.Directives.Authentication = d.Authentication
	conf.Directives.CurrentUser = d.CurrentUser

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(conf))
	srv.SetErrorPresenter(mw.ErrorPresenter())
	srv.SetRecoverFunc(mw.Recover())
	srv.AroundOperations(mw.Operation())

	return srv
}

func NewPlayGroundHandler(r generated.ResolverRoot) http.Handler {
	p := playground.Handler("GraphQL playground", "/query")

	return p
}
