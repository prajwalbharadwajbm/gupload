package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prajwalbharadwajbm/gupload/internal/config"
	"github.com/prajwalbharadwajbm/gupload/internal/db"
	"github.com/prajwalbharadwajbm/gupload/internal/logger"
)

const version = "1.0.0"

func init() {
	config.LoadConfigs()
	initializeGlobalLogger()
	loadDatabaseClient()
	logger.Log.Info("loaded all configs")
}

func initializeGlobalLogger() {
	env := config.AppConfigInstance.GeneralConfig.Env
	logLevel := config.AppConfigInstance.GeneralConfig.LogLevel
	logger.InitializeGlobalLogger(logLevel, env, version+"-gupload-service")
	logger.Log.Info("loaded the global logger")
}

func loadDatabaseClient() {
	db.GetClient()
}

func main() {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.AppConfigInstance.GeneralConfig.Port),
		Handler:      Routes(),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		logger.Log.Fatal("failed to serve http server", err)
	}
}
