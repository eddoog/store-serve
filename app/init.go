package app

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func initEnvironment() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.New().Info("Environment variables loaded")
}
