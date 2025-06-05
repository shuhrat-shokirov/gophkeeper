package card

import (
	"context"
	"fmt"
)

func (r *repo) Save(ctx context.Context, data *Data) (int64, error) {
	var id int64
	err := r.conn.QueryRow(ctx, `
INSERT INTO cards (user_id, title, pan, cvv, expiry, note, created_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id`,
		data.UserID, data.Title, data.Pan,
		data.Cvv, data.Expiry, data.Note, data.CreatedAt).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("save card: %w", err)
	}

	return id, nil
}
