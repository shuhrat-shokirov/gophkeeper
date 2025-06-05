package server

import "time"

type LoginAndPass struct {
	Login      string
	Pass       string
	Title      string
	Note       string
	ModifiedAt time.Time
}

type Text struct {
	Title      string
	Content    string
	Note       string
	ModifiedAt time.Time
}

type Card struct {
	Title      string
	Pan        string
	Cvv        string
	Expiry     string
	Note       string
	ModifiedAt time.Time
}

type Binary struct {
	Title      string
	Content    []byte
	Note       string
	ModifiedAt time.Time
}
