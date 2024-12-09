package rest

import (
	"fmt"
	"os"

	"github.com/eddoog/store-serve/controller"
	"github.com/eddoog/store-serve/service"
	"github.com/gofiber/fiber/v2"
)

func InitRest(service *service.Service) {
	app = fiber.New()

	controller := controller.InitController(service)

	controllerList := initRoutes(controller)

	controllerList.registerRoutes()

	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
