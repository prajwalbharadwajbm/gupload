package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/prajwalbharadwajbm/gupload/internal/utils"
)

type GeneralConfig struct {
	Env      string
	LogLevel string
	Port     string
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
	AppConfigInstance.GeneralConfig.Env = utils.GetEnv("APP_ENV", "dev").(string)
	AppConfigInstance.GeneralConfig.LogLevel = utils.GetEnv("LOG_LEVEL", "info").(string)
	AppConfigInstance.GeneralConfig.Port = utils.GetEnv("PORT", "8080").(string)
}

func loadDatabaseConfigs() {
	AppConfigInstance.DB.Host = utils.GetEnv("DB_HOST", "").(string)
	AppConfigInstance.DB.Port = (utils.GetEnv("DB_PORT", "")).(int)
	AppConfigInstance.DB.DBname = utils.GetEnv("DB_NAME", "").(string)
}
