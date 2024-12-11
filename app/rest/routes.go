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
