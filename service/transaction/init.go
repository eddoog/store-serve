package transaction

import (
	"github.com/eddoog/store-serve/domains/models"
	"github.com/eddoog/store-serve/repository/transaction"
)

type ITransactionService interface {
	GetUserTransactions(userID uint) ([]models.Transaction, error)
	Checkout(userID uint) error
	CancelTransaction(txID uint, userID uint) error
	ProcessPayment(txID uint, userID uint) error
}

type TransactionService struct {
	repository transaction.ITransactionRepository
}

func InitTransactionService(
	repository transaction.ITransactionRepository,
) ITransactionService {
	return &TransactionService{
		repository: repository,
	}
}
