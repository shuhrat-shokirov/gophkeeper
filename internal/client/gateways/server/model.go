package server

import "time"

type LoginAndPass struct {
	Login     string
	Pass      string
	Title     string
	Note      string
	CreatedAt time.Time
}

type Text struct {
	Title     string
	Content   string
	Note      string
	CreatedAt time.Time
}

type Card struct {
	Title     string
	Pan       string
	Cvv       string
	Expiry    string
	Note      string
	CreatedAt time.Time
}

type Binary struct {
	Title     string
	Content   []byte
	Note      string
	CreatedAt time.Time
}

type ListItem struct {
	ID        int64
	Title     string
	CreatedAt int64
	UpdatedAt int64
}

type LoginInfo struct {
	ID int64
	LoginAndPass
	UpdatedAt time.Time
}

type TextInfo struct {
	ID int64
	Text
	UpdatedAt time.Time
}

type CardInfo struct {
	ID int64
	Card
	UpdatedAt time.Time
}

type BinaryInfo struct {
	ID int64
	Binary
	UpdatedAt time.Time
}
