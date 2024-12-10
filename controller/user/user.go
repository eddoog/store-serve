package user

import (
	"github.com/eddoog/store-serve/middleware"
	"github.com/gofiber/fiber/v2"
)

func (u *UserController) Profile(ctx *fiber.Ctx) error {
	userClaims := ctx.Locals("user")

	if userClaims == nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	claims, ok := userClaims.(middleware.UserClaims)
	if !ok {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve user claims",
		})
	}

	user, err := u.UserService.GetProfile(uint(claims.UserID))

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve user profile",
		})
	}

	return ctx.JSON(fiber.Map{
		"user": user,
	})
}
