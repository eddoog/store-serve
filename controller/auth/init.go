package auth

import (
	"github.com/eddoog/store-serve/service/auth"
)

type IAuthController interface {
	Login()
	Register()
}

type AuthController struct {
	AuthService auth.IAuthService
}

func NewAuthController(authService auth.IAuthService) IAuthController {
	return &AuthController{
		AuthService: authService,
	}
}
