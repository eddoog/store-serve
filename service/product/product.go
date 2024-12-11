package product

import (
	"errors"
	"strconv"
	"strings"

	"github.com/eddoog/store-serve/domains/entities"
	"github.com/eddoog/store-serve/domains/models"
)

func (p *ProductService) GetProducts(params entities.SearchProduct) (entities.SearchProductResponse, error) {
	var categoryIDs []uint
	if params.Category != "" {
		for _, idStr := range strings.Split(params.Category, ",") {
			id, err := strconv.Atoi(idStr)
			if err == nil {
				categoryIDs = append(categoryIDs, uint(id))
			}
		}
	}

	params.CategoryIDs = categoryIDs

	products, err := p.ProductRepository.GetProducts(params)
	if err != nil {
		return entities.SearchProductResponse{}, err
	}
	return products, nil
}

func (p *ProductService) GetProductByID(id uint) (models.Product, error) {
	product, err := p.ProductRepository.GetProductByID(id)

	if err != nil {
		if err.Error() == "record not found" {
			return models.Product{}, errors.New("product not found")
		}

		return models.Product{}, err
	}

	return product, nil
}
