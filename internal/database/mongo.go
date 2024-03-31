package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type MigrateFunc func(context context.Context, db *mongo.Database) error

var (
	Mongo         *mongo.Database
	upMigration   []MigrateFunc
	downMigration []MigrateFunc
)

func MigrationRegister(up MigrateFunc, down MigrateFunc) {
	upMigration = append(upMigration, up)
	downMigration = append(downMigration, down)
}

func UpMIgration() error {
	for i := 0; i < len(upMigration); i++ {
		err := upMigration[i](context.Background(), Mongo)
		if err != nil {
			return err
		}
	}
	return nil
}

func DownMigration() error {
	for i := 0; i < len(downMigration); i++ {
		err := downMigration[i](context.Background(), Mongo)
		if err != nil {
			return err
		}
	}
	return nil
}
