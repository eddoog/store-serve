package transaction

import (
	"github.com/eddoog/store-serve/domains/models"
	"gorm.io/gorm"
)

type ITransactionRepository interface {
	GetUserTransactions(userID uint) ([]models.Transaction, error)
	Checkout(userID uint) error
	CancelTransaction(txID uint, userID uint) error
	GetTransaction(txID uint) (*models.Transaction, error)
	UpdateTransaction(transaction *models.Transaction) error
}

type TransactionRepository struct {
	db *gorm.DB
}

func InitTransactionRepository(
	db *gorm.DB,
) ITransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}
