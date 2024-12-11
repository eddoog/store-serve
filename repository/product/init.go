package product

import (
	"github.com/eddoog/store-serve/domains/entities"
	"github.com/eddoog/store-serve/domains/models"
	"gorm.io/gorm"
)

type IProductRepository interface {
	GetProducts(params entities.SearchProduct) (entities.SearchProductResponse, error)
	GetProductByID(id uint) (models.Product, error)
}

type ProductRepository struct {
	db *gorm.DB
}

func InitProductRepository(db *gorm.DB) IProductRepository {
	return &ProductRepository{
		db: db,
	}
}
