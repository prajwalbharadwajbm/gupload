package db

import (
	"database/sql"
	"fmt"

	"github.com/prajwalbharadwajbm/gupload/internal/config"
	"github.com/prajwalbharadwajbm/gupload/internal/logger"
)

const pgConnStrFormat = "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"

func GetClient() *sql.DB {
	connStr := buildConnString()

	client, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Log.Fatal("failed to connect to db", err)
	}

	err = client.Ping()
	if err != nil {
		logger.Log.Fatal("failed to ping db", err)
	}
	logger.Log.Info("connected to database")

	return client
}

func buildConnString() string {
	return fmt.Sprintf(
		pgConnStrFormat,
		config.AppConfigInstance.DB.Host,
		config.AppConfigInstance.DB.Port,
		config.AppConfigInstance.DB.User,
		config.AppConfigInstance.DB.Password,
		config.AppConfigInstance.DB.DBname,
	)
}
