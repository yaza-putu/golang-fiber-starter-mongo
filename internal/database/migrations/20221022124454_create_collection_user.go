package migrations

import (
	"context"

	migrate "github.com/xakep666/mongo-migrate"
	"github.com/yaza-putu/golang-fiber-starter-mongo/internal/apps/auth/entity"
	"github.com/yaza-putu/golang-fiber-starter-mongo/internal/database"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	const collectionName = "users"

	migrate.NewMigrate(database.Mongo, migrate.Migration{
		Version:     1,
		Description: "User collections",
		Up: func(ctx context.Context, db *mongo.Database) error {
			user := entity.User{}
			_, err := db.Collection(collectionName).InsertOne(context.Background(), user)
			if err != nil {
				return err
			}
			return nil
		},
		Down: func(ctx context.Context, db *mongo.Database) error {
			db.Collection(collectionName).Drop(context.Background())
			return nil
		},
	})
}
