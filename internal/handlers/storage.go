package handlers

import (
	"fmt"
	"net/http"

	"github.com/prajwalbharadwajbm/gupload/internal/db/repository"
	"github.com/prajwalbharadwajbm/gupload/internal/interceptor"
	"github.com/prajwalbharadwajbm/gupload/internal/logger"
)

func StorageRemaining(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	storageQuota, storageRemaining, err := repository.GetStorageRemaining(ctx)
	if err != nil {
		logger.Log.Error("unable to get storage remaining", err)
		interceptor.SendErrorResponse(w, "", http.StatusInternalServerError)
		return
	}
	storageQuotaInMB, storageRemainingInMB := convertToHumanReadable(storageQuota, storageRemaining)
	response := map[string]interface{}{
		"storage_quota":     storageQuotaInMB,
		"storage_remaining": storageRemainingInMB,
	}
	interceptor.SendSuccessResponse(w, response, http.StatusOK)
}

func convertToHumanReadable(storageQuota, storageRemaining int) (string, string) {
	quotaMB := float64(storageQuota) / 1024 / 1024
	remainingMB := float64(storageRemaining) / 1024 / 1024

	return fmt.Sprintf("%.3f MB", quotaMB), fmt.Sprintf("%.3f MB", remainingMB)
}
