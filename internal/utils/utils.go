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
