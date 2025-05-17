package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/prajwalbharadwajbm/gupload/internal/config"
	"github.com/prajwalbharadwajbm/gupload/internal/logger"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	envelope := map[string]interface{}{
		"status": "available",
		"application-details": map[string]interface{}{
			"version":     "1.0.0",
			"environment": config.AppConfigInstance.GeneralConfig.Env,
		},
	}
	healthCheckObj, err := json.Marshal(envelope)
	if err != nil {
		logger.Log.Error("failed to marshal health check response", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(healthCheckObj)
}
