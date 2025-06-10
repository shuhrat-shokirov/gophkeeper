package auth

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"gophkeeper/internal/server/repositories/session"
)

//nolint:lll,gocritic
const testPrivateKey = "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFb3dJQkFBS0NBUUVBenlLT2NkdXNHSU8vVS82Z0UxbDR0Z3Z4KzR5TUExQ1pQRkYwRnhKbzJaS1ZTY3I4CklONUVVQ1A5VHhLMUUyS3NsZ0tNc0JIZ2JXSWN2RzBQaU12VFo0dTB3SWlPaTVSMDNlK3I5V1NqOG1xSCs3UjUKNVZ3cm1GVWxXTEVsbk9RODJ2L3ljaVdoT2RSVEVieXE2WEFnaU5BckZyL3NFRTBacHFPdlVSeEFmeG5Qb1ZGdwpzODVKZVNBWEdXNmhvRytxaHhCL2d3YmJGTlYraW11YmVGenRyd1NJUGdmNjR2d2RoWnpDY1JZOVRUK1dyRm16CmJObmZjcEszb2RFZzgrM3VSWlg3QW9kdFNhOTk2T0xRTnBTSVFnUDRiMnBXem5hc0NlRHU3ZlBWME5GNkJmSG0KYXArN0R2TkU5NW5XMVFIdjM2UUczaFVUZmdGWUFSaHNDazkwcndJREFRQUJBb0lCQUZ1Wkx2L1h3b2VHdjNYUAovSThCK25rYTJEUkM1Mm5SMnluSzVYa01lWlI1bDQ0dDl3Zzc4bDYwUTZFVHAwSysySTV2NnpJemZaa3hrWDZjCkJnb2JCTTVlQUIxQ1pqTUFnQnZqRUpxd21qV3ArWitNSkhtU3BHNjFmSkgzcUtmMFlKc0NJMzlwOTUzQXNNbC8Kc3Q4UGFEdklQcjNOT29ITTdxSjc4UnYvejkvRVNSdVBuUFZPaEtQSXh6OFZ4Z2p5eGRlemM5Y0NvbnZLbzV0bgpXUGFjOStlYkwvSTFFWDdHZ1IwVnZzbDdtc0RoVVorcnkraUdnM3pBVzhIOGJ5NUFGR3lkdUNjeGtrWmVUVFRFClR1SllVSGdCTWN5NUtCNStxdzNvRk1QVnFNeTE3RktCRncwMW8vQlE3WlFCdERNUFhoWVJyY2RDRk9BOThia1QKbzlnQzFka0NnWUVBNTVxZ0NUUzFibk9xY3lmOFpMdkM2ZFdyVWQ1Q1VZVEZiNWZWUGlYRmxpTTExR2FRbWxGUgpyTmxjTU83ZEVHNVNvUmkxNzJHQVJZK24xZi84QXlGN2g0WHlqanFkbnNybmhFckFJRXNiN1h1cFRzM3RqL1pXCjhJYWFFZ0VxUExITHVrdHhyWjh2aHdtamxlSXhrVmVGR0N1S3JSMFc0bWRlV3ArNmJjYnJ2TmtDZ1lFQTVQUWIKSHI4c2FhdTIrSHVTTm90czl5ZUhNdmRKczkvZzdmeHVzUGs4STRrVHdGZHpZeUhNSmcvdzEwTTY3VFNiYzJFcAp1SXlxSlpCVUNDNkVleUF3cjJERjBiRkJEWGpqZ3hFQzYrMU1QdXlXV2JWam1EaWRreFA4WVRkb0k5ZjVrbklXClBxcUhSMWdEY3lJVW14UE5UOXVqZTBBQXJWZXRZZUp4bmU1ajZNY0NnWUVBby84RGIwRlpiMXFMeVhyNDUwTmsKNHpzZlVwczFEcEFiVmNlSGdiZ3hUdnlqczBEbEI4Q3BPdUcydkJlSGhZajVEWVYzM29lRjBydkVObTVLdnRUSQpxZEFaVHNrR3IxZ3gwNlV5b2l0TkhUNUJSc0hlZytBRTg4Lzc3Ti9TVHFQL0JHMURrNU55amdZdlJZU2pZSzArCld6MEp0MGN2MnlVaTFMemh2N1hwV3hFQ2dZQWNVSmdlQ0ZTTXlRQzY0RVZuMjN4aFlKRVcyNEJRNzRvWXhKUkgKN0xya1JpcWNLZlNLT1A3UFlqOU56L0cwcmtIZlZnL2IxQUdpM2FPVzAzSHM3RUU1SDBXM3RpMHVabG4wdHFEZQozcDBFVnl3TThpTGNDM3hwV1Jwb1Izcm9tK2d3bFUxcytKZjhXY1VyY3ZhTGF6cUQrc3pRREUxSklzTzlqRXl5CjFHMmt0d0tCZ0FhdjY0UlhjWFM0SGRrK1JJMUlmSTdkR0h1b3FQRjB4ZVhzcnZCeTNWeUxpK1RHRzNraXNMMFkKY1F4S1RxODdUckpSWEE5YjFYU0ZHV0VvenVOQ2VCZjNEMGhTTWpkcXg2SEgrQ0MrTmdNOGtZUDBycWQ0N1plbAp3UzRyVWdpbm91TlFZYVlRaEpLQVd0Y21JN0JLNGhBb1JOLzRubE9hSXZDa2dPWGVpRW8zCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg=="

func Test_service_RefreshToken(t *testing.T) {
	t.Run("err on session get", func(t *testing.T) {
		service, cache, emailGateway, userRepo, sessionRepo := testInit(t)
		defer cache.AssertExpectations(t)
		defer emailGateway.AssertExpectations(t)
		defer userRepo.AssertExpectations(t)
		defer sessionRepo.AssertExpectations(t)

		sessionRepo.On("Get", t.Context(), "token").
			Return(nil, assert.AnError)

		accessToken, err := service.RefreshToken(t.Context(), "token")
		assert.Error(t, err)
		assert.Empty(t, accessToken)
	})

	t.Run("err on jwt generate", func(t *testing.T) {
		err := os.Setenv("JWT_PRIVATE_KEY", "")
		require.NoError(t, err)

		service, cache, emailGateway, userRepo, sessionRepo := testInit(t)

		defer cache.AssertExpectations(t)
		defer emailGateway.AssertExpectations(t)
		defer userRepo.AssertExpectations(t)
		defer sessionRepo.AssertExpectations(t)

		sessionRepo.On("Get", t.Context(), "token").
			Return(&session.Session{UserID: 1}, nil)

		accessToken, err := service.RefreshToken(t.Context(), "token")
		assert.Error(t, err)
		assert.Empty(t, accessToken)
	})

	t.Run("success", func(t *testing.T) {
		err := os.Setenv("JWT_PRIVATE_KEY", testPrivateKey)
		require.NoError(t, err)

		service, cache, emailGateway, userRepo, sessionRepo := testInit(t)
		defer cache.AssertExpectations(t)
		defer emailGateway.AssertExpectations(t)
		defer userRepo.AssertExpectations(t)
		defer sessionRepo.AssertExpectations(t)

		sessionRepo.On("Get", t.Context(), "token").
			Return(&session.Session{UserID: 1}, nil)

		accessToken, err := service.RefreshToken(t.Context(), "token")
		assert.NoError(t, err)
		assert.NotEmpty(t, accessToken)
	})
}
