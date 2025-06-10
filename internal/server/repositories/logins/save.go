package logins

import (
	"context"
	"fmt"
)

func (r *repo) Save(ctx context.Context, login *LoginData) (int, error) {
	var id int
	err := r.dbConn.QueryRow(ctx, `
INSERT INTO logins (user_id, login, password, title, note, created_at)
VALUES ($1, $2, $3, $4, $5, $6)
returning id`,
		login.UserID, login.Login, login.Password,
		login.Title, login.Note, login.CreatedAt,
	).
		Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to save login: %w", err)
	}

	return id, nil
}
