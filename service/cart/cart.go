package cart

import "github.com/eddoog/store-serve/domains/entities"

func (s *CartService) AddToCart(userID uint, req entities.AddToCartRequest) error {
	return s.repo.AddProductToCart(userID, req.ProductID, req.Quantity)
}

func (s *CartService) GetCart(userID uint) ([]entities.CartItemResponse, float64, error) {
	return s.repo.GetCart(userID)
}

func (s *CartService) RemoveCartItem(userID uint, productID *uint, quantity int, clearAll bool) error {
	return s.repo.RemoveCartItem(userID, productID, quantity, clearAll)
}
