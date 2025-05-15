package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/prajwalbharadwajbm/gupload/internal/handlers"
)

func Routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/healthCheck", handlers.HealthCheck)

	router.HandlerFunc(http.MethodPost, "/register", handlers.Register)

	return router
}
