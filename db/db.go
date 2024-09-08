package db

import (
	"golang-rest-api/utils"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	logger := utils.NewLogger()

	logger.Info("Starting database connction...")
	// Retrieve enviroment variables from .env file only in the development enviroment.
	if os.Getenv("GO_ENV") == "development" {
		logger.Info("Read .env file")
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
	}

	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	logger.Info("Database Connected")
	return db
}

func CloseDB(db *gorm.DB) {
	logger := utils.NewLogger()

	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		logger.Fatal(err.Error())
	}
}
