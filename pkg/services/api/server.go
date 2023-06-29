package api

import (
	"database/sql"
	"delegationz/pkg/services/tzkt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Serve(port string, dbClient *sql.DB) error {
	srv := echo.New()
	srv.Use(middleware.Logger())
	tzktClient := tzkt.NewTzktClient()
	srv.GET("/xtz/delegations", DelegationsHandler(dbClient))
	srv.GET("/v0/xtz/delegations", QuickDelegationsHandler(tzktClient))
	return srv.Start(":8080")
}
