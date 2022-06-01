package main

import (
	"context"
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/jonator8/go-api/internal"
	"github.com/jonator8/go-api/internal/config"
	"github.com/jonator8/go-api/internal/controllers"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"log"
	"net/http"
)

func main() {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()
	logger := log.Logger{}

	conf, err := config.GetConfiguration()
	if err != nil {
		logger.Fatal("ERROR: ", err)
	}

	dsn := "postgres://" +
		conf.Db.User +
		":" +
		conf.Db.Password +
		"@" +
		conf.Db.Host +
		":" +
		conf.Db.Port +
		"/" +
		conf.Db.Database +
		"?sslmode=disable"

	sqlDb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqlDb, pgdialect.New())
	cmd := internal.NewCommandsBootstrap(db)

	app := internal.NewApp(&logger, cmd)

	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/v1").Subrouter()
	apiRouter.HandleFunc("/healthcheck", controllers.HealthCheckController(app)).Methods(http.MethodGet)
	apiRouter.HandleFunc("/seed", controllers.SeedController(app)).Methods(http.MethodPost)

	err = http.ListenAndServe(":"+conf.App.Port, router)
	if err != nil {
		logger.Fatal("ERROR: ", err)
	}
}
