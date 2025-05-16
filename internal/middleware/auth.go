package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/prajwalbharadwajbm/gupload/internal/interceptor"
	"github.com/prajwalbharadwajbm/gupload/internal/logger"
	"github.com/prajwalbharadwajbm/gupload/internal/service/auth"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			logger.Log.Info("missing authorization header")
			interceptor.SendErrorResponse(w, "GUPLD107", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			logger.Log.Info("invalid authorization header format")
			interceptor.SendErrorResponse(w, "GUPLD107", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]

		claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			logger.Log.Info("invalid token: %v", err)
			interceptor.SendErrorResponse(w, "GUPLD107", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userId", claims.UserID)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
