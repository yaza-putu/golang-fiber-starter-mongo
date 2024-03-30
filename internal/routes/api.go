package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yaza-putu/golang-fiber-starter-mongo/internal/http/response"
)

func Api(router fiber.Router) {
	router.Get("/", func(c *fiber.Ctx) error {
		time.Sleep(20 * time.Second)
		return c.Status(fiber.StatusOK).JSON(response.Api(
			response.SetCode(fiber.StatusOK),
			response.SetMessage("OK"),
		))
	})
}
