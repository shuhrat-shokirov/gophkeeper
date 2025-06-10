package texts

import (
	"context"
	"fmt"
)

func (r *repo) Save(ctx context.Context, data *TextData) (int64, error) {
	var id int64
	err := r.conn.QueryRow(ctx, `
INSERT INTO texts (user_id, title, content, note, created_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING id`,
		data.UserID, data.Title, data.Content,
		data.Note, data.CreatedAt).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to save text: %w", err)
	}

	return id, nil
}
