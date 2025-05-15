package main

import (
	"github.com/prajwalbharadwajbm/gupload/internal/config"
	"github.com/prajwalbharadwajbm/gupload/internal/db"
	"github.com/prajwalbharadwajbm/gupload/internal/logger"
)

const version = "1.0.0"

func init() {
	loadAllConfigs()
	initializeGlobalLogger()
	loadDatabaseClient()
}

func initializeGlobalLogger() {
	env := config.AppConfigInstance.GeneralConfig.Env
	logLevel := config.AppConfigInstance.GeneralConfig.LogLevel
	logger.InitializeGlobalLogger(logLevel, env, version+"-file-upload-service")
	logger.Log.Info("loaded the global logger")
}

func loadAllConfigs() {
	config.LoadConfigs()
	logger.Log.Info("loaded all configs")
}

func loadDatabaseClient() {
	db.GetClient()
}

func main() {

}
