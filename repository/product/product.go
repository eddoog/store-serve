package product

import (
	"github.com/eddoog/store-serve/domains/entities"
	"github.com/eddoog/store-serve/domains/models"
	"github.com/sirupsen/logrus"
)

func (p *ProductRepository) GetProducts(params entities.SearchProduct) (entities.SearchProductResponse, error) {
	offset := (params.Page - 1) * params.Limit

	var products []models.Product

	query := p.db.Model(&models.Product{}).Preload("Category").Joins("JOIN product_category ON product.category_id = product_category.id")

	if len(params.CategoryIDs) > 0 {
		query = query.Where("product.category_id IN ?", params.CategoryIDs)
	}

	if params.Search != "" {
		query = query.Where("product.name LIKE ?", "%"+params.Search+"%")
	}

	if params.MinPrice > 0 {
		query = query.Where("product.price >= ?", params.MinPrice)
	}

	if params.MaxPrice > 0 {
		query = query.Where("product.price <= ?", params.MaxPrice)
	}

	var totalItems int64
	query.Count(&totalItems)
	query.Offset(offset).Limit(params.Limit).Find(&products)

	if query.Error != nil {
		logrus.Error(query.Error)
		return entities.SearchProductResponse{}, nil
	}

	logrus.Info(query.Statement.SQL.String())

	var productResponse []entities.ProductResponse
	for _, product := range products {
		productResponse = append(productResponse, entities.ProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			CategoryID:  product.CategoryID,
			Category: entities.CategoryResponse{
				ID:          product.Category.ID,
				Name:        product.Category.Name,
				Description: product.Category.Description,
			},
			Stock:     product.Stock,
			CreatedAt: product.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: product.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return entities.SearchProductResponse{
		TotalItems: totalItems,
		Data:       productResponse,
	}, nil

}

func (p *ProductRepository) GetProductByID(id uint) (models.Product, error) {
	var product models.Product
	query := p.db.Preload("Category").First(&product, id)
	if query.Error != nil {
		logrus.Error(query.Error)
		return models.Product{}, query.Error
	}
	return product, nil
}
