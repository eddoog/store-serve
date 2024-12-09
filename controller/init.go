package controller

import "github.com/eddoog/store-serve/service"

type Controller struct {
}

func InitController(service *service.Service) *Controller {
	return &Controller{}
}
