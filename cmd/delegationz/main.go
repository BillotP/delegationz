package main

import (
	"delegationz/pkg/services/api"
	"delegationz/pkg/services/db"
	"delegationz/pkg/services/importer"
)

var dbURL = "postgres://postgres:supersecret@localhost:5432/dev"

func main() {
	dbClient := db.Get(dbURL)
	go api.Serve("8080", dbClient)
	importer.Run(dbClient, 800)
}
