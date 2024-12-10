package middleware

import (
	"strings"

	"github.com/eddoog/store-serve/controller/auth"
	"github.com/gofiber/fiber/v2"
)

type UserClaims struct {
	UserID uint    `json:"user_id"`
	Email  string  `json:"email"`
	Exp    float64 `json:"exp"`
}

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

	userClaims := UserClaims{
		UserID: uint(claims["user_id"].(float64)),
		Email:  claims["email"].(string),
		Exp:    claims["exp"].(float64),
	}

	c.Locals("user", userClaims)
	return c.Next()
}
