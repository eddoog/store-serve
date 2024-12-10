package auth

import (
	"github.com/eddoog/store-serve/domains"
	"github.com/eddoog/store-serve/domains/entities"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func (a *AuthController) Login(ctx *fiber.Ctx) error {
	return nil
}

func (a *AuthController) Register(ctx *fiber.Ctx) error {
	var userRegister entities.UserRegister

	if err := ctx.BodyParser(&userRegister); err != nil {
		logrus.Error(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if err := a.Validator.Struct(&userRegister); err != nil {
		validationErrors := domains.CustomValidationError(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  validationErrors,
		})
	}

	err := a.AuthService.Register(userRegister)

	if err != nil {
		logrus.Error(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
	})
}
