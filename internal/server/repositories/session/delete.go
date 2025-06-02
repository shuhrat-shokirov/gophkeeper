package session

import (
	"context"
	"fmt"
)

func (r *repo) Delete(ctx context.Context, token string) error {
	_, err := r.dbConn.Exec(ctx, "DELETE FROM sessions WHERE refresh_token = $1", token)
	if err != nil {
		return fmt.Errorf("delete session: %w", err)
	}

	return nil
}
