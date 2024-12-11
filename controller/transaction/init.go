package transaction

import (
	"github.com/eddoog/store-serve/service/transaction"
	"github.com/gofiber/fiber/v2"
)

type ITransactionController interface {
	GetTransactions(ctx *fiber.Ctx) error
	Checkout(ctx *fiber.Ctx) error
	CancelTransaction(ctx *fiber.Ctx) error
}

type TransactionController struct {
	TransactionService transaction.ITransactionService
}

func NewTransactionController(
	transaction transaction.ITransactionService,
) ITransactionController {
	return &TransactionController{
		TransactionService: transaction,
	}
}
