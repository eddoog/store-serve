package app

import (
	"github.com/eddoog/store-serve/app/config"
	"github.com/eddoog/store-serve/app/rest"
)

func StartApp() {
	environment := initEnvironment()
	db = initDatabase()

	if environment == "development" {
		models := initModels()

		migrateModels(db, models)
	}

	InitRedis()
	cacheService := config.InitCacheService(rdb)
	repository := config.InitRepository(db, cacheService)
	service := config.InitService(repository)

	rest.InitRest(service, cacheService)
}
