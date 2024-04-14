package models

import "time"

type User struct {
	ID            uint `gorm:"primary_key"`
	AccountNumber string
	FullName      string
	Email         string
	Password      string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
