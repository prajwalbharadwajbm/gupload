package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/prajwalbharadwajbm/gupload/internal/logger"
)

func FetchDataFromRequestBody[T any](request *http.Request) (T, error) {
	var data T

	body, err := io.ReadAll(request.Body)
	if err != nil {
		logger.Log.Error("unable to read request body", err)
		return data, err
	}
	defer request.Body.Close()

	err = json.Unmarshal(body, &data)
	if err != nil {
		logger.Log.Error("unable to unmarshal request body", err)
		return data, err
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
