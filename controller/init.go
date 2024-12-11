package controller

import (
	"github.com/eddoog/store-serve/app/config"
	"github.com/eddoog/store-serve/controller/auth"
	"github.com/eddoog/store-serve/controller/cart"
	"github.com/eddoog/store-serve/controller/product"
	"github.com/eddoog/store-serve/controller/user"
)

type Controller struct {
	AuthController    auth.IAuthController
	UserController    user.IUserController
	ProductController product.IProductController
	CartController    cart.ICartController
}

func InitController(service *config.Service) *Controller {
	return &Controller{
		AuthController:    auth.NewAuthController(service.AuthService),
		UserController:    user.NewUserController(service.UserService),
		ProductController: product.NewProductController(service.ProductService),
		CartController:    cart.NewCartController(service.CartService),
	}
}
