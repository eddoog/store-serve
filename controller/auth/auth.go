package auth

import (
	"github.com/eddoog/store-serve/domains"
	"github.com/eddoog/store-serve/domains/entities"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func (a *AuthController) Login(ctx *fiber.Ctx) error {
	var userLogin entities.UserLogin

	if err := ctx.BodyParser(&userLogin); err != nil {
		logrus.Error(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if err := a.Validator.Struct(&userLogin); err != nil {
		validationErrors := domains.CustomValidationError(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  validationErrors,
		})
	}

	user, err := a.AuthService.Login(userLogin)

	if err != nil {
		logrus.Error(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	token, err := GenerateJWT(user.ID, user.Email)

	if err != nil {
		logrus.Error(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to generate token",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token":   token,
		"message": "Login successful",
	})
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
