package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yaza-putu/golang-fiber-starter-mongo/internal/http/response"
)

func Api(router fiber.Router) {
	router.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(response.Api(
			response.SetCode(fiber.StatusOK),
			response.SetMessage("OK"),
		))
	})
}
