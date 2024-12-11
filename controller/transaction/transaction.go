package transaction

import (
	"strconv"

	"github.com/eddoog/store-serve/middleware"
	"github.com/gofiber/fiber/v2"
)

func (t *TransactionController) GetTransactions(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(middleware.UserClaims)
	if user == (middleware.UserClaims{}) {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	transactions, err := t.TransactionService.GetUserTransactions(user.UserID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve transactions",
			"error":   err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"data": transactions,
	})
}

func (t *TransactionController) Checkout(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(middleware.UserClaims)
	if user == (middleware.UserClaims{}) {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	err := t.TransactionService.Checkout(user.UserID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Checkout successful",
	})
}

func (t *TransactionController) CancelTransaction(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(middleware.UserClaims)
	if user == (middleware.UserClaims{}) {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	txIDStr := ctx.Params("id")
	txID, err := strconv.Atoi(txIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid transaction ID",
		})
	}

	if err := t.TransactionService.CancelTransaction(uint(txID), user.UserID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to cancel transaction",
			"error":   err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Transaction canceled successfully",
	})
}

func (t *TransactionController) HandlePayment(ctx *fiber.Ctx) error {
	// Extract user from context
	user := ctx.Locals("user").(middleware.UserClaims)
	if user == (middleware.UserClaims{}) {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// Get transaction ID from URL parameters
	txIDStr := ctx.Params("id")
	txID, err := strconv.Atoi(txIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid transaction ID",
		})
	}

	// Call service to process payment
	err = t.TransactionService.ProcessPayment(uint(txID), user.UserID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Payment processing failed",
			"error":   err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Payment successful",
	})
}
