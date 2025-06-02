package emailtotp

type Request struct {
	To      string
	Subject string
	Body    string
}
