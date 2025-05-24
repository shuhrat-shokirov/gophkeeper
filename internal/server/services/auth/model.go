package auth

type Registration struct {
	Email    string
	Password string
}

type ConfirmResponse struct {
	UserId int
	Token  string
}
