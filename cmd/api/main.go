package main

import (
	"context"
	"fmt"

	"github.com/hizzuu/plate-backend/conf"
	"github.com/hizzuu/plate-backend/internal/infrastructure/db"
	"github.com/hizzuu/plate-backend/internal/infrastructure/firebase"
	"github.com/hizzuu/plate-backend/internal/infrastructure/graphql"
	"github.com/hizzuu/plate-backend/internal/infrastructure/logger"
	"github.com/hizzuu/plate-backend/internal/infrastructure/server"
	"github.com/hizzuu/plate-backend/internal/infrastructure/storage"
	"github.com/hizzuu/plate-backend/internal/registry"
)

func main() {
	conf.ReadConfig(conf.ReadConfigOption{})
	ctx := context.Background()
	logger := logger.New()

	defer func() {
		if rec := recover(); rec != nil {
			logger.Fatal(fmt.Errorf("recovered from: %v", rec))
		}
	}()

	conn, err := db.NewMysqlDB()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	dbClient := db.NewClient(conn)

	strClient, err := storage.New(ctx)
	if err != nil {
		panic(err)
	}

	fb, err := firebase.New(ctx)
	if err != nil {
		panic(err)
	}

	authClient, err := firebase.NewAuthClient(ctx, fb)
	if err != nil {
		panic(err)
	}

	r := registry.New(dbClient, strClient)
	ctrl := r.NewController()

	srv := graphql.New(logger, ctrl)

	server := server.New(logger, authClient, srv)
	if err := server.Listen(); err != nil {
		panic(err)
	}
}
