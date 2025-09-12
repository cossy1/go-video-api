package database

import (
	"go-api/entity"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")

	if dsn == "" {
		log.Fatal("Database url is not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&entity.User{}, &entity.Video{})

	DB = db
	return DB
}
