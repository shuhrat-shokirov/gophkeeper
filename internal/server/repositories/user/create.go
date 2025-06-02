package user

import (
	"context"
	"fmt"
)

func (r *repo) CreateUser(ctx context.Context, user *User) (int, error) {
	var id int

	err := r.dbConn.QueryRow(ctx, `INSERT INTO users (email, password)
VALUES ($1, $2)
RETURNING id`, user.Email, user.Password).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("could not create user: %w", err)
	}

	return id, nil
}
