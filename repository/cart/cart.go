package cart

import (
	"fmt"

	"github.com/eddoog/store-serve/domains/entities"
	"github.com/eddoog/store-serve/domains/models"
)

func (c *CartRepository) AddProductToCart(userID uint, productID uint, quantity int) error {
	tx := c.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var product models.Product
	if err := tx.First(&product, "id = ?", productID).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("product not found")
	}

	if product.Stock < quantity {
		tx.Rollback()
		return fmt.Errorf("insufficient stock")
	}

	var cart models.Cart
	if err := tx.FirstOrCreate(&cart, models.Cart{UserID: userID}).Error; err != nil {
		tx.Rollback()
		return err
	}

	var cartItem models.CartItem
	if err := tx.First(&cartItem, "cart_id = ? AND product_id = ?", cart.ID, productID).Error; err == nil {
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

	if err := tx.Save(&cartItem).Error; err != nil {
		tx.Rollback()
		return err
	}

	product.Stock -= quantity
	if err := tx.Save(&product).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}

func (r *CartRepository) GetCart(userID uint) ([]entities.CartItemResponse, float64, error) {
	var items []entities.CartItemResponse
	var total float64

	var cartID uint
	err := r.db.Table("cart").Select("id").Where("user_id = ?", userID).Pluck("id", &cartID).Error
	if err != nil {
		return nil, 0, err
	}

	if cartID == 0 {
		return items, 0, nil
	}

	err = r.db.Table("cart_item").
		Select("cart_item.product_id, product.name, cart_item.price, cart_item.quantity, (cart_item.price * cart_item.quantity) AS subtotal").
		Joins("JOIN product ON cart_item.product_id = product.id").
		Where("cart_item.cart_id = ?", cartID).
		Scan(&items).Error

	if err != nil {
		return nil, 0, err
	}

	err = r.db.Table("cart_item").
		Select("COALESCE(SUM(cart_item.price * cart_item.quantity), 0)").
		Where("cart_item.cart_id = ?", cartID).
		Row().Scan(&total)

	if err != nil {
		return nil, 0, err
	}

	if len(items) == 0 {
		return []entities.CartItemResponse{}, 0, nil
	}

	return items, total, nil
}

func (r *CartRepository) RemoveCartItem(userID uint, productID *uint, quantity int, clearAll bool) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var cart models.Cart
	if err := tx.Where("user_id = ?", userID).First(&cart).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("cart not found")
	}

	if clearAll {
		var cartItems []models.CartItem
		if err := tx.Where("cart_id = ?", cart.ID).Find(&cartItems).Error; err != nil {
			tx.Rollback()
			return err
		}

		for _, item := range cartItems {
			var product models.Product
			if err := tx.First(&product, "id = ?", item.ProductID).Error; err != nil {
				tx.Rollback()
				return fmt.Errorf("product not found for item ID: %d", item.ProductID)
			}
			product.Stock += item.Quantity
			if err := tx.Save(&product).Error; err != nil {
				tx.Rollback()
				return err
			}
		}

		if err := tx.Where("cart_id = ?", cart.ID).Delete(&models.CartItem{}).Error; err != nil {
			tx.Rollback()
			return err
		}

		if err := tx.Commit().Error; err != nil {
			return fmt.Errorf("failed to commit transaction: %v", err)
		}

		return nil
	}

	// Remove specific item if productID is provided
	if productID != nil {
		var cartItem models.CartItem
		if err := tx.Where("cart_id = ? AND product_id = ?", cart.ID, *productID).First(&cartItem).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("item not found in cart")
		}

		var product models.Product
		if err := tx.First(&product, "id = ?", cartItem.ProductID).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("product not found")
		}

		if quantity >= cartItem.Quantity {
			product.Stock += cartItem.Quantity
			if err := tx.Save(&product).Error; err != nil {
				tx.Rollback()
				return err
			}
			if err := tx.Delete(&cartItem).Error; err != nil {
				tx.Rollback()
				return err
			}
		} else {
			cartItem.Quantity -= quantity
			product.Stock += quantity

			if err := tx.Save(&cartItem).Error; err != nil {
				tx.Rollback()
				return err
			}

			if err := tx.Save(&product).Error; err != nil {
				tx.Rollback()
				return err
			}
		}

		if err := tx.Commit().Error; err != nil {
			return fmt.Errorf("failed to commit transaction: %v", err)
		}

		return nil
	}

	tx.Rollback()
	return fmt.Errorf("no product ID provided and clearAll is false")
}
