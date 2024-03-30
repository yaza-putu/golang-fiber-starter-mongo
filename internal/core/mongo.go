package core

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/yaza-putu/golang-fiber-starter-mongo/internal/config"
	"github.com/yaza-putu/golang-fiber-starter-mongo/internal/database"
	"github.com/yaza-putu/golang-fiber-starter-mongo/internal/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Mongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if database.Mongo == nil {
		credential := options.Credential{
			Username: config.DB().User,
			Password: config.DB().Password,
		}

		// connect with auth
		if config.DB().User != "" {
			client, err := mongo.Connect(ctx, options.Client().SetAuth(credential).ApplyURI(fmt.Sprintf("mongodb://%s:%d", config.DB().Host, config.DB().Port)))
			logger.New(err, logger.SetType(logger.FATAL))
			// set instance
			database.Mongo = client.Database(config.DB().Name)
			// check connection
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if client.Ping(ctx, nil) != nil {
				logger.New(errors.New("Failed to connect mongo server"), logger.SetType(logger.FATAL))
			}
		} else {
			client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", config.DB().Host, config.DB().Port)))
			logger.New(err, logger.SetType(logger.FATAL))
			// set instance
			database.Mongo = client.Database(config.DB().Name)
			// check connection
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if client.Ping(ctx, nil) != nil {
				logger.New(errors.New("Failed to connect mongo server"), logger.SetType(logger.FATAL))
			}
		}
	}
}
