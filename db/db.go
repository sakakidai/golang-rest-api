package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	fmt.Println("Starting database connction...")
	// Retrieve enviroment variables from .env file only in the development enviroment.
	if os.Getenv("GO_ENV") == "dev" {
		fmt.Println("Read .env file")
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

	fmt.Println("Database Connected")
	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}
