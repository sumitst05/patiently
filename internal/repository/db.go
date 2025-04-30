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
	if err := db.AutoMigrate(&models.User{}, &models.Patient{}, &models.RegistrationHistory{}); err != nil {
		log.Fatal("Auto migration failed:", err)
	}

	// re-add foregin key for patient after migration
	err = db.Exec(`
		ALTER TABLE registration_histories
		DROP CONSTRAINT IF EXISTS fk_registration_histories_patient;
	`).Error
	if err != nil {
		log.Fatal("Failed to drop foreign key constraint:", err)
	}

	err = db.Exec(`
		ALTER TABLE registration_histories
		ADD CONSTRAINT fk_registration_histories_patient
		FOREIGN KEY (patient_id) REFERENCES patients(id) ON DELETE CASCADE;
	`).Error
	if err != nil {
		log.Fatal("Failed to add foreign key constraint:", err)
	}

	DB = db
	log.Println("Connected to PostgreSQL")
}
