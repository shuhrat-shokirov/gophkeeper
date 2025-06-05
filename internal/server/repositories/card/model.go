package card

import "time"

type Data struct {
	UserID    int64
	Pan       string
	Cvv       string
	Expiry    string
	Title     string
	Note      string
	CreatedAt time.Time
}
