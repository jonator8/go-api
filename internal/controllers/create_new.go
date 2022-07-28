package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/jonator8/go-api/internal"
	"github.com/jonator8/go-api/internal/commands"
	"net/http"
)

type CreateNewRequest struct {
	Title string `json:"title" validate:"required"`
	Body  string `json:"body" validate:"required"`
}

func CreateNewController(app *internal.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateNewRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			ErrorResponse(fmt.Errorf("create_new_cotroller: error formating json: %s", err), w)
			return
		}

		err = app.Validator.Struct(req)
		if err != nil {
			ErrorResponse(internal.TranslateErrors(err, *app.Translator), w)
		}

		response, err := app.Commands.CreateNew.Handle(r.Context(), commands.CreateNewCommand{
			Title: req.Title,
			Body:  req.Body,
		})
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
