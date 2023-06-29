package main

import (
	"delegationz/pkg/services/api"
	"delegationz/pkg/services/db"
)

var dbURL = "postgres://postgres:supersecret@localhost:5432/dev"

func main() {
	dbClient := db.Get(dbURL)
	api.Serve("8080", dbClient)
}
