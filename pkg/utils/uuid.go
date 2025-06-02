package utils

import (
	"github.com/google/uuid"
	"github.com/lithammer/shortuuid/v4"
)

func GenerateShortUUID() string {
	return shortuuid.New()
}

func GenerateUUID() string {
	newV7, err := uuid.NewV7()
	if err != nil {
		return uuid.New().String()
	}

	return newV7.String()
}
