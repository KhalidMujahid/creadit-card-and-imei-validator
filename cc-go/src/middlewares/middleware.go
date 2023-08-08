package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func PageNotFound(c *fiber.Ctx) error {
	return c.Next()
}

func ErrorHandler(c *fiber.Ctx) error {
	statusCode := c.Status(200)

	if statusCode == 200 {
		statusCode = 500
	} else {
		statusCode = statusCode
	}
	return c.Status(statusCode).JSON(fiber.Map{
		"error": c.Error(),
	})
}
