package cart

import (
	"fmt"

	"github.com/eddoog/store-serve/domains/entities"
	"github.com/eddoog/store-serve/domains/models"
)

func (c *CartRepository) AddProductToCart(userID uint, productID uint, quantity int) error {
	var product models.Product
	if err := c.db.First(&product, "id = ?", productID).Error; err != nil {
		return fmt.Errorf("product not found")
	}

	if product.Stock < quantity {
		return fmt.Errorf("insufficient stock")
	}

	var cart models.Cart
	if err := c.db.FirstOrCreate(&cart, models.Cart{UserID: userID}).Error; err != nil {
		return err
	}

	var cartItem models.CartItem
	if err := c.db.First(&cartItem, "cart_id = ? AND product_id = ?", cart.ID, productID).Error; err == nil {
		cartItem.Quantity += quantity
		cartItem.Price = product.Price
	} else {
		cartItem = models.CartItem{
			CartID:    cart.ID,
			ProductID: productID,
			Quantity:  quantity,
			Price:     product.Price,
		}
	}

	if err := c.db.Save(&cartItem).Error; err != nil {
		return err
	}

	product.Stock -= quantity

	return c.db.Save(&product).Error
}

func (r *CartRepository) GetCart(userID uint) ([]entities.CartItemResponse, float64, error) {
	var items []entities.CartItemResponse
	var total float64

	err := r.db.Table("cart_item").
		Select("cart_item.product_id, product.name, cart_item.price, cart_item.quantity, (cart_item.price * cart_item.quantity) AS subtotal").
		Joins("JOIN product ON cart_item.product_id = product.id").
		Where("cart_item.cart_id = (SELECT id FROM cart WHERE user_id = ?)", userID).
		Scan(&items).Error

	if err != nil {
		return nil, 0, err
	}

	err = r.db.Table("cart_item").
		Select("SUM(cart_item.price * cart_item.quantity)").
		Joins("JOIN product ON cart_item.product_id = product.id").
		Where("cart_item.cart_id = (SELECT id FROM cart WHERE user_id = ?)", userID).
		Row().Scan(&total)

	if err != nil {
		return nil, 0, err
	}

	return items, total, nil
}

func (r *CartRepository) RemoveCartItem(userID uint, productID *uint, quantity int, clearAll bool) error {
	var cart models.Cart
	if err := r.db.Where("user_id = ?", userID).First(&cart).Error; err != nil {
		return fmt.Errorf("cart not found")
	}

	if clearAll {
		var cartItems []models.CartItem
		if err := r.db.Where("cart_id = ?", cart.ID).Find(&cartItems).Error; err != nil {
			return err
		}

		for _, item := range cartItems {
			var product models.Product
			if err := r.db.First(&product, "id = ?", item.ProductID).Error; err != nil {
				return fmt.Errorf("product not found for item ID: %d", item.ProductID)
			}
			product.Stock += item.Quantity
			if err := r.db.Save(&product).Error; err != nil {
				return err
			}
		}

		return r.db.Where("cart_id = ?", cart.ID).Delete(&models.CartItem{}).Error
	}

	// Remove specific item if productID is provided
	if productID != nil {
		var cartItem models.CartItem
		if err := r.db.Where("cart_id = ? AND product_id = ?", cart.ID, *productID).First(&cartItem).Error; err != nil {
			return fmt.Errorf("item not found in cart")
		}

		var product models.Product
		if err := r.db.First(&product, "id = ?", cartItem.ProductID).Error; err != nil {
			return fmt.Errorf("product not found")
		}

		if quantity >= cartItem.Quantity {
			product.Stock += cartItem.Quantity
			if err := r.db.Save(&product).Error; err != nil {
				return err
			}
			return r.db.Delete(&cartItem).Error
		}

		cartItem.Quantity -= quantity
		product.Stock += quantity

		if err := r.db.Save(&cartItem).Error; err != nil {
			return err
		}

		return r.db.Save(&product).Error
	}

	return fmt.Errorf("no product ID provided and clearAll is false")
}
