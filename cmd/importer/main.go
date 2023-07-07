package main

import (
	"delegationz/pkg/api"
	"delegationz/pkg/db"
	"delegationz/pkg/importer"
	"delegationz/pkg/tzkt"
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
	if err := db.Init(dbURL()); err != nil {
		log.Printf("[ERROR] Failed to connect to db : %s\n", err)
		os.Exit(1)
	}
	dbClient := db.Get()
	apiclient := tzkt.NewTzktClient()
	importr := importer.New(dbClient, apiclient,
		importer.WithVerbose(true),
		importer.WithPageSize(800),
		importer.WithWatch(true),
		importer.WithReset(false))
	importr.Run()
}
