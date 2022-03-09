package server

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/hizzuu/plate-backend/conf"
	"github.com/hizzuu/plate-backend/internal/infrastructure/firebase"
	"github.com/hizzuu/plate-backend/internal/infrastructure/logger"
)

type server struct {
	logger logger.Logger
}

type Server interface {
	Listen(srv *handler.Server, authClient firebase.AuthClient) error
}

func New(
	logger logger.Logger,
	authClient firebase.AuthClient,
	srv *handler.Server,
) *server {
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", cors(requestID(authentication(srv, authClient))))
	return &server{
		logger: logger,
	}
}

func (s *server) Listen() error {
	s.logger.Infof("Start the server...")
	return http.ListenAndServe(":"+conf.C.Api.Port, nil)
}
