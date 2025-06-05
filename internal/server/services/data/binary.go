package data

import (
	"context"
	"fmt"
	"time"

	"gophkeeper/internal/server/repositories/binary"
	"gophkeeper/pkg/aes"
)

func (s *service) SaveBinary(ctx context.Context, data *BinaryData) error {
	_, err := s.binaryRepo.Save(ctx, &binary.Data{
		UserID:    data.UserID,
		Title:     data.Title,
		Content:   aes.MustEncrypt(string(data.Content)),
		Note:      aes.MustEncrypt(data.Note),
		CreatedAt: time.Unix(0, data.CreatedAt),
	})
	if err != nil {
		return fmt.Errorf("save binary: %w", err)
	}

	return nil

}
