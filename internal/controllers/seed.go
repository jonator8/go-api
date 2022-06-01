package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/jonator8/go-api/internal"
	"net/http"
)

type SeedRequest struct {
	News     bool `json:"news"`
	Comments bool `json:"comments"`
}

func SeedController(app *internal.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req SeedRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			ErrorResponse(fmt.Errorf("seed_cotroller: error formating json: %s", err), w)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		jsonResponse, err := json.Marshal(req)
		if err != nil {
			app.Logger.Fatal("send_controller: ", err)
		}

		_, err = w.Write(jsonResponse)
		if err != nil {
			app.Logger.Fatal("send_controller: ", err)
		}
	}
}
