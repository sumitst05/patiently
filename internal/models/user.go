package models

import "time"

type User struct {
	Id           uint   `gorm:"primaryKey"`
	Name         string `gorm:"not null"`
	Email        string `gorm:"uniqueIndex;not null"`
	PasswordHash string `gorm:"not null"`
	Role         string `gorm:"type:VARCHAR(20);not null"` // doctor or receptionist
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
