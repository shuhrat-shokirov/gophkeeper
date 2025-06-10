package binary

import "time"

type Data struct {
	UserID    int64
	Title     string
	Content   string
	Note      string
	CreatedAt time.Time
}
