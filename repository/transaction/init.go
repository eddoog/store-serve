package transaction

import (
	"github.com/eddoog/store-serve/domains/models"
	"github.com/eddoog/store-serve/service/cache"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ITransactionRepository interface {
	GetUserTransactions(userID uint) ([]models.Transaction, error)
	Checkout(ctx *fiber.Ctx, userID uint) error
	CancelTransaction(ctx *fiber.Ctx, txID uint, userID uint) error
	GetTransaction(txID uint) (*models.Transaction, error)
	UpdateTransaction(transaction *models.Transaction) error
}

type TransactionRepository struct {
	db           *gorm.DB
	CacheService cache.ICacheService
}

func InitTransactionRepository(
	db *gorm.DB,
	cacheService cache.ICacheService,
) ITransactionRepository {
	return &TransactionRepository{
		db:           db,
		CacheService: cacheService,
	}
}
