package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jonator8/go-api/internal/controllers"
	"net/http"
)

func main() {
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()

	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api/v1").Subrouter()
	apiRouter.HandleFunc("/healthcheck", controllers.HealthCheck).Methods(http.MethodGet)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
}
