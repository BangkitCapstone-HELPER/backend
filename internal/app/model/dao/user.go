package dao

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// ID          uuid.UUID `gorm:"index;primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name        string `gorm:"not null;size:255"`
	Email       string `gorm:"index;unique;not null"`
	Password    string `gorm:"not null;size:255"`
	PhoneNumber string `gorm:"size:255"`
	IsAdmin     bool   `gorm:"default:0"`
	Address     string `gorm:"not null;size:255"`
}
