package main

import (
	"context"
	"log"

	"github.com/hizzuu/plate-backend/conf"
	"github.com/hizzuu/plate-backend/ent/migrate"
	"github.com/hizzuu/plate-backend/internal/infrastructure/db"
)

func main() {
	conf.ReadConfig(conf.ReadConfigOption{})
	ctx := context.Background()

	conn, err := db.NewMysqlDB()
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer conn.Close()

	c := db.NewClient(conn)
	defer conn.Close()

	err = c.Client(ctx).Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	log.Println("migration succeeded")
}
