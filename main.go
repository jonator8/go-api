package main

import (
	"github.com/gorilla/mux"
	"github.com/jonator8/go-api/internal"
	"github.com/jonator8/go-api/internal/config"
	"github.com/jonator8/go-api/internal/controllers"
	"log"
	"net/http"
)

func main() {
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	app := internal.NewApp(&log.Logger{})

	conf, err := config.GetConfiguration()
	if err != nil {
		app.Logger.Fatal("ERROR: ", err)
	}

	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/v1").Subrouter()
	apiRouter.HandleFunc("/healthcheck", controllers.HealthCheckController(app)).Methods(http.MethodGet)

	err = http.ListenAndServe(":"+conf.App.Port, router)
	if err != nil {
		app.Logger.Fatal("ERROR: ", err)
	}
}
