package logins

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"

	"gophkeeper/internal/server/errorx"
)

func (r *repo) List(ctx context.Context, userID int64, pg Pagination) ([]List, error) {
	rows, err := r.dbConn.Query(ctx, `
SELECT id, title, created_at, updated_at
from logins 
where user_id = $1
limit $2
offset $3
`, userID, pg.GetLimit(), pg.GetOffset())
	if err != nil {
		return nil, fmt.Errorf("error while getting logins list: %w", err)
	}
	defer rows.Close()

	var logins = make([]List, 0, pg.GetLimit())
	for rows.Next() {
		var login List
		if err := rows.Scan(&login.ID, &login.Title, &login.CreatedAt, &login.UpdatedAt); err != nil {
			return nil, fmt.Errorf("error while scanning login: %w", err)
		}
		logins = append(logins, login)
	}

	if len(logins) == 0 {
		return nil, errorx.ErrNotFound
	}

	return logins, nil
}

func (r *repo) GetByID(ctx context.Context, id int64) (*Info, error) {
	var info = &Info{}
	err := r.dbConn.QueryRow(ctx, `
select id, user_id, login, password, title, note, created_at, updated_at
from logins
where id = $1`, id).
		Scan(&info.ID, &info.UserID, &info.Login, &info.Password,
			&info.Title, &info.Note, &info.CreatedAt, &info.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errorx.ErrNotFound
		}
		return nil, fmt.Errorf("error while getting login: %w", err)
	}

	return info, nil
}
