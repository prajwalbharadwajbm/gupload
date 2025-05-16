package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/prajwalbharadwajbm/gupload/internal/db"
)

func AddUser(ctx context.Context, username string, password []byte) (string, error) {
	db := db.GetClient()

	userId := uuid.New().String()

	query := `INSERT INTO users (id, username, password) VALUES ($1, $2, $3)`
	_, err := db.ExecContext(ctx, query, userId, username, password)
	if err != nil {
		return "", err
	}
	return userId, nil
}

func GetUserByUsername(ctx context.Context, userName string) (string, []byte, error) {
	db := db.GetClient()

	var userId string
	var hashedPassword []byte

	query := `SELECT id, password FROM users WHERE username = $1`
	err := db.QueryRowContext(ctx, query, userName).Scan(&userId, &hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil, errors.New("user not found")
		}
		return "", nil, err
	}
	return userId, hashedPassword, nil
}
