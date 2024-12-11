package cart

import (
	"github.com/eddoog/store-serve/service/cart"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ICartController interface {
	Store(ctx *fiber.Ctx) error
	ViewCart(ctx *fiber.Ctx) error
	RemoveCartItem(ctx *fiber.Ctx) error
}

type CartController struct {
	CartService cart.ICartService
	Validator   *validator.Validate
}

func NewCartController(cartService cart.ICartService) ICartController {
	return &CartController{
		CartService: cartService,
		Validator:   validator.New(),
	}
}
