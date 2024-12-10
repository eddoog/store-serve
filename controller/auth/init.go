package auth

import (
	"github.com/eddoog/store-serve/service/auth"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type IAuthController interface {
	Login(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
}

type AuthController struct {
	AuthService auth.IAuthService
	Validator   *validator.Validate
}

func NewAuthController(authService auth.IAuthService) IAuthController {
	return &AuthController{
		AuthService: authService,
		Validator:   validator.New(),
	}
}
