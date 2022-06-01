package main

import (
	"github.com/gorilla/mux"
	"github.com/jonator8/go-api/internal"
	"github.com/jonator8/go-api/internal/controllers"
	"log"
	"net/http"
)

func main() {
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	app := internal.NewApp(&log.Logger{})

	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/v1").Subrouter()
	apiRouter.HandleFunc("/healthcheck", controllers.HealthCheckController(app)).Methods(http.MethodGet)

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		app.Logger.Fatal("ERROR: ", err)
	}
}
