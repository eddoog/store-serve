package product

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/eddoog/store-serve/domains/entities"
	"github.com/eddoog/store-serve/domains/models"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func (p *ProductController) Index(ctx *fiber.Ctx) error {
	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Query("limit", "10"))
	minPrice, _ := strconv.ParseFloat(ctx.Query("min_price", "0"), 64)
	maxPrice, _ := strconv.ParseFloat(ctx.Query("max_price", "0"), 64)

	searchParams := entities.SearchProduct{
		Category: ctx.Query("category"),
		Page:     page,
		Limit:    limit,
		Search:   ctx.Query("search"),
		MinPrice: minPrice,
		MaxPrice: maxPrice,
	}

	cacheKey := fmt.Sprintf("products:index:page_%d:limit_%d:category_%s:search_%s:minPrice_%f:maxPrice_%f",
		page, limit, searchParams.Category, searchParams.Search, minPrice, maxPrice)

	var cachedProducts entities.SearchProductResponse
	if err := p.CacheService.Get(ctx, cacheKey, &cachedProducts); err == nil {
		return ctx.JSON(fiber.Map{
			"current_page": page,
			"total_items":  cachedProducts.TotalItems,
			"data":         cachedProducts.Data,
		})
	}

	products, err := p.ProductService.GetProducts(searchParams)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := p.CacheService.Set(ctx, cacheKey, products, 5*time.Minute); err != nil {
		logrus.Printf("Failed to set cache: %v", err)
	}

	return ctx.JSON(fiber.Map{
		"current_page": page,
		"total_items":  products.TotalItems,
		"data":         products.Data,
	})
}

func (p *ProductController) Show(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	cacheKey := fmt.Sprintf("product:%d", id)

	var cachedProduct models.Product
	if err := p.CacheService.Get(ctx, cacheKey, &cachedProduct); err == nil {
		// Return cached product
		return ctx.JSON(fiber.Map{
			"data": cachedProduct,
		})
	}

	// Fetch the product from the database if not found in cache
	product, err := p.ProductService.GetProductByID(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := p.CacheService.Set(ctx, cacheKey, product, 5*time.Minute); err != nil {
		log.Printf("Failed to set cache: %v", err)
	}

	return ctx.JSON(fiber.Map{
		"data": product,
	})
}
