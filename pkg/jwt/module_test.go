package jwt

import (
	"encoding/base64"
	"errors"
	"os"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"

	"gophkeeper/pkg/config"
)

const (
	testPrivateKey = "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFb3dJQkFBS0NBUUVBenlLT2NkdXNHSU8vVS82Z0UxbDR0Z3Z4KzR5TUExQ1pQRkYwRnhKbzJaS1ZTY3I4CklONUVVQ1A5VHhLMUUyS3NsZ0tNc0JIZ2JXSWN2RzBQaU12VFo0dTB3SWlPaTVSMDNlK3I5V1NqOG1xSCs3UjUKNVZ3cm1GVWxXTEVsbk9RODJ2L3ljaVdoT2RSVEVieXE2WEFnaU5BckZyL3NFRTBacHFPdlVSeEFmeG5Qb1ZGdwpzODVKZVNBWEdXNmhvRytxaHhCL2d3YmJGTlYraW11YmVGenRyd1NJUGdmNjR2d2RoWnpDY1JZOVRUK1dyRm16CmJObmZjcEszb2RFZzgrM3VSWlg3QW9kdFNhOTk2T0xRTnBTSVFnUDRiMnBXem5hc0NlRHU3ZlBWME5GNkJmSG0KYXArN0R2TkU5NW5XMVFIdjM2UUczaFVUZmdGWUFSaHNDazkwcndJREFRQUJBb0lCQUZ1Wkx2L1h3b2VHdjNYUAovSThCK25rYTJEUkM1Mm5SMnluSzVYa01lWlI1bDQ0dDl3Zzc4bDYwUTZFVHAwSysySTV2NnpJemZaa3hrWDZjCkJnb2JCTTVlQUIxQ1pqTUFnQnZqRUpxd21qV3ArWitNSkhtU3BHNjFmSkgzcUtmMFlKc0NJMzlwOTUzQXNNbC8Kc3Q4UGFEdklQcjNOT29ITTdxSjc4UnYvejkvRVNSdVBuUFZPaEtQSXh6OFZ4Z2p5eGRlemM5Y0NvbnZLbzV0bgpXUGFjOStlYkwvSTFFWDdHZ1IwVnZzbDdtc0RoVVorcnkraUdnM3pBVzhIOGJ5NUFGR3lkdUNjeGtrWmVUVFRFClR1SllVSGdCTWN5NUtCNStxdzNvRk1QVnFNeTE3RktCRncwMW8vQlE3WlFCdERNUFhoWVJyY2RDRk9BOThia1QKbzlnQzFka0NnWUVBNTVxZ0NUUzFibk9xY3lmOFpMdkM2ZFdyVWQ1Q1VZVEZiNWZWUGlYRmxpTTExR2FRbWxGUgpyTmxjTU83ZEVHNVNvUmkxNzJHQVJZK24xZi84QXlGN2g0WHlqanFkbnNybmhFckFJRXNiN1h1cFRzM3RqL1pXCjhJYWFFZ0VxUExITHVrdHhyWjh2aHdtamxlSXhrVmVGR0N1S3JSMFc0bWRlV3ArNmJjYnJ2TmtDZ1lFQTVQUWIKSHI4c2FhdTIrSHVTTm90czl5ZUhNdmRKczkvZzdmeHVzUGs4STRrVHdGZHpZeUhNSmcvdzEwTTY3VFNiYzJFcAp1SXlxSlpCVUNDNkVleUF3cjJERjBiRkJEWGpqZ3hFQzYrMU1QdXlXV2JWam1EaWRreFA4WVRkb0k5ZjVrbklXClBxcUhSMWdEY3lJVW14UE5UOXVqZTBBQXJWZXRZZUp4bmU1ajZNY0NnWUVBby84RGIwRlpiMXFMeVhyNDUwTmsKNHpzZlVwczFEcEFiVmNlSGdiZ3hUdnlqczBEbEI4Q3BPdUcydkJlSGhZajVEWVYzM29lRjBydkVObTVLdnRUSQpxZEFaVHNrR3IxZ3gwNlV5b2l0TkhUNUJSc0hlZytBRTg4Lzc3Ti9TVHFQL0JHMURrNU55amdZdlJZU2pZSzArCld6MEp0MGN2MnlVaTFMemh2N1hwV3hFQ2dZQWNVSmdlQ0ZTTXlRQzY0RVZuMjN4aFlKRVcyNEJRNzRvWXhKUkgKN0xya1JpcWNLZlNLT1A3UFlqOU56L0cwcmtIZlZnL2IxQUdpM2FPVzAzSHM3RUU1SDBXM3RpMHVabG4wdHFEZQozcDBFVnl3TThpTGNDM3hwV1Jwb1Izcm9tK2d3bFUxcytKZjhXY1VyY3ZhTGF6cUQrc3pRREUxSklzTzlqRXl5CjFHMmt0d0tCZ0FhdjY0UlhjWFM0SGRrK1JJMUlmSTdkR0h1b3FQRjB4ZVhzcnZCeTNWeUxpK1RHRzNraXNMMFkKY1F4S1RxODdUckpSWEE5YjFYU0ZHV0VvenVOQ2VCZjNEMGhTTWpkcXg2SEgrQ0MrTmdNOGtZUDBycWQ0N1plbAp3UzRyVWdpbm91TlFZYVlRaEpLQVd0Y21JN0JLNGhBb1JOLzRubE9hSXZDa2dPWGVpRW8zCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg=="
	testPublicKey  = "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlJQklqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FROEFNSUlCQ2dLQ0FRRUF6eUtPY2R1c0dJTy9VLzZnRTFsNAp0Z3Z4KzR5TUExQ1pQRkYwRnhKbzJaS1ZTY3I4SU41RVVDUDlUeEsxRTJLc2xnS01zQkhnYldJY3ZHMFBpTXZUClo0dTB3SWlPaTVSMDNlK3I5V1NqOG1xSCs3UjU1VndybUZVbFdMRWxuT1E4MnYveWNpV2hPZFJURWJ5cTZYQWcKaU5BckZyL3NFRTBacHFPdlVSeEFmeG5Qb1ZGd3M4NUplU0FYR1c2aG9HK3FoeEIvZ3diYkZOVitpbXViZUZ6dApyd1NJUGdmNjR2d2RoWnpDY1JZOVRUK1dyRm16Yk5uZmNwSzNvZEVnOCszdVJaWDdBb2R0U2E5OTZPTFFOcFNJClFnUDRiMnBXem5hc0NlRHU3ZlBWME5GNkJmSG1hcCs3RHZORTk1blcxUUh2MzZRRzNoVVRmZ0ZZQVJoc0NrOTAKcndJREFRQUIKLS0tLS1FTkQgUFVCTElDIEtFWS0tLS0tCg=="
	testToken      = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDk1NDA2NTcsImlhdCI6MTc0OTUzODg1NywiaXNzIjoiZ29waGtlZXBlciIsIm5iZiI6MTc0OTUzODg1NywidXNlcl9pZCI6MTF9.lOg4GN_4ybxxWlFY4e_ZR_bHeHQ1R9z1xmmXxCUvkBZ0vZRQPXsm0yuyUteKEYJP3B7vheuP8JFXtvADP4fukXxzVTxdubpDZTxN-zlntbRAdOZmHb29BtPgqhIDnpSwUtucPxgqMlhrzR6ZL1zJF4xl9BeYEAI3-8EFMfzVItd_pDD3_pvI_Vd8-Gg7Slrm9rjlHec60OugTVuWKJ-_MNVQsp-_cIvxOj8Mipz8oB2_UsWTJH-vK5DngaMvsxnHcq6gA7aDzwfSqjb2hur2oSqXX_xCblex2ntxxzNjbJpzJGQkeprIvVogksoGkqRcjhBbdcNKtqmG_hqSf8GPkg"
)

