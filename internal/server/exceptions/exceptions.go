package exceptions

import "errors"

var (
	ErrNotFound            = errors.New("not found")
	ErrAlreadyExists       = errors.New("already exists")
	ErrInvalidOTP          = errors.New("invalid OTP")
	ErrOTPExpired          = errors.New("OTP expired")
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrTokenNotFound       = errors.New("token not found")
)
