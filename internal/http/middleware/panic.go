package middleware

import (
	"fmt"
	"runtime"

	"github.com/gofiber/fiber/v2"
	"github.com/yaza-putu/golang-fiber-starter-mongo/internal/http/response"
	"github.com/yaza-putu/golang-fiber-starter-mongo/internal/pkg/logger"
)

func PanicMiddleware(c *fiber.Ctx) error {
	defer func() {
		if err := recover(); err != nil {
			// Retrieve the stack trace.
			// only capture 4 error stack
			var pcs [4]uintptr
			n := runtime.Callers(3, pcs[:])

			frames := runtime.CallersFrames(pcs[:n])
			errors := []string{}
			for {
				frame, more := frames.Next()
				// capture error with location file, line code and function to easy debug
				errors = append(errors, fmt.Sprintf("\t%s:%d %s\n", frame.File, frame.Line, frame.Function))
				if !more {
					break
				}
			}
			logger.New(fmt.Errorf("panic recover : %v at %v", err, errors))
			c.Status(fiber.StatusInternalServerError).JSON(response.Api(
				response.SetCode(fiber.StatusInternalServerError),
				response.SetMessage("Internal Server Error"),
			))
		}
	}()
	return c.Next()
}
