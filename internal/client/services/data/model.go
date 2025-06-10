package data

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
