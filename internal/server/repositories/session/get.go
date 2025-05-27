package session

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"gophkeeper/internal/server/exceptions"
)

func (r *repo) Get(ctx context.Context, refreshToken string) (*Session, error) {
	var session Session
	err := r.dbConn.QueryRow(ctx, `
SELECT id, user_id, refresh_token, created_at, updated_at, expired_at
FROM sessions
WHERE refresh_token = $1 and expired_at > NOW()`,
		refreshToken).Scan(
		&session.ID,
		&session.UserID,
		&session.RefreshToken,
		&session.CreatedAt,
		&session.UpdatedAt,
		&session.ExpiredAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, exceptions.ErrNotFound
		}
		return nil, fmt.Errorf("get session: %w", err)
	}

	return &session, nil

}
