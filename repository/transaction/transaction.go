package transaction

import (
	"fmt"
	"time"

	"github.com/eddoog/store-serve/domains/models"
)

func (r *TransactionRepository) GetUserTransactions(userID uint) ([]models.Transaction, error) {
	var transactions []models.Transaction
	if err := r.db.Preload("Items").Where("user_id = ?", userID).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *TransactionRepository) Checkout(userID uint) error {
	tx := r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var cart models.Cart
	if err := tx.Where("user_id = ?", userID).First(&cart).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("cart not found: %v", err)
	}

	if cart.ID == 0 {
		tx.Rollback()
		return fmt.Errorf("cart is empty")
	}

	var cartItems []models.CartItem
	if err := tx.Where("cart_id = ?", cart.ID).Find(&cartItems).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to fetch cart items: %v", err)
	}

	if len(cartItems) == 0 {
		tx.Rollback()
		return fmt.Errorf("cart is empty")
	}

	productIDs := make([]uint, len(cartItems))
	for i, item := range cartItems {
		productIDs[i] = item.ProductID
	}

	var products []models.Product
	if err := tx.Where("id IN ?", productIDs).Find(&products).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to fetch products: %v", err)
	}

	productMap := make(map[uint]*models.Product)
	for _, product := range products {
		productMap[product.ID] = &product
	}

	var total float64
	for _, item := range cartItems {
		product, exists := productMap[item.ProductID]
		if !exists {
			tx.Rollback()
			return fmt.Errorf("product not found for ID %d", item.ProductID)
		}
		if product.Stock < item.Quantity {
			tx.Rollback()
			return fmt.Errorf("insufficient stock for product ID %d", item.ProductID)
		}
		total += float64(item.Quantity) * item.Price
	}

	transaction := models.Transaction{
		UserID:    userID,
		Total:     total,
		Status:    "pending",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create transaction: %v", err)
	}

	for _, item := range cartItems {
		product := productMap[item.ProductID]
		transactionItem := models.TransactionItem{
			TransactionID: transaction.ID,
			ProductID:     item.ProductID,
			Quantity:      item.Quantity,
			Price:         item.Price,
			Subtotal:      float64(item.Quantity) * item.Price,
		}
		if err := tx.Create(&transactionItem).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to create transaction item: %v", err)
		}

		product.Stock -= item.Quantity
		if err := tx.Save(product).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to update product stock: %v", err)
		}
	}

	if err := tx.Where("cart_id = ?", cart.ID).Delete(&models.CartItem{}).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to clear cart: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}

func (r *TransactionRepository) CancelTransaction(txID uint, userID uint) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var transaction models.Transaction
	if err := tx.First(&transaction, txID).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("transaction not found")
	}

	if transaction.UserID != userID {
		tx.Rollback()
		return fmt.Errorf("unauthorized to cancel this transaction")
	}

	if !transaction.DeletedAt.Time.IsZero() {
		tx.Rollback()
		return fmt.Errorf("transaction already canceled")
	}

	var transactionItems []models.TransactionItem
	if err := tx.Where("transaction_id = ?", txID).Find(&transactionItems).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to fetch transaction items: %v", err)
	}

	productIDs := make([]uint, len(transactionItems))
	for i, item := range transactionItems {
		productIDs[i] = item.ProductID
	}

	var products []models.Product
	if err := tx.Where("id IN ?", productIDs).Find(&products).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to fetch products: %v", err)
	}

	productMap := make(map[uint]*models.Product)
	for _, product := range products {
		productMap[product.ID] = &product
	}

	for _, item := range transactionItems {
		product, exists := productMap[item.ProductID]
		if !exists {
			tx.Rollback()
			return fmt.Errorf("product not found for ID %d", item.ProductID)
		}
		product.Stock += item.Quantity
		if err := tx.Save(product).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to update product stock: %v", err)
		}
	}

	// Soft delete transaction items
	for _, item := range transactionItems {
		if err := tx.Delete(&item).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to delete transaction item: %v", err)
		}
	}

	// Soft delete transaction
	if err := tx.Delete(&transaction).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete transaction: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}

func (r *TransactionRepository) GetTransaction(txID uint) (*models.Transaction, error) {
	var transaction models.Transaction
	if err := r.db.First(&transaction, txID).Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *TransactionRepository) UpdateTransaction(transaction *models.Transaction) error {
	return r.db.Save(transaction).Error
}
