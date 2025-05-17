package handlers

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/prajwalbharadwajbm/gupload/internal/db/repository"
	"github.com/prajwalbharadwajbm/gupload/internal/interceptor"
	"github.com/prajwalbharadwajbm/gupload/internal/logger"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		logger.Log.Error("unable to parse form data", err)
		interceptor.SendErrorResponse(w, "GUPLD001", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		logger.Log.Error("unable to get file", err)
		interceptor.SendErrorResponse(w, "GUPLD201", http.StatusBadRequest)
		return
	}
	defer file.Close()

	filepath, ok, err := uploadUserFiles(ctx, file, header)
	if err != nil {
		logger.Log.Error("unable to upload user files", err)
		if err.Error() != "" && err.Error() == "file with name already exists" {
			interceptor.SendErrorResponse(w, "File with this name already exists. Please rename your file and try again.", http.StatusConflict)
			return
		}
		interceptor.SendErrorResponse(w, "GUPLD202", http.StatusBadRequest)
		return
	}
	if !ok {
		interceptor.SendErrorResponse(w, "GUPLD203", http.StatusInsufficientStorage)
		return
	}
	response := map[string]interface{}{
		"filepath": filepath,
	}
	logger.Log.Infof("Successfully uploaded file for user_id: %s", ctx.Value("userId"))
	interceptor.SendSuccessResponse(w, response, http.StatusOK)
}

func uploadUserFiles(ctx context.Context, file multipart.File, header *multipart.FileHeader) (string, bool, error) {
	username, err := repository.GetUsernameByUserID(ctx)
	if err != nil {
		logger.Log.Error("unable to get username by user ID", err)
		return "", false, fmt.Errorf("failed to get username: %w", err)
	}

	// Check if a file with the same name already exists
	exists, err := repository.CheckFileExists(ctx, header.Filename)
	if err != nil {
		return "", false, fmt.Errorf("unable to check if file exists: %w", err)
	}
	if exists {
		return "", false, errors.New("file with name already exists")
	}

	ok, err := checkStorageAvailability(ctx, header)
	if err != nil {
		return "", false, fmt.Errorf("unable to check storage availability: %w", err)
	}
	if !ok {
		return "", ok, nil
	}

	dirPath := fmt.Sprintf("./gupload/%s/", username)
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return "", false, fmt.Errorf("unable to create directory: %w", err)
	}

	filepath := filepath.Join(dirPath, header.Filename)
	dst, err := os.Create(filepath)
	if err != nil {
		return "", false, fmt.Errorf("unable to create file: %w", err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return "", false, fmt.Errorf("unable to copy file contents: %w", err)
	}

	err = repository.CreateFileLogs(ctx, filepath, header.Filename, header.Size, header.Header.Get("Content-Type"))
	if err != nil {
		return "", false, fmt.Errorf("unable to create file logs: %w", err)
	}

	err = repository.UpdateStorageQuota(ctx, header.Size)
	if err != nil {
		return "", false, fmt.Errorf("unable to update storage quota: %w", err)
	}

	return filepath, true, nil
}

func checkStorageAvailability(ctx context.Context, header *multipart.FileHeader) (bool, error) {
	userStorage, err := repository.GetStorageInfoByUserID(ctx)
	if err != nil {
		return false, fmt.Errorf("unable to get user storage info: %w", err)
	}

	fileSize := header.Size
	if int64(userStorage.UsedBytes)+fileSize > int64(userStorage.MaxBytes) {
		logger.Log.Info("User storage quota exceeded")
		return false, nil
	}
	return true, nil
}
