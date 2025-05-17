package handlers

import (
	"net/http"

	"github.com/prajwalbharadwajbm/gupload/internal/db/repository"
	"github.com/prajwalbharadwajbm/gupload/internal/interceptor"
	"github.com/prajwalbharadwajbm/gupload/internal/logger"
	"github.com/prajwalbharadwajbm/gupload/internal/utils"
)

func StorageRemaining(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	storageQuota, storageRemaining, err := repository.GetStorageRemaining(ctx)
	if err != nil {
		logger.Log.Error("unable to get storage remaining", err)
		interceptor.SendErrorResponse(w, "", http.StatusInternalServerError)
		return
	}
	humanReadableStorageQuota := utils.FormatBytes(int64(storageQuota))
	humanReadableStorageRemaining := utils.FormatBytes(int64(storageRemaining))

	response := map[string]interface{}{
		"storage_quota":     humanReadableStorageQuota,
		"storage_remaining": humanReadableStorageRemaining,
	}
	interceptor.SendSuccessResponse(w, response, http.StatusOK)
}
