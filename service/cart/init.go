package cart

import (
	"github.com/eddoog/store-serve/domains/entities"
	"github.com/eddoog/store-serve/repository/cart"
)

type ICartService interface {
	AddToCart(uint, entities.AddToCartRequest) error
	GetCart(uint) ([]entities.CartItemResponse, float64, error)
	RemoveCartItem(userID uint, productID *uint, quantity int, clearAll bool) error
}

type CartService struct {
	repo cart.ICartRepository
}

func InitCartService(
	repo cart.ICartRepository,
) ICartService {
	return &CartService{
		repo: repo,
	}
}
