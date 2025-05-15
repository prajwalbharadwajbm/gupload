package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/prajwalbharadwajbm/gupload/internal/utils"
)

type GeneralConfig struct {
	Env      string
	LogLevel string
	Port     int
}

type appConfig struct {
	GeneralConfig GeneralConfig
	DB            DB
	JWTSecret     string
}

type DB struct {
	Host     string
	Port     int
	User     string
	Password string
	DBname   string
}

func LoadConfigs() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file:", err)
	}

	loadGeneralConfigs()
	loadDatabaseConfigs()
}

var AppConfigInstance appConfig

func loadGeneralConfigs() {
	AppConfigInstance.GeneralConfig.Env = utils.GetEnv("APP_ENV", "dev")
	AppConfigInstance.GeneralConfig.LogLevel = utils.GetEnv("LOG_LEVEL", "info")
	portStr := utils.GetEnv("PORT", "8080")
	AppConfigInstance.GeneralConfig.Port = utils.StringToInt(portStr)
}

func loadDatabaseConfigs() {
	AppConfigInstance.DB.Host = utils.GetEnv("DB_HOST", "localhost")
	portStr := utils.GetEnv("DB_PORT", "5432")
	AppConfigInstance.DB.Port = utils.StringToInt(portStr)
	AppConfigInstance.DB.User = utils.GetEnv("DB_USER", "postgres")
	AppConfigInstance.DB.Password = utils.GetEnv("DB_PASSWORD", "")
	AppConfigInstance.DB.DBname = utils.GetEnv("DB_NAME", "fileupload")
}
