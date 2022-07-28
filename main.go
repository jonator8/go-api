package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jonator8/go-api/internal"
	"github.com/jonator8/go-api/internal/config"
	"github.com/jonator8/go-api/internal/controllers"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"net/http"
)

func main() {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := logrus.New()

	conf, err := config.GetConfiguration()
	if err != nil {
		logger.WithFields(logrus.Fields{
			"error":   err,
			"message": "failed to GetConfig",
		}).Fatal("ERROR: ", err)
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

	app := internal.NewApp(logger, cmd)

	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/v1").Subrouter()
	apiRouter.HandleFunc("/healthcheck", controllers.HealthCheckController(app)).Methods(http.MethodGet)
	apiRouter.HandleFunc("/seed", controllers.SeedController(app)).Methods(http.MethodPost)
	apiRouter.HandleFunc("/news", controllers.CreateNewController(app)).Methods(http.MethodGet)

	fmt.Println("API listening at " + conf.App.Host + ":" + conf.App.Port)
	err = http.ListenAndServe(conf.App.Host+":"+conf.App.Port, router)
	if err != nil {
		logger.Fatal("ERROR: ", err)
	}
}
