package main

import (
	"delegationz/pkg/db"
	"delegationz/pkg/services/api"
	"delegationz/pkg/services/tzkt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var dbURL = "postgres://postgres:supersecret@localhost:5432/dev"

func main() {
	srv := echo.New()
	srv.Use(middleware.Logger())
	tzktClient := tzkt.NewTzktClient()
	dbclient := db.Get(dbURL)
	srv.GET("/v0/xtz/delegations", api.QuickDelegationsHandler(tzktClient))
	srv.GET("/xtz/delegations", api.DelegationsHandler(dbclient))
	log.Fatal(srv.Start(":8080"))
}
