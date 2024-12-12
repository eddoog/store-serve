package transaction

import (
	"fmt"

	"github.com/eddoog/store-serve/domains/models"
	"github.com/gofiber/fiber/v2"
)

func (t *TransactionService) GetUserTransactions(userID uint) ([]models.Transaction, error) {
	return t.repository.GetUserTransactions(userID)
}

func (t *TransactionService) Checkout(ctx *fiber.Ctx, userID uint) error {
	err := t.repository.Checkout(ctx, userID)
	if err != nil {
		return err
	}

	return nil
}

func (t *TransactionService) CancelTransaction(ctx *fiber.Ctx, txID uint, userID uint) error {
	return t.repository.CancelTransaction(ctx, txID, userID)
}

func (t *TransactionService) ProcessPayment(txID uint, userID uint) error {
	transaction, err := t.repository.GetTransaction(txID)
	if err != nil {
		return err
	}

	if transaction.UserID != userID {
		return fmt.Errorf("unauthorized to update this transaction")
	}

	// IMPROVEMENT: Implement actual payment processing
	paymentSuccess := true

	if paymentSuccess {
		transaction.Status = "paid"
	} else {
		transaction.Status = "failed"
	}

	err = t.repository.UpdateTransaction(transaction)
	if err != nil {
		return err
	}

	return nil
}
