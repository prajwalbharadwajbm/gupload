package db

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "github.com/lib/pq"
	"github.com/prajwalbharadwajbm/gupload/internal/config"
	"github.com/prajwalbharadwajbm/gupload/internal/logger"
)

const pgConnStrFormat = "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"

var (
	client *sql.DB
	once   sync.Once
)

// GetClient returns the database client, initializing it if necessary
func GetClient() *sql.DB {
	once.Do(func() {
		initializeClient()
	})
	return client
}

// Initialize the database client
func initializeClient() {
	connStr := buildConnString()

	var err error
	client, err = sql.Open("postgres", connStr)
	if err != nil {
		logger.Log.Fatal("failed to connect to db", err)
	}

	configureConnPoolParams(client)

	err = client.Ping()
	if err != nil {
		logger.Log.Fatal("failed to ping db", err)
	}
	logger.Log.Info("connected to database")
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

func configureConnPoolParams(client *sql.DB) {
	client.SetMaxOpenConns(25)
	client.SetMaxIdleConns(10)
	client.SetConnMaxLifetime(5 * time.Minute)
	client.SetConnMaxIdleTime(3 * time.Minute)
}
