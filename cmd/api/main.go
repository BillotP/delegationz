package main

import (
	"delegationz/pkg/api"
	"delegationz/pkg/db"
	"log"
	"os"
)

var appName = "delegationz-api"

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

func main() {
	log.Printf("[INFO] Starting %s v%s", appName, api.VERSION)
	if err := db.Init(dbURL()); err != nil {
		log.Printf("[ERROR] Failed to connect to db : %s\n", err)
		os.Exit(1)
	}
	dbClient := db.Get()
	api.Serve(listeningPort(), false, dbClient)
}
