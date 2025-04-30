package models

import "time"

type RegistrationHistory struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	PatientID   uint      `gorm:"not null" json:"-"`
	Patient     Patient   `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE;" json:"patient"`
	Action      string    `gorm:"not null" json:"action"` // create or update
	ChangedByID uint      `gorm:"not null" json:"-"`
	ChangedBy   User      `gorm:"foreignKey:ChangedByID" json:"changed_by"`
	Timestamp   time.Time `json:"timestamp"`
	OldValue    string    `gorm:"type:text" json:"old_value"`
	NewValue    string    `gorm:"type:text" json:"new_value"`
}
