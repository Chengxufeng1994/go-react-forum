package model

import (
	"time"
)

type User struct {
	ID        int    `gorm:"primaryKey"`
	Username  string `gorm:"not null"`
	Email     string `gorm:"unique;not null;index:email_idx"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
