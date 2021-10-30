package db

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"long2ice/longurl/config"
	"long2ice/longurl/ent"
	"long2ice/longurl/ent/migrate"
)

var Client *ent.Client

func init() {
	var err error
	Client, err = ent.Open(config.DatabaseConfig.Type, config.DatabaseConfig.Dsn)
	if err != nil {
		log.Fatalf("Connect to database error: %v", err)
	}
	err = Client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("Failed creating schema resources: %v", err)
	}
}
