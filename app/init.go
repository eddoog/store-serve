package app

import (
	"os"

	"github.com/eddoog/store-serve/domains/models"
	"github.com/joho/godotenv"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initEnvironment() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.New().Info("Environment variables loaded")
}

func initDatabase() *gorm.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	logrus.New().Info("Database connected")

	return db
}

func tableExists(db *gorm.DB, tableName string) bool {
	var exists bool
	query := `SELECT EXISTS (
		SELECT 1 
		FROM information_schema.tables 
		WHERE table_name = ?
	)`
	err := db.Raw(query, tableName).Scan(&exists).Error
	if err != nil {
		logrus.Fatalf("Error checking if table exists: %v", err)
	}
	return exists
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
		tableName := getTableName(db, model)
		if tableExists(db, tableName) {
			logrus.Printf("Table '%s' already exists. Skipping migration.", tableName)
		} else {
			logrus.Printf("Table '%s' does not exist. Running migration.", tableName)
			err := db.AutoMigrate(model)
			if err != nil {
				logrus.Fatalf("Failed to migrate table '%s': %v", tableName, err)
			}
			logrus.Printf("Migration for table '%s' completed successfully.", tableName)
		}
	}
}

func initModels() []interface{} {
	return []interface{}{
		&models.User{},
		&models.ProductCategory{},
		&models.Product{},
	}
}
