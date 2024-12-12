package config

import (
	"github.com/eddoog/store-serve/repository/auth"
	"github.com/eddoog/store-serve/repository/cart"
	"github.com/eddoog/store-serve/repository/product"
	"github.com/eddoog/store-serve/repository/transaction"
	"github.com/eddoog/store-serve/repository/user"
	"github.com/eddoog/store-serve/service/cache"
	"gorm.io/gorm"
)

type Repository struct {
	AuthRepository        auth.IAuthRepository
	UserRepository        user.IUserRepository
	ProductRepository     product.IProductRepository
	CartRepository        cart.ICartRepository
	TransactionRepository transaction.ITransactionRepository
}

func InitRepository(db *gorm.DB, cacheService cache.ICacheService) *Repository {
	return &Repository{
		AuthRepository:        auth.InitAuthRepository(db),
		UserRepository:        user.InitUserRepository(db),
		ProductRepository:     product.InitProductRepository(db),
		CartRepository:        cart.InitCartRepository(db),
		TransactionRepository: transaction.InitTransactionRepository(db, cacheService),
	}
}
