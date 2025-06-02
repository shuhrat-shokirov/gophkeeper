package utils

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"html/template"
	"math/big"
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
	// используем html/template вместо text/template для безопасности и корректной обработки HTML
	parse, err := template.New("otp").Parse(templateHTML)
	if err != nil {
		return "", fmt.Errorf("error parsing OTP template: %w", err)
	}

	data := struct {
		OTP string
	}{
		OTP: otp,
	}

	var tmpl bytes.Buffer
	if err := parse.Execute(&tmpl, data); err != nil {
		return "", fmt.Errorf("error executing OTP template: %w", err)
	}

	return tmpl.String(), nil
}

var templateHTML = `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>One-Time Password</title>
</head>
<body style="font-family: Arial, sans-serif; background-color: #f9f9f9; padding: 20px;">
	<div style="max-width: 600px; margin: auto; background-color: #fff; padding: 20px; border-radius: 8px; box-shadow: 0 2px 4px rgba(0,0,0,0.1);">
		<h1 style="color: #333;">Your One-Time Password</h1>
		<p style="font-size: 16px;">Your OTP is: <strong>{{.OTP}}</strong></p>
		<p style="font-size: 14px; color: #555;">This OTP is valid for 5 minutes.</p>
		<p style="font-size: 14px; color: #555;">Please do not share this OTP with anyone.</p>
		<p style="font-size: 14px; color: #555;">Thank you for using our service!</p>
	</div>
</body>
</html>`
