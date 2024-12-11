package rest

import (
	"github.com/eddoog/store-serve/controller"
	"github.com/eddoog/store-serve/middleware"
	"github.com/gofiber/fiber/v2"
)

type Routes struct {
	controller *controller.Controller
}

func initRoutes(controller *controller.Controller) *Routes {
	return &Routes{controller: controller}
}

func (r *Routes) registerRoutes() {
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	r.AuthRoutes()
	r.UserRoutes()
	r.ProductRoutes()
	r.CartRoutes()
	r.TransactionRoutes()
}

func (r *Routes) AuthRoutes() {
	authGroup := app.Group("/auth")

	authGroup.Post("/login", r.controller.AuthController.Login)
	authGroup.Post("/register", r.controller.AuthController.Register)
}

func (r *Routes) UserRoutes() {
	userGroup := app.Group("/user")

	userGroup.Get("/me", middleware.JWTMiddleware, r.controller.UserController.Profile)
}

func (r *Routes) ProductRoutes() {
	productGroup := app.Group("/products").Use(middleware.JWTMiddleware)

	productGroup.Get("/", r.controller.ProductController.Index)
	productGroup.Get("/:id", r.controller.ProductController.Show)
}

func (r *Routes) CartRoutes() {
	cartGroup := app.Group("/cart").Use(middleware.JWTMiddleware)

	cartGroup.Get("/", r.controller.CartController.ViewCart)
	cartGroup.Post("/", r.controller.CartController.Store)
	cartGroup.Delete("/:id", r.controller.CartController.RemoveCartItem)
}

func (r *Routes) TransactionRoutes() {
	transactionGroup := app.Group("/transaction").Use(middleware.JWTMiddleware)

	transactionGroup.Get("/", r.controller.TransactionController.GetTransactions)
	transactionGroup.Post("/checkout", r.controller.TransactionController.Checkout)
	transactionGroup.Delete("/:id/cancel", r.controller.TransactionController.CancelTransaction)
	transactionGroup.Post("/:id/pay", r.controller.TransactionController.HandlePayment)
}
