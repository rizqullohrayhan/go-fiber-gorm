package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func ApiKeyMiddleware(ctx *fiber.Ctx) error {
	apiKey := ctx.Get("X-API-Key")
	if apiKey == "" || apiKey != os.Getenv("API_KEY") {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}
	return ctx.Next()
}