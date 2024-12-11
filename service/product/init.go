package product

import (
	"github.com/eddoog/store-serve/domains/entities"
	"github.com/eddoog/store-serve/domains/models"
	"github.com/eddoog/store-serve/repository/product"
)

type IProductService interface {
	GetProducts(params entities.SearchProduct) (entities.SearchProductResponse, error)
	GetProductByID(id uint) (models.Product, error)
}

type ProductService struct {
	ProductRepository product.IProductRepository
}

func InitProductService(productRepository product.IProductRepository) IProductService {
	return &ProductService{
		ProductRepository: productRepository,
	}
}
