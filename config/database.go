package config

import (
	"fmt"
	"log"

	"github.com/haris97m/go-fiber/model/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase(config Config) *gorm.DB {
	DB_HOST := config.Get("DB_HOST")
	DB_PORT := config.Get("DB_PORT")
	DB_USER := config.Get("DB_USER")
	DB_PASS := config.Get("DB_PASS")
	DB_NAME := config.Get("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database")
	}
	fmt.Println("Connected to database")

	if err := db.AutoMigrate(&domain.Author{}, &domain.Book{}, &domain.Publisher{}); err != nil {
		log.Fatal("Failed to migrate database")
	}
	fmt.Println("Database was migrated")

	return db
}
