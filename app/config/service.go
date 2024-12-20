package config

import (
	"github.com/eddoog/store-serve/service/auth"
	"github.com/eddoog/store-serve/service/cache"
	"github.com/eddoog/store-serve/service/cart"
	"github.com/eddoog/store-serve/service/product"
	"github.com/eddoog/store-serve/service/transaction"
	"github.com/eddoog/store-serve/service/user"
	"github.com/valkey-io/valkey-go"
)

type Service struct {
	AuthService        auth.IAuthService
	UserService        user.IUserService
	ProductService     product.IProductService
	CartService        cart.ICartService
	TransactionService transaction.ITransactionService
}

func InitService(
	repository *Repository,
) *Service {
	return &Service{
		AuthService:        auth.InitAuthService(repository.AuthRepository),
		UserService:        user.InitUserService(repository.UserRepository),
		ProductService:     product.InitProductService(repository.ProductRepository),
		CartService:        cart.InitCartService(repository.CartRepository),
		TransactionService: transaction.InitTransactionService(repository.TransactionRepository),
	}
}

func InitCacheService(rdb valkey.Client) cache.ICacheService {
	return cache.NewCacheService(rdb)
}
