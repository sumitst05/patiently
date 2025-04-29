package models

import (
	"time"
)

type Patient struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Age         int    `gorm:"not null"`
	Gender      string `gorm:"type:VARCHAR(10)"`
	Address     string
	Phone       string `gorm:"type:VARCHAR(15)"`
	CreatedByID uint   `gorm:"not null"` // receptionist who registered the patient
	CreatedBy   User   `gorm:"foreignKey:CreatedByID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
