package database

import (
	"fmt"
	"go-api/entity"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	password := os.Getenv("DB_PASSWORD")
	dsn := fmt.Sprintf("host=localhost user=postgres password=%s dbname=go-video-api port=5432 sslmode=disable TimeZone=Asia/Shanghai", password)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&entity.User{}, &entity.Video{})

	DB = db
	return DB
}
