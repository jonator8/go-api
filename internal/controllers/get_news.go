package controllers

import (
	"encoding/json"
	"github.com/jonator8/go-api/internal"
	"net/http"
)

func GetNewsController(app *internal.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response, err := app.Queries.GetNews.Handle(r.Context())
		if err != nil {
			ErrorResponse(err, w)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		if err = json.NewEncoder(w).Encode(response); err != nil {
			ErrorResponse(err, w)
			return
		}
	}
}
