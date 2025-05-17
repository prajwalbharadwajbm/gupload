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

func GetUsernameByUserID(ctx context.Context) (string, error) {
	db := db.GetClient()

	var username string
	query := `SELECT username FROM users WHERE id = $1::uuid`
	err := db.QueryRowContext(ctx, query, ctx.Value("userId").(string)).Scan(&username)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("user not found")
		}
		return "", err
	}
	return username, nil
}
