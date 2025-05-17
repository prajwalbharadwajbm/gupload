package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/prajwalbharadwajbm/gupload/internal/handlers"
	"github.com/prajwalbharadwajbm/gupload/internal/middleware"
)

func Routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/healthCheck", handlers.HealthCheck)

	router.HandlerFunc(http.MethodPost, "/register", handlers.Register)
	router.HandlerFunc(http.MethodPost, "/login", handlers.Login)

	// Protected Endpoints
	router.HandlerFunc(http.MethodPost, "/upload", middleware.AuthMiddleware(handlers.Upload))

	return router
}
