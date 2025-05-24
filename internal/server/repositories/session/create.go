package session

import (
	"context"
	"fmt"
)

func (r *repo) Create(ctx context.Context, session *Session) error {
	_, err := r.dbConn.Exec(ctx, `INSERT INTO sessions (user_id, refresh_token)
	VALUES ($1, $2)`,
		session.UserID,
		session.RefreshToken,
	)
	if err != nil {
		return fmt.Errorf("create session: %w", err)
	}

	return nil

}
