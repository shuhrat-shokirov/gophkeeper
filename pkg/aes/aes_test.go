package aes

import (
	"reflect"
	"testing"
)

func Test_decrypt(t *testing.T) {
	type args struct {
		key        []byte
		cipherText []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "valid decryption",
			args: args{
				key:        []byte("0123456789abcdef"), // 16 bytes key for AES-128
				cipherText: []byte(`test message 123 123`),
			},
			want:    []byte{163, 66, 138, 193},
			wantErr: false,
		},
		{
			name: "invalid key length",
			args: args{
				key:        []byte("shortkey"), // Invalid key length for AES
				cipherText: []byte(`test message 123 123`),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid cipher text",
			args: args{
				key: []byte("0123456789abcdef"), // 16 bytes key for AES-128
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty cipher text",
			args: args{
				key:        []byte("0123456789abcdef"), // 16 bytes key for AES-128
				cipherText: []byte{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "nil key",
			args: args{
				key:        nil,
				cipherText: []byte(`test message 123 123`),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "nil cipher text",
			args: args{
				key:        []byte("0123456789abcdef"), // 16 bytes key for AES-128
				cipherText: nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decrypt(tt.args.key, tt.args.cipherText)
			if (err != nil) != tt.wantErr {
				t.Errorf("decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encrypt(t *testing.T) {
	type args struct {
		key       []byte
		plainText []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid encryption",
			args: args{
				key:       []byte("0123456789abcdef"), // 16 bytes key for AES-128
				plainText: []byte("test message 123 123"),
			},
			wantErr: false,
		},
		{
			name: "invalid key length",
			args: args{
				key:       []byte("shortkey"), // Invalid key length for AES
				plainText: []byte("test message 123 123"),
			},
			wantErr: true,
		},
		{
			name: "empty plain text",
			args: args{
				key:       []byte("0123456789abcdef"), // 16 bytes key for AES-128
				plainText: []byte{},
			},
			wantErr: false,
		},
		{
			name: "nil key",
			args: args{
				key:       nil,
				plainText: []byte("test message 123 123"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := encrypt(tt.args.key, tt.args.plainText)
			if (err != nil) != tt.wantErr {
				t.Errorf("encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(got) == 0 {
				t.Errorf("encrypt() got = %v, want non-empty result", got)
			}
		})
	}
}

func TestMustEncrypt(t *testing.T) {
	type args struct {
		plainText string
		secretKey string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "valid encryption and decryption",
			args: args{
				plainText: "test message 123 123",
				secretKey: "0123456789abcdef", // 16 bytes key for AES-128
			},
		},
		{
			name: "empty plain text",
			args: args{
				plainText: "",
				secretKey: "0123456789abcdef", // 16 bytes key for AES-128
			},
		},
		{
			name: "invalid key length",
			args: args{
				plainText: "test message 123 123",
				secretKey: "shortkey", // Invalid key length for AES
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set the secret key globally for the test
			secretKey = tt.args.secretKey

			got := MustEncrypt(tt.args.plainText)
			if got == "" {
				t.Errorf("MustEncrypt() returned empty string, expected non-empty result")
			}

			t.Logf("MustEncrypt() got = %v", got)
		})
	}
}

func TestMustDecrypt(t *testing.T) {
	type args struct {
		b64CipherText string
		secretKey     string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "valid decryption",
			args: args{
				b64CipherText: "OXc9pK297DNXohUrFKBJXpLB1iLwhuFkjznYZqIbnd5E6DgF", // Base64 encoded "test message 123 123"
				secretKey:     "0123456789abcdef",                                 // 16 bytes key for AES-128
			},
			want: "test message 123 123",
		},
		{
			name: "valid decr",
			args: args{
				b64CipherText: "upLsffPsLAO9xtUCvVYhhccfYub0I90JZylSsxIRQnPvRd5K", // Base64 encoded "test message 123 123"
				secretKey:     "0123456789abcdef",                                 // 16 bytes key for AES-128
			},
			want: "test message 123 123",
		},
		{
			name: "invalid secret key",
			args: args{
				b64CipherText: "OXc9pK297DNXohUrFKBJXpLB1iLwhuFkjznYZqIbnd5E6DgF", // Base64 encoded "test message 123 123"
				secretKey:     "shortkey",                                         // Invalid key length for AES
			},
			want: "OXc9pK297DNXohUrFKBJXpLB1iLwhuFkjznYZqIbnd5E6DgF", // Should return the original base64 string
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set the secret key globally for the test
			secretKey = tt.args.secretKey

			if got := MustDecrypt(tt.args.b64CipherText); got != tt.want {
				t.Errorf("MustDecrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
