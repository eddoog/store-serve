package cart

import (
	"strconv"

	"github.com/eddoog/store-serve/domains/entities"
	"github.com/eddoog/store-serve/middleware"
	"github.com/gofiber/fiber/v2"
)

func (c *CartController) Store(ctx *fiber.Ctx) error {
	var req entities.AddToCartRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	if err := c.Validator.Struct(req); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  err.Error(),
		})
	}

	user := ctx.Locals("user").(middleware.UserClaims)

	if user == (middleware.UserClaims{}) {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	if err := c.CartService.AddToCart(user.UserID, req); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to add to cart",
			"error":   err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Product added to cart successfully",
	})
}

func (c *CartController) ViewCart(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(middleware.UserClaims)
	if user == (middleware.UserClaims{}) {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	items, total, err := c.CartService.GetCart(user.UserID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve cart items",
			"error":   err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Cart retrieved successfully",
		"data":    items,
		"total":   total,
	})
}

func (c *CartController) RemoveCartItem(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(middleware.UserClaims)
	if user == (middleware.UserClaims{}) {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	clearAll := ctx.Query("all") == "true"
	var productID *uint
	var quantity int

	if !clearAll {
		productIDParam := ctx.Query("product_id")
		quantityParam := ctx.Query("quantity")

		if productIDParam != "" {
			id, err := strconv.Atoi(productIDParam)
			if err != nil {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"message": "Invalid product ID",
				})
			}
			idUint := uint(id)
			productID = &idUint
		}

		if quantityParam != "" {
			qty, err := strconv.Atoi(quantityParam)
			if err != nil || qty <= 0 {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"message": "Invalid quantity",
				})
			}
			quantity = qty
		} else {
			quantity = 1
		}
	}

	err := c.CartService.RemoveCartItem(user.UserID, productID, quantity, clearAll)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Cart item(s) removed successfully",
	})
}
