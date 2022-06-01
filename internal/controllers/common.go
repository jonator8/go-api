package controllers

import (
	"encoding/json"
	"net/http"
)

type StatusError struct {
	Code int    `json:"code"`
	Err  string `json:"err"`
}

func ErrorResponse(err error, w http.ResponseWriter) {
	response := StatusError{
		Code: http.StatusBadRequest,
		Err:  err.Error(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Code)
	jsonResponse, _ := json.Marshal(response)
	w.Write(jsonResponse)
}
