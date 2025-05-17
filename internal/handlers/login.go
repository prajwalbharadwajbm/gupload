package handlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/prajwalbharadwajbm/gupload/internal/db/repository"
	"github.com/prajwalbharadwajbm/gupload/internal/dtos"
	"github.com/prajwalbharadwajbm/gupload/internal/interceptor"
	"github.com/prajwalbharadwajbm/gupload/internal/logger"
	"github.com/prajwalbharadwajbm/gupload/internal/service/auth"
	"github.com/prajwalbharadwajbm/gupload/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userData, err := utils.FetchDataFromRequestBody[dtos.User](r)
	if err != nil {
		logger.Log.Error("unable to fetch request body", err)
		interceptor.SendErrorResponse(w, "GULPD001", http.StatusBadRequest)
		return
	}

	authenticated, userId, err := authenticateUser(ctx, userData)
	if err != nil {
		logger.Log.Error("authentication error", err)
		interceptor.SendErrorResponse(w, "GUPLD102", http.StatusInternalServerError)
		return
	}
	if !authenticated {
		logger.Log.Info("invalid username or password")
		if errors.Is(err, errors.New("user not found")) {
			interceptor.SendErrorResponse(w, "GUPLD102", http.StatusNotFound)
		} else {
			interceptor.SendErrorResponse(w, "GUPLD102", http.StatusUnauthorized)
		}
		return
	}
	token, err := auth.GenerateToken(userId)
	if err != nil {
		logger.Log.Error("failed to generate token", err)
		interceptor.SendErrorResponse(w, "GUPLD106", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"token":   token,
		"user_id": userId,
	}
	interceptor.SendSuccessResponse(w, response, http.StatusOK)
}

func authenticateUser(ctx context.Context, userData dtos.User) (bool, string, error) {
	userId, hashedPassword, err := repository.GetUserByUsername(ctx, userData.Username)
	if err != nil {
		return false, "", fmt.Errorf("unable to fetch user by username: %w", err)
	}
	if userId == "" {
		return false, "", nil
	}
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(userData.Password))
	if err != nil {
		return false, "", nil
	}
	return true, userId, nil
}
