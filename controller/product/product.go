package product

import (
	"strconv"

	"github.com/eddoog/store-serve/domains/entities"
	"github.com/gofiber/fiber/v2"
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

	products, err := p.ProductService.GetProducts(searchParams)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"current_page": page,
		"total_items":  products.TotalItems,
		"data":         products.Data,
	})
}

func (p *ProductController) Show(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	product, err := p.ProductService.GetProductByID(uint(id))

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"data": product,
	})
}
