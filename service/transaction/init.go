package transaction

import (
	"github.com/eddoog/store-serve/domains/models"
	"github.com/eddoog/store-serve/repository/transaction"
	"github.com/gofiber/fiber/v2"
)

type ITransactionService interface {
	GetUserTransactions(userID uint) ([]models.Transaction, error)
	Checkout(ctx *fiber.Ctx, userID uint) error
	CancelTransaction(ctx *fiber.Ctx, txID uint, userID uint) error
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
