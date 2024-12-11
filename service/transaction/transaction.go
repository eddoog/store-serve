package transaction

import "github.com/eddoog/store-serve/domains/models"

func (t *TransactionService) GetUserTransactions(userID uint) ([]models.Transaction, error) {
	return t.repository.GetUserTransactions(userID)
}

func (t *TransactionService) Checkout(userID uint) error {
	err := t.repository.Checkout(userID)
	if err != nil {
		return err
	}

	return nil
}

func (t *TransactionService) CancelTransaction(txID uint, userID uint) error {
	return t.repository.CancelTransaction(txID, userID)
}
