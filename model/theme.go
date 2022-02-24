package model

import "time"

type Theme struct {
	ID          int
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UserID      int
}
