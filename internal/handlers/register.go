package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/prajwalbharadwajbm/gupload/internal/db/repository"
	"github.com/prajwalbharadwajbm/gupload/internal/dtos"
	"github.com/prajwalbharadwajbm/gupload/internal/interceptor"
	"github.com/prajwalbharadwajbm/gupload/internal/logger"
	"github.com/prajwalbharadwajbm/gupload/internal/utils"
	"github.com/prajwalbharadwajbm/gupload/internal/validator"
	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userData, err := utils.FetchDataFromRequestBody[dtos.User](r)
	if err != nil {
		logger.Log.Error("unable to fetch request body", err)
		interceptor.SendErrorResponse(w, "GUPLD001", http.StatusBadRequest)
		return
	}

	valid, err := validateRequestBody(userData)
	if !valid || err != nil {
		logger.Log.Info("request body is not valid %v", err)
		interceptor.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, err := registerUser(ctx, userData)
	if err != nil {
		logger.Log.Error("unable to register user", err)
		interceptor.SendErrorResponse(w, "GUPLD101", http.StatusBadRequest)
		return
	}
	response := map[string]interface{}{
		"userID": userID,
	}
	interceptor.SendSuccessResponse(w, response, http.StatusOK)
}

func registerUser(ctx context.Context, userData dtos.User) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("unable to hash password: %w", err)
	}
	userId, err := repository.AddUser(ctx, userData.Username, hashedPassword)
	if err != nil {
		return "", fmt.Errorf("unable to add user: %w", err)
	}
	err = repository.CreateStorageQuota(ctx, userId)
	if err != nil {
		return "", fmt.Errorf("unable to create storage quota: %w", err)
	}
	return userId, nil
}

func validateRequestBody(userData dtos.User) (bool, error) {
	if valid, err := validator.IsValidUsername(userData.Username); !valid || err != nil {
		return false, fmt.Errorf("invalid username: %w", err)
	}

	if valid, err := validator.IsValidPassword(userData.Username, userData.Password); !valid || err != nil {
		return false, fmt.Errorf("invalid password: %w", err)
	}
	return true, nil
}
