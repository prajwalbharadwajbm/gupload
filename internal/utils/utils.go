package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func FetchDataFromRequestBody[T any](request *http.Request) (T, error) {
	var data T

	body, err := io.ReadAll(request.Body)
	if err != nil {
		return data, fmt.Errorf("unable to read request body: %w", err)
	}
	defer request.Body.Close()

	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, fmt.Errorf("unable to unmarshal request body: %w", err)
	}
	return data, nil
}

func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func FormatBytes(sizeInBytes int64) string {
	const (
		KB = 1024
		MB = 1024 * KB
		GB = 1024 * MB
		TB = 1024 * GB
	)

	var unit string
	var value float64

	switch {
	case sizeInBytes >= TB:
		unit = "TB"
		value = float64(sizeInBytes) / TB
	case sizeInBytes >= GB:
		unit = "GB"
		value = float64(sizeInBytes) / GB
	case sizeInBytes >= MB:
		unit = "MB"
		value = float64(sizeInBytes) / MB
	case sizeInBytes >= KB:
		unit = "KB"
		value = float64(sizeInBytes) / KB
	default:
		return fmt.Sprintf("%d B", sizeInBytes)
	}

	return fmt.Sprintf("%.2f %s", value, unit)
}
