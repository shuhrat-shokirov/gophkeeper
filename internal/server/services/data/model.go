package data

type LoginData struct {
	UserId    int64
	Login     string
	Password  string
	Title     string
	Note      string
	CreatedAt int64
}

type LoginInfo struct {
	ID int64
	LoginData
	UpdatedAt int64
}

type TextData struct {
	UserID    int64
	Title     string
	Content   string
	Note      string
	CreatedAt int64
}

type CardData struct {
	UserID    int64
	Pan       string
	Cvv       string
	Expiry    string
	Title     string
	Note      string
	CreatedAt int64
}

type BinaryData struct {
	UserID    int64
	Title     string
	Content   []byte
	Note      string
	CreatedAt int64
}

type LoginListItem struct {
	ID        int64
	Title     string
	CreatedAt int64
	UpdatedAt int64
}
