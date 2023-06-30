package main

import (
	"delegationz/pkg/services/api"
	"delegationz/pkg/services/db"
	"delegationz/pkg/services/importer"
	"log"
	"os"
)

var appName = "delegationz-importer"

// dbURL must be in a dedicated config struct and be loaded securely via env or secret manager service
func dbURL() string {
	if v, ok := os.LookupEnv("DATABASE_URL"); ok && v != "" {
		return v
	}
	return "postgres://postgres:supersecret@localhost:5432/dev"
}

func main() {
	log.Printf("[INFO] Starting %s v%s", appName, api.VERSION)
	dbclient := db.Get(dbURL())
	importer.Run(dbclient, 800, false, true, true)
}
