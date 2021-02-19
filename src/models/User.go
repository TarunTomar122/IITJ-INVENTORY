package models

import (
	"github.com/jinzhu/gorm"
)

// User Model
type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"primaryKey;not null"`
	Password string `gorm:"not null"`
}