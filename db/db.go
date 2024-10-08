package db

import (
	"golang-rest-api/utils"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	logger := utils.Logger()
	logger.Info("Database connecting...")

	dsn := os.Getenv("DB_URL")
	logger.Debug("Connecting to " + dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	logger.Info("Database Connected")
	return db
}

func CloseDB(db *gorm.DB) {
	logger := utils.Logger()

	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		logger.Fatal(err.Error())
	}
}
