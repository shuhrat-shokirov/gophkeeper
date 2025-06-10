package data

import "time"

type LoginData struct {
	Login string
	Pass  string
	Title string
	Note  string
}

type TextData struct {
	Title   string
	Content string
	Note    string
}

type CardData struct {
	Title  string
	Pan    string
	Cvv    string
	Expiry string
	Note   string
}

type FileData struct {
	Title string
	Path  string
	Note  string
}

type ListItem struct {
	ID        int64
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type LoginInfo struct {
	ID int64
	LoginData
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TextInfo struct {
	ID int64
	TextData
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CardInfo struct {
	ID int64
	CardData
	CreatedAt time.Time
	UpdatedAt time.Time
}

type BinaryInfo struct {
	ID        int64
	Title     string
	Content   []byte
	Note      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
