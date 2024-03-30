package core

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yaza-putu/golang-fiber-starter-mongo/internal/config"
	"github.com/yaza-putu/golang-fiber-starter-mongo/internal/pkg/logger"
	"github.com/yaza-putu/golang-fiber-starter-mongo/internal/routes"
)

func Server() {
	app := fiber.New(fiber.Config{
		AppName: config.App().Name,
	})

	// call route api
	routes.Api(app)

	// gracefully shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// start server
	go func() {
		if err := app.Listen(fmt.Sprintf(":%d", config.App().Port)); err != nil && err != http.ErrServerClosed {
			logger.New(errors.New("shutting down the server"), logger.SetType(logger.FATAL))
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.ShutdownWithContext(ctx); err != nil {
		logger.New(err, logger.SetType(logger.FATAL))
	} else {
		fmt.Println("Gracefully Server Shutdown")
	}
}
