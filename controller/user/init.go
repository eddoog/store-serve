package user

import (
	"github.com/eddoog/store-serve/service/user"
	"github.com/gofiber/fiber/v2"
)

type IUserController interface {
	Profile(ctx *fiber.Ctx) error
}

type UserController struct {
	UserService user.IUserService
}

func NewUserController(userService user.IUserService) IUserController {
	return &UserController{
		UserService: userService,
	}
}
