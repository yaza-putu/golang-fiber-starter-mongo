package core

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yaza-putu/golang-fiber-starter-mongo/internal/config"
	"github.com/yaza-putu/golang-fiber-starter-mongo/internal/http/middleware"
	"github.com/yaza-putu/golang-fiber-starter-mongo/internal/pkg/logger"
	"github.com/yaza-putu/golang-fiber-starter-mongo/internal/routes"
)

func Server() {
	app := fiber.New(fiber.Config{
		AppName: config.App().Name,
	})

	// handle panic error with recover
	app.Use(middleware.PanicMiddleware)
	// call route api
	routes.Api(app)

	// gracefully shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// start server
	go func() {
		if err := app.Listen(fmt.Sprintf(":%d", config.App().Port)); err != nil && err != http.ErrServerClosed {
			logger.New(err, logger.SetType(logger.FATAL))
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	app.ShutdownWithContext(ctx)
}
