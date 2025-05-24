package session

import "time"

type Session struct {
	ID           int
	UserID       int
	RefreshToken string
	CreatedAt    time.Time
	ExpiresAt    time.Time
	UpdatedAt    time.Time
}
