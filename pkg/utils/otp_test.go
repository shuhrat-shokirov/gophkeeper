package utils

import (
	"strings"
	"testing"
)

func TestGenerateOTP(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "valid OTP generation",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GenerateOTP()
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateOTP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGenerateOTPMessage(t *testing.T) {
	type args struct {
		otp string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "valid OTP message",
			args: args{
				otp: "123456",
			},
			want:    "Your OTP is 123456",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateOTPMessage(tt.args.otp)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateOTPMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == "" {
				t.Errorf("GenerateOTPMessage() got = %v, want non-empty string", got)
				return
			}

			if !strings.Contains(got, tt.args.otp) {
				t.Errorf("GenerateOTPMessage() got = %v, want to contain OTP %v", got, tt.args.otp)
			}
		})
	}
}
