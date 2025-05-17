package repository

import (
	"context"
	"database/sql"

	"github.com/prajwalbharadwajbm/gupload/internal/db"
	"github.com/prajwalbharadwajbm/gupload/internal/db/models"
)

func GetStorageInfoByUserID(ctx context.Context) (models.StorageQuota, error) {
	db := db.GetClient()

	storageInfo := models.StorageQuota{}

	query := `SELECT max_bytes, used_bytes FROM storage_quota WHERE user_id = $1::uuid`
	err := db.QueryRowContext(ctx, query, ctx.Value("userId")).Scan(&storageInfo.MaxBytes, &storageInfo.UsedBytes)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.StorageQuota{}, err
		}
		return models.StorageQuota{}, err
	}
	return storageInfo, nil

}

func CreateStorageQuota(ctx context.Context, userId string) error {
	db := db.GetClient()

	query := `INSERT INTO storage_quota (user_id, max_bytes, used_bytes) VALUES ($1::uuid, $2, $3)`
	_, err := db.ExecContext(ctx, query, userId, 50<<20, 0)
	if err != nil {
		return err
	}
	return nil
}

func UpdateStorageQuota(ctx context.Context, size int64) error {
	db := db.GetClient()

	query := `UPDATE storage_quota SET used_bytes = used_bytes + $1 WHERE user_id = $2::uuid`
	_, err := db.ExecContext(ctx, query, size, ctx.Value("userId"))
	if err != nil {
		return err
	}
	return nil
}
