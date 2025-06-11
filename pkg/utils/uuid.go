package utils

import (
	"github.com/lithammer/shortuuid/v4"
)

func GenerateShortUUID() string {
	return shortuuid.New()
}
