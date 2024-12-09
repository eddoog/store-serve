package app

import (
	"github.com/eddoog/store-serve/app/config"
	"github.com/eddoog/store-serve/app/rest"
	"github.com/eddoog/store-serve/repository"
)

func StartApp() {
	initEnvironment()

	repository := repository.InitRepository()
	service := config.InitService(repository)

	rest.InitRest(service)
}