func TestNew(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		newConfig, err := config.NewConfig()
		require.NoError(t, err)
		require.NotNil(t, newConfig)

		jwt, err := New(Params{
			Config: newConfig,
		})
		require.NoError(t, err)
		require.NotNil(t, jwt)
	})
	t.Run("invalid private key", func(t *testing.T) {
		err := os.Setenv("JWT_PRIVATE_KEY", "invalid_private_key")
		require.NoError(t, err)

		newConfig, err := config.NewConfig()
		require.NoError(t, err)
		require.NotNil(t, newConfig)

		jwt, err := New(Params{
			Config: newConfig,
		})
		require.Error(t, err)
		require.Nil(t, jwt)
	})
}

func Test_jwtI_GenerateTokenPair(t *testing.T) {
	t.Run("success", func(t *testing.T) {

		err := os.Setenv("JWT_PRIVATE_KEY", testPrivateKey) // base64 encoded

		newConfig, err := config.NewConfig()
		require.NoError(t, err)
		require.NotNil(t, newConfig)

		jwt, err := New(Params{
			Config: newConfig,
		})
		require.NoError(t, err)
		require.NotNil(t, jwt)

		tokenPair, err := jwt.GenerateTokenPair(1)
		require.NoError(t, err)
		require.NotNil(t, tokenPair)
		require.NotEmpty(t, tokenPair.AccessToken)
		require.NotEmpty(t, tokenPair.RefreshToken)
	})
	t.Run("invalid private key", func(t *testing.T) {
		err := os.Setenv("JWT_PRIVATE_KEY", "invalid_private_key")
		require.NoError(t, err)

		newConfig, err := config.NewConfig()
		require.NoError(t, err)
		require.NotNil(t, newConfig)

		jwt, err := New(Params{
			Config: newConfig,
		})
		require.Error(t, err)
		require.Nil(t, jwt)
	})
}

