package config

import (
	"github.com/eddoog/store-serve/service/auth"
	"github.com/eddoog/store-serve/service/cart"
	"github.com/eddoog/store-serve/service/product"
	"github.com/eddoog/store-serve/service/user"
)

type Service struct {
	AuthService    auth.IAuthService
	UserService    user.IUserService
	ProductService product.IProductService
	CartService    cart.ICartService
}

func InitService(
	repository *Repository,
) *Service {
	return &Service{
		AuthService:    auth.InitAuthService(repository.AuthRepository),
		UserService:    user.InitUserService(repository.UserRepository),
		ProductService: product.InitProductService(repository.ProductRepository),
		CartService:    cart.InitCartService(repository.CartRepository),
	}
}
