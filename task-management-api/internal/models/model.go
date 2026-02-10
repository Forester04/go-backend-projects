package models

import (
	"time"

	"gorm.io/gorm"
)

type Status string

const (
	TODO        Status = "todo"
	IN_PROGRESS Status = "in_progress"
	DONE        Status = "done"
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
	Projects  []*Project
	Tasks     []*Task
}

type Project struct {
	gorm.Model
	Name        string  `gorm:"type:varchar(100);not null"`
	Description *string `gorm:"type:text"`
	UserID      uint    `gorm:"not null;index"`
	Tasks       []*Task
}

type Task struct {
	gorm.Model
	Title       string  `gorm:"type:varchar(150);not null"`
	Description *string `gorm:"type:text"`
	Status      Status  `gorm:"type:status;not null;default:'todo'"`
	UserID      uint    `gorm:"not null;index"`
	ProjectID   uint    `gorm:"not null;index"`
}
