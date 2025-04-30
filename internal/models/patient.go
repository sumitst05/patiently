package models

import (
	"time"
)

type Patient struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Age         int       `gorm:"not null" json:"age"`
	Gender      string    `gorm:"type:VARCHAR(10)" json:"gender"`
	Address     string    `gorm:"type:text" json:"address"`
	Phone       string    `gorm:"type:VARCHAR(15)" json:"phone"`
	CreatedByID uint      `gorm:"not null" json:"-"` // receptionist who registered the patient
	CreatedBy   User      `gorm:"foreignKey:CreatedByID" json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
