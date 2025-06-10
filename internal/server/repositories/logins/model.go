package logins

import "time"

type LoginData struct {
	UserID    int64
	Login     string
	Password  string
	Title     string
	Note      string
	CreatedAt time.Time
}
