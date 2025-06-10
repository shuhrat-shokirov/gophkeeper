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

type Pagination struct {
	Limit  int64
	Offset int64
}

const limitDefault = 4

func (p *Pagination) GetLimit() int64 {
	if p.Limit <= 0 {
		return limitDefault
	}

	return int64(p.Limit)
}

func (p *Pagination) GetOffset() int64 {
	if p.Limit <= 0 {
		return 0
	}

	return int64(p.Offset * p.Limit)
}

type List struct {
	ID        int64
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Info struct {
	ID int64
	LoginData
	UpdatedAt time.Time
}
