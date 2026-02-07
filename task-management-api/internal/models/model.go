package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string `gorm:"type:varchar(100);not null;unique"`
	Password  string `gorm:"not null"`
	Username  string `gorm:"type:varchar(100);unique;not null"`
	FirstName string `gorm:"type:varchar(100);not null"`
	LastName  string `gorm:"type:varchar(100);not null"`
	Phone     *string
	BirthDate *time.Time
}
