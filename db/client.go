package db

import (
	"context"
	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/long2ice/longurl/config"
	"github.com/long2ice/longurl/ent"
	"github.com/long2ice/longurl/ent/migrate"
	log "github.com/sirupsen/logrus"
	"time"
)

var Client *ent.Client

func init() {
	var err error
	drv, err := sql.Open(config.DatabaseConfig.Type, config.DatabaseConfig.Dsn)
	if err != nil {
		panic(err)
	}
	db := drv.DB()
	db.SetConnMaxLifetime(time.Hour)
	db.SetConnMaxIdleTime(time.Hour)
	Client = ent.NewClient(ent.Driver(drv))
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
