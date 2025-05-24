package utils

import (
	"crypto/rand"
	"fmt"
	"html/template"
	"math/big"

	"go.uber.org/zap/buffer"
)

func GenerateOTP() (string, error) {
	const otpLength = 6
	const digits = "0123456789"

	otp := make([]byte, otpLength)
	for i := range otp {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return "", fmt.Errorf("failed to generate random index: %w", err)
		}

		otp[i] = digits[randomIndex.Int64()]
	}

	return string(otp), nil
}

func GenerateOTPMessage(otp string) (string, error) {
	parse, err := template.New("otp").Parse(templateHTML)
	if err != nil {
		return "", fmt.Errorf("error parsing OTP template: %w", err)
	}

	data := struct {
		OTP string
	}{
		OTP: otp,
	}

	var tmpl buffer.Buffer
	if err := parse.Execute(&tmpl, data); err != nil {
		return "", fmt.Errorf("error executing OTP template: %w", err)
	}

	return tmpl.String(), nil
}

var templateHTML = `<html>
<head>
	<title>One-Time Password</title>
</head>
<body>
	<h1>Your One-Time Password</h1>
	<p>Your OTP is: {{.OTP}}</p>
	<p>This OTP is valid for 5 minutes.</p>
	<p>Please do not share this OTP with anyone.</p>
	<p>Thank you for using our service!</p>
</body>
</html>`
