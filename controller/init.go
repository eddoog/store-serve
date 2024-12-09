package controller

import (
	"github.com/eddoog/store-serve/app/config"
	"github.com/eddoog/store-serve/controller/auth"
)

type Controller struct {
	AuthController auth.IAuthController
}

func InitController(service *config.Service) *Controller {
	return &Controller{
		AuthController: auth.NewAuthController(service.AuthService),
	}
}
