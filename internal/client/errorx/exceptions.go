package errorx

import "errors"

var (
	ErrUserAlreadyExists   = errors.New("user already exists")
	ErrEmailInvalidFormat  = errors.New("invalid email format")
	ErrPasswordTooShort    = errors.New("password is too short")
	ErrOtpExpired          = errors.New("OTP code has expired")
	ErrOtpInvalid          = errors.New("invalid OTP code")
	ErrUserNotFound        = errors.New("user not found")
	ErrTokenNotFound       = errors.New("token not found")
	ErrRefreshTokenInvalid = errors.New("invalid refresh token")
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrNotFound            = errors.New("not found")
)
