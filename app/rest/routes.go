package rest

import (
	"github.com/eddoog/store-serve/controller"
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
}

func (r *Routes) AuthRoutes() {
	authGroup := app.Group("/auth")

	authGroup.Post("/login", r.controller.AuthController.Login)
	authGroup.Post("/register", r.controller.AuthController.Register)
}
