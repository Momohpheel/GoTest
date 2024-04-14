package models

import "time"

type Transaction struct {
	ID        uint `gorm:"primary_key"`
	Reference string
	AccountID uint
	Account   User
	Amount    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
