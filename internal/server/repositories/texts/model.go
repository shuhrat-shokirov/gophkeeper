package texts

import "time"

type TextData struct {
	UserID    int64
	Title     string
	Content   string
	Note      string
	CreatedAt time.Time
}
