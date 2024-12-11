package cart

import (
	"github.com/eddoog/store-serve/domains/entities"
	"gorm.io/gorm"
)

type ICartRepository interface {
	AddProductToCart(userID uint, productID uint, quantity int) error
	GetCart(userID uint) ([]entities.CartItemResponse, float64, error)
	RemoveCartItem(userID uint, productID *uint, quantity int, clearAll bool) error
}

type CartRepository struct {
	db *gorm.DB
}

func InitCartRepository(
	db *gorm.DB,
) ICartRepository {
	return &CartRepository{
		db: db,
	}
}
