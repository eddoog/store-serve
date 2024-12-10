package middleware

import (
	"strings"

	"github.com/eddoog/store-serve/controller/auth"
	"github.com/gofiber/fiber/v2"
)

func JWTMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Missing or invalid token",
		})
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := auth.ParseJWT(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "token invalid",
			"error":   err.Error(),
		})
	}

	c.Locals("user", claims)
	return c.Next()
}
