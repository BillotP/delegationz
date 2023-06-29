package main

import (
	"delegationz/pkg/services/api"
	"delegationz/pkg/services/tzkt"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	srv := echo.New()
	tzktClient := tzkt.NewTzktClient()
	srv.GET("/xtz/delegations", api.DelegationsHandler(tzktClient))
	log.Fatal(srv.Start(":8080"))
}
