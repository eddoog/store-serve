package app

import (
	"github.com/eddoog/store-serve/app/config"
	"github.com/eddoog/store-serve/app/rest"
)

func StartApp() {
	initEnvironment()
	db = initDatabase()

	models := initModels()

	migrateModels(db, models)

	repository := config.InitRepository(db)
	service := config.InitService(repository)

	rest.InitRest(service)
}
