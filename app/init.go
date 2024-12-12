package app

import (
	"os"

	"github.com/eddoog/store-serve/domains/models"
	"github.com/joho/godotenv"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initEnvironment() string {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.New().Info("Environment variables loaded")

	return os.Getenv("ENVIRONMENT")
}

func initDatabase() *gorm.DB {
	pgBouncerURL := os.Getenv("PGBOUNCER_URL")
	if pgBouncerURL == "" {
		panic("PGBOUNCER_URL environment variable is not set")
	}

	db, err := gorm.Open(postgres.Open(pgBouncerURL), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		panic(err)
	}

	session := db.Session(&gorm.Session{PrepareStmt: true})
	logrus.New().Info("Database connected")

	return session
}

func getTableName(db *gorm.DB, model interface{}) string {
	stmt := &gorm.Statement{DB: db}
	err := stmt.Parse(model)
	if err != nil {
		logrus.Fatalf("Failed to parse model: %v", err)
	}
	return stmt.Schema.Table
}

func migrateModels(db *gorm.DB, models []interface{}) {
	for _, model := range models {
		err := db.AutoMigrate(model)
		if err != nil {
			logrus.Fatalf("Failed to migrate table '%s': %v", getTableName(db, model), err)
		}
		logrus.Printf("Migration for table '%s' completed successfully.", getTableName(db, model))
	}
}

func initModels() []interface{} {
	return []interface{}{
		&models.User{},
		&models.ProductCategory{},
		&models.Product{},
		&models.Cart{},
		&models.CartItem{},
		&models.Transaction{},
		&models.TransactionItem{},
	}
}
