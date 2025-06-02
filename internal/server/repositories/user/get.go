package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"

	"gophkeeper/internal/server/errorx"
)

func (r *repo) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	user := &User{}

	err := r.dbConn.QueryRow(ctx, `SELECT id,
       email,
       password,
       created_at,
       updated_at
FROM users
WHERE email = $1`, email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errorx.ErrNotFound
		}

		return nil, fmt.Errorf("could not get user: %w", err)
	}

	return user, nil
}

func (r *repo) GetUserByID(ctx context.Context, id int) (*User, error) {
	user := &User{}

	err := r.dbConn.QueryRow(ctx, `SELECT id,
	   email,
	   password,
	   created_at,
	   updated_at
FROM users
WHERE id = $1`, id).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errorx.ErrNotFound
		}

		return nil, fmt.Errorf("could not get user by id: %w", err)
	}

	return user, nil
}
