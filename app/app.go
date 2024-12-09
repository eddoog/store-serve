package app

import (
	"github.com/eddoog/store-serve/app/rest"
	"github.com/eddoog/store-serve/repository"
	"github.com/eddoog/store-serve/service"
)

func StartApp() {
	initEnvironment()

	repository := repository.InitRepository()
	service := service.InitService(repository)

	rest.InitRest(service)
}