func Test_jwtI_GenerateOnlyAccessToken(t *testing.T) {
	t.Run("invalid private key", func(t *testing.T) {
		err := os.Setenv("JWT_PRIVATE_KEY", "invalid_private_key")
		require.NoError(t, err)

		newConfig, err := config.NewConfig()
		require.NoError(t, err)
		require.NotNil(t, newConfig)

		jwt, err := New(Params{
			Config: newConfig,
		})
		require.Error(t, err)
		require.Nil(t, jwt)
	})
	t.Run("success", func(t *testing.T) {
		err := os.Setenv("JWT_PRIVATE_KEY", testPrivateKey)
		require.NoError(t, err)

		newConfig, err := config.NewConfig()
		require.NoError(t, err)
		require.NotNil(t, newConfig)

		jwt, err := New(Params{
			Config: newConfig,
		})
		require.NoError(t, err)
		require.NotNil(t, jwt)

		token, err := jwt.GenerateOnlyAccessToken(1)
		require.NoError(t, err)

		require.NotEmpty(t, token)
	})
}

func Test_generateSecureToken(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid length",
			args: args{
				length: 32,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateSecureToken(tt.args.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateSecureToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(got) == 0 {
				t.Errorf("generateSecureToken() got = %v, want non-empty result", got)
			}
		})
	}
}

func TestParse(t *testing.T) {
	bytes, err := base64.StdEncoding.DecodeString(testPublicKey)
	require.NoError(t, err)
	require.NotNil(t, bytes)

	type args struct {
		token     string
		publicKey []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "invalid private key",
			args: args{
				token:     "invalid_token",
				publicKey: []byte("invalid_public_key"),
			},
			wantErr: jwt.ErrKeyMustBePEMEncoded,
		},
		{
			name: "invalid token",
			args: args{
				token:     "invalid_token",
				publicKey: bytes,
			},
			wantErr: jwt.ErrTokenMalformed,
		},
		{
			name: "valid token",
			args: args{
				token:     testToken,
				publicKey: bytes,
			},
			wantErr: ErrTokenExpired,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Parse(tt.args.token, tt.args.publicKey)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	t.Run("valide token", func(t *testing.T) {
		err := os.Setenv("JWT_PRIVATE_KEY", testPrivateKey)
		require.NoError(t, err)

		newConfig, err := config.NewConfig()
		require.NoError(t, err)
		require.NotNil(t, newConfig)

		jwt, err := New(Params{
			Config: newConfig,
		})
		require.NoError(t, err)
		require.NotNil(t, jwt)

		token, err := jwt.GenerateOnlyAccessToken(1)
		require.NoError(t, err)
		require.NotEmpty(t, token)

		parsedToken, err := Parse(token, bytes)
		require.NoError(t, err)
		require.NotNil(t, parsedToken)

		claimExpiration, err := parsedToken.GetExpirationTime()
		require.NoError(t, err)
		require.NotNil(t, claimExpiration)

		userID, err := GetUserIDFromClaims(parsedToken)
		require.NoError(t, err)
		require.Equal(t, 1, userID)
	})
}

func TestGetUserIDFromClaims(t *testing.T) {
	type args struct {
		claims jwt.Claims
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "valid claims",
			args: args{
				claims: jwt.MapClaims{
					"iss":     "gophkeeper",
					"exp":     1749554065,
					"iat":     1749553857,
					"nbf":     1749553857,
					"user_id": 11.,
				},
			},
			want:    11,
			wantErr: false,
		},
		{
			name: "missing user_id",
			args: args{
				claims: jwt.MapClaims{
					"iss": "gophkeeper",
					"exp": 1749554065,
					"iat": 1749553857,
					"nbf": 1749553857,
				},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "invalid claims type",
			args: args{
				claims: nil,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUserIDFromClaims(tt.args.claims)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserIDFromClaims() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetUserIDFromClaims() got = %v, want %v", got, tt.want)
			}
		})
	}
}
