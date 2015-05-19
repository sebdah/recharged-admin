package main

import (
	"fmt"
	"net/http"

	goLogging "github.com/op/go-logging"
	"github.com/sebdah/recharged-admin/config"
	"github.com/sebdah/recharged-admin/database"
	"github.com/sebdah/recharged-admin/logging"
	"github.com/sebdah/recharged-admin/models"
	"github.com/sebdah/recharged-admin/routers"
)

var log goLogging.Logger

func main() {
	// Configure logging
	logging.Setup()

	// Set the environment
	log.Info("Starting re:charged admin service")
	log.Info("Environment: %s", config.Env)

	// Create databases if needed
	database.EnsureAllDatabases()
	models.EnsureAllIndexes()

	log.Info("Starting webserver on port %d", config.Config.GetInt("port"))
	http.ListenAndServe(fmt.Sprintf(":%d", config.Config.GetInt("port")), routers.HttpInterceptor(routers.Router()))
}
