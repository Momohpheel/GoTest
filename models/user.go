package models

import "time"

type User struct {
	ID             uint `gorm:"primary_key"`
	FullName       string
	Email          string
	AccountBalance uint
	Password       string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
