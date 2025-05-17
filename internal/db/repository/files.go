package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/prajwalbharadwajbm/gupload/internal/db"
	"github.com/prajwalbharadwajbm/gupload/internal/db/models"
	"github.com/prajwalbharadwajbm/gupload/internal/utils"
)

func CreateFileLogs(ctx context.Context, filepath string, filename string, size int64, contentType string) error {
	db := db.GetClient()
	uuid := uuid.New()

	dbCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `INSERT INTO files (id, user_id, file_path, filename, size_bytes, content_type) VALUES ($1::uuid, $2::uuid, $3, $4, $5, $6)`
	_, err := db.ExecContext(dbCtx, query, uuid, ctx.Value("userId"), filepath, filename, size, contentType)
	if err != nil {
		return err
	}
	return nil
}

func GetFilesByUserId(ctx context.Context) ([]map[string]any, error) {
	db := db.GetClient()

	dbCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `SELECT filename, file_path, size_bytes, content_type FROM files WHERE user_id = $1::uuid`
	rows, err := db.QueryContext(dbCtx, query, ctx.Value("userId"))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []map[string]any
	for rows.Next() {
		files := models.Files{}
		err := rows.Scan(&files.Filename, &files.FilePath, &files.SizeBytes, &files.ContentType)
		if err != nil {
			return nil, err
		}
		filteredResponse := map[string]any{
			"filename": files.Filename,
			"filepath": files.FilePath,
			"Size":     utils.FormatBytes(int64(files.SizeBytes)),
			"content":  files.ContentType,
		}
		result = append(result, filteredResponse)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func CheckFileExists(ctx context.Context, filename string) (bool, error) {
	db := db.GetClient()

	dbCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var count int
	checkQuery := `SELECT COUNT(*) FROM files WHERE user_id = $1::uuid AND filename = $2 AND is_deleted = FALSE`
	err := db.QueryRowContext(dbCtx, checkQuery, ctx.Value("userId"), filename).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
