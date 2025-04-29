package repository

import (
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/sumitst05/patiently/internal/config"
	"github.com/sumitst05/patiently/internal/models"
)

var DB *gorm.DB

func InitDB() {
	_ = godotenv.Load()

	cfg := config.LoadDBConfig()
	dsn := cfg.DSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// auto-migrations
	if err := db.AutoMigrate(&models.User{}, &models.Patient{}); err != nil {
		log.Fatal("Auto migration failed:", err)
	}

	DB = db
	log.Println("Connected to PostgreSQL")
}
