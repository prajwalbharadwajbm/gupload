package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/prajwalbharadwajbm/gupload/internal/db"
)

func AddUser(ctx context.Context, username string, password []byte) (string, error) {
	db := db.GetClient()
	userId := uuid.New().String()
	query := `INSERT INTO users (id, username, password) VALUES ($1, $2, $3)`
	err := db.QueryRowContext(ctx, query, userId, username, password)
	if err != nil {
		return "", err.Err()
	}
	return userId, nil
}
