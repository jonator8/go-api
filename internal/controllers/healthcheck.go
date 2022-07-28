package controllers

import (
	"encoding/json"
	"github.com/jonator8/go-api/internal"
	"net/http"
)

type HealthCheckResponse struct {
	Message string `json:"message"`
}

func HealthCheckController(app *internal.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := HealthCheckResponse{
			Message: "API is up and running",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			app.Logger.Error("healthCheckController: ", err)
		}

		_, err = w.Write(jsonResponse)
		if err != nil {
			app.Logger.Error("healthCheckController: ", err)
		}
	}
}
