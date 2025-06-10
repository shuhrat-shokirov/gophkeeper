package utils

import "testing"

func TestValidateEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid email",
			args: args{
				email: "123@gmail.com",
			},
			want: true,
		},
		{
			name: "invalid email with special characters",
			args: args{
				email: "123@!gmail.com",
			},
			want: false,
		},
		{
			name: "invalid email with spaces",
			args: args{
				email: "123 @gmail.com",
			},
			want: false,
		},
		{
			name: "invalid email with missing domain",
			args: args{
				email: "123@.com",
			},
			want: false,
		},
		{
			name: "invalid email with missing local part",
			args: args{
				email: "@gmail.com",
			},
			want: false,
		},
		{
			name: "invalid email with too short local part",
			args: args{
				email: "@g.com",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateEmail(tt.args.email); got != tt.want {
				t.Errorf("ValidateEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
