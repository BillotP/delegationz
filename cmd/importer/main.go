package main

import (
	"delegationz/pkg/services/db"
	"delegationz/pkg/services/importer"
)

var dbURL = "postgres://postgres:supersecret@localhost:5432/dev"

func main() {
	dbclient := db.Get(dbURL)
	importer.Run(dbclient, 800)
}
