package main

import (
	"delegationz/pkg/api"
	"delegationz/pkg/db"
	"delegationz/pkg/importer"
	"delegationz/pkg/tzkt"
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

// withFrontend must be in a dedicated config struct and be loaded via env at runtime
func withFrontend() bool {
	if v, ok := os.LookupEnv("FRONTEND"); ok && v != "" {
		v = strings.ToLower(v)
		return v == "true" || v == "1"
	}
	return false
}

// verbose must be in a dedicated config struct and be loaded via env at runtime
func verbose() bool {
	if v, ok := os.LookupEnv("VERBOSE"); ok && v != "" {
		v = strings.ToLower(v)
		return v == "true" || v == "1"
	}
	return false
}

// The above (mostly duplicated) methods surely need a refacto  and a proper config package!

func main() {
	log.Printf("[INFO] Starting %s v%s", appName, api.VERSION)
	if err := db.Init(dbURL()); err != nil {
		log.Printf("[ERROR] Failed to connect to db : %s\n", err)
		os.Exit(1)
	}
	dbClient := db.Get()
	// Starting API on separate goroutine
	go api.Serve(listeningPort(), withFrontend(), dbClient)
	// Running delegation updater on main routine
	apiclient := tzkt.NewTzktClient()
	importr := importer.New(dbClient, apiclient,
		importer.WithVerbose(verbose()),
		importer.WithPageSize(800),
		importer.WithWatch(true),
		importer.WithReset(false))
	importr.Run()
}
