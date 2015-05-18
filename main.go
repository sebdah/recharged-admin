package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sebdah/recharged-admin/config"
	"github.com/sebdah/recharged-admin/database"
	"github.com/sebdah/recharged-admin/routers"
)

func main() {
	// Set the environment
	log.Printf("Using environment '%s'\n", config.Env)

	// Create databases if needed
	database.EnsureDatabases()
	database.EnsureIndexes()

	log.Printf("Starting webserver on port %d\n", config.Config.GetInt("port"))
	http.ListenAndServe(fmt.Sprintf(":%d", config.Config.GetInt("port")), routers.Router())
}
