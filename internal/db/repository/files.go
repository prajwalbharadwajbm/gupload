package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/prajwalbharadwajbm/gupload/internal/db"
)

func CreateFileLogs(ctx context.Context, filepath string, filename string, size int64, contentType string) error {
	db := db.GetClient()
	uuid := uuid.New()

	query := `INSERT INTO files (id, user_id, file_path, filename, size_bytes, content_type) VALUES ($1::uuid, $2::uuid, $3, $4, $5, $6)`
	_, err := db.ExecContext(ctx, query, uuid, ctx.Value("userId"), filepath, filename, size, contentType)
	if err != nil {
		return err
	}
	return nil
}
