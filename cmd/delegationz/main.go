package main

import (
	"delegationz/pkg/services/api"
	"delegationz/pkg/services/db"
	"delegationz/pkg/services/importer"
	"log"
	"os"
	"strings"
)

var appName = "delegationz"

// dbURL must be in a dedicated config struct and be loaded securely via env or secret manager service
func dbURL() string {
	if v, ok := os.LookupEnv("DATABASE_URL"); ok && v != "" {
		return v
	}
	return "postgres://postgres:supersecret@localhost:5432/dev"
}

// listeningPort must be in a dedicated config struct and be loaded via env at runtime
func listeningPort() string {
	if v, ok := os.LookupEnv("PORT"); ok && v != "" {
		return v
	}
	return "8080"
}

// staticPath must be in a dedicated config struct and be loaded via env at runtime
func staticPath() *string {
	if v, ok := os.LookupEnv("STATIC_PATH"); ok && v != "" {
		return &v
	}
	return nil
}

// verbose must be in a dedicated config struct and be loaded via env at runtime
func verbose() bool {
	if v, ok := os.LookupEnv("STATIC_PATH"); ok && v != "" {
		v = strings.ToLower(v)
		return v == "true" || v == "1"
	}
	return false
}

// The above (mostly duplicated) methods surely need a refacto  and a proper config package!

func main() {
	log.Printf("[INFO] Starting %s v%s", appName, api.VERSION)
	dbClient := db.Get(dbURL())
	// Starting API on goroutine
	go api.Serve(listeningPort(), staticPath(), dbClient)
	// Running delegation updater on main routine
	importer.Run(dbClient, 800, true, false, verbose())
}
