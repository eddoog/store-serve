package product

import (
	"github.com/eddoog/store-serve/service/product"
	"github.com/gofiber/fiber/v2"
)

type IProductController interface {
	Index(ctx *fiber.Ctx) error
	Show(ctx *fiber.Ctx) error
}

type ProductController struct {
	ProductService product.IProductService
}

func NewProductController(productService product.IProductService) IProductController {
	return &ProductController{
		ProductService: productService,
	}
}
