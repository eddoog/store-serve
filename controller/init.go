package controller

import (
	"github.com/eddoog/store-serve/app/config"
	"github.com/eddoog/store-serve/controller/auth"
	"github.com/eddoog/store-serve/controller/cart"
	"github.com/eddoog/store-serve/controller/product"
	"github.com/eddoog/store-serve/controller/transaction"
	"github.com/eddoog/store-serve/controller/user"
	"github.com/eddoog/store-serve/service/cache"
)

type Controller struct {
	AuthController        auth.IAuthController
	UserController        user.IUserController
	ProductController     product.IProductController
	CartController        cart.ICartController
	TransactionController transaction.ITransactionController
}

func InitController(service *config.Service, cacheService cache.ICacheService) *Controller {
	return &Controller{
		AuthController:        auth.NewAuthController(service.AuthService),
		UserController:        user.NewUserController(service.UserService),
		ProductController:     product.NewProductController(service.ProductService, cacheService),
		CartController:        cart.NewCartController(service.CartService),
		TransactionController: transaction.NewTransactionController(service.TransactionService),
	}
}
