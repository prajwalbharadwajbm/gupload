package handlers

import (
	"net/http"

	"github.com/prajwalbharadwajbm/gupload/internal/db/repository"
	"github.com/prajwalbharadwajbm/gupload/internal/interceptor"
	"github.com/prajwalbharadwajbm/gupload/internal/logger"
)

func FetchFiles(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	files, err := repository.GetFilesByUserId(ctx)
	if err != nil {
		logger.Log.Error("unable to get files", err)
		interceptor.SendErrorResponse(w, "", http.StatusInternalServerError)
		return
	}

	interceptor.SendSuccessResponse(w, files, http.StatusOK)
}
