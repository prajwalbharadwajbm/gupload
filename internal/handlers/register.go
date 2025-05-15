package handlers

import (
	"context"
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

	responseID, err := registerUser(ctx, userData)
	if err != nil {
		logger.Log.Error("unable to register user", err)
		interceptor.SendErrorResponse(w, "GUPLD101", http.StatusBadRequest)
		return
	}
	interceptor.SendSuccessResponse(w, responseID, http.StatusOK)
}

func registerUser(ctx context.Context, userData dtos.User) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Log.Error("unable to hash password", err)
		return "", err
	}
	userId, err := repository.AddUser(ctx, userData.Username, hashedPassword)
	if err != nil {
		logger.Log.Error("unable to add user", err)
		return "", err
	}
	return userId, nil
}

func validateRequestBody(userData dtos.User) (bool, error) {
	if valid, err := validator.IsValidUsername(userData.Username); !valid || err != nil {
		logger.Log.Error("username: %s is not valid", err)
		return false, err
	}

	if valid, err := validator.IsValidPassword(userData.Username, userData.Password); !valid || err != nil {
		logger.Log.Error("password doesn't follow org policy/rule: %s", err)
		return false, err
	}
	return true, nil
}
