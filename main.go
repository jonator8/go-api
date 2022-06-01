package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()

	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api/v1").Subrouter()
	apiRouter.HandleFunc("/healthcheck", HealthCheck).Methods(http.MethodGet)
	http.Handle("/", router)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	//specify status code
	w.WriteHeader(http.StatusOK)

	//update response writer
	fmt.Fprintf(w, "API is up and running")
}
