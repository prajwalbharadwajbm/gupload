package utils

import "os"

func GetEnv(key, fallback string) any {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
