package database

import (
	"context"

	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/mongo"
)

var Mongo *mongo.Database

func UpMIgration() error {
	migrate.SetDatabase(Mongo)
	return migrate.Up(context.Background(), migrate.AllAvailable)
}

func DownMigration() error {
	migrate.SetDatabase(Mongo)
	if err := migrate.Down(context.Background(), migrate.AllAvailable); err != nil {
		return err
	}

	Mongo.Collection("migrations").Drop(context.Background())
	return nil
}
