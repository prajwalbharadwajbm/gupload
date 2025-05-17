package handlers

import (
	"net/http"
	"strconv"

	"github.com/prajwalbharadwajbm/gupload/internal/db/repository"
	"github.com/prajwalbharadwajbm/gupload/internal/interceptor"
	"github.com/prajwalbharadwajbm/gupload/internal/logger"
)

var (
	page  = 1
	limit = 10
)

func FetchFiles(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	if pageStr != "" {
		pageInt, err := strconv.Atoi(pageStr)
		if err == nil && pageInt > 0 {
			page = pageInt
		} else {

		}
	}

	if limitStr != "" {
		limitInt, err := strconv.Atoi(limitStr)
		if err == nil && limitInt > 0 {
			limit = limitInt
		} else {

		}
	}

	files, err := repository.GetFilesByUserId(ctx)
	if err != nil {
		logger.Log.Error("unable to get files", err)
		interceptor.SendErrorResponse(w, "", http.StatusInternalServerError)
		return
	}

	totalFiles := len(files)
	startIndex := (page - 1) * limit
	endIndex := startIndex + limit

	if startIndex >= totalFiles {
		paginatedResponse := map[string]interface{}{
			"files": []map[string]any{},
			"pagination": map[string]int{
				"current_page": page,
				"page_size":    limit,
				"total_items":  totalFiles,
				"total_pages":  (totalFiles + limit - 1) / limit,
			},
		}
		interceptor.SendSuccessResponse(w, paginatedResponse, http.StatusOK)
		return
	}

	// 25 files
	// 3rd page with limit 10
	// startIndex = (3-1)*10 = 20
	// endIndex = 20+10 = 30 (30 is out of bound, and will cause panic)
	if endIndex > totalFiles {
		endIndex = totalFiles
	}

	paginatedFiles := files[startIndex:endIndex]

	paginatedResponse := map[string]interface{}{
		"files": paginatedFiles,
		"pagination": map[string]int{
			"current_page": page,
			"page_size":    limit,
			"total_items":  totalFiles,
			"total_pages":  (totalFiles + limit - 1) / limit,
		},
	}

	interceptor.SendSuccessResponse(w, paginatedResponse, http.StatusOK)
}
