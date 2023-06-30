package api

import (
	"database/sql"
	"delegationz/pkg/repository"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func Serve(port string, withapppath *string, dbClient *sql.DB) error {
	srv := echo.New()
	srv.HideBanner = true
	srv.HidePort = true
	// Create a custom logger with console format
	logger := log.New(os.Stdout, "[Echo] ", log.LstdFlags)

	// Configure Echo to use the custom logger
	srv.Logger.SetOutput(logger.Writer())
	// Custom time format
	timeCustomFormat := "2006/01/02 15:04:05"
	// Rate limit memory store (20req/s)
	limiterStore := middleware.NewRateLimiterMemoryStore(20)
	// Middleware
	srv.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output:           logger.Writer(),
		Format:           "${time_custom} [${method}] ${uri} ${status} ${latency_human}\n",
		CustomTimeFormat: timeCustomFormat,
	}))
	srv.Use(middleware.CORS())
	srv.Use(middleware.Gzip())
	srv.Use(middleware.Recover())
	srv.Use(middleware.RateLimiter(limiterStore))

	srv.GET("/health", HealthHandler)
	srv.GET("/xtz/sync", func(c echo.Context) error {
		lastSyncedItem, err := repository.Delegations(
			qm.OrderBy("id desc"),
		).One(c.Request().Context(), dbClient)
		if err != nil {
			log.Printf("[ERROR] Failed to get last delegation : %v", err)
			return c.JSON(ErrSomethingBad.Code, ErrSomethingBad)
		}
		return c.JSON(http.StatusOK, lastSyncedItem)
	})
	srv.GET("/xtz/delegations", DelegationsHandler(dbClient))
	if withapppath != nil && *withapppath != "" {
		log.Printf("[INFO] Serving static app for default route from folder %s", *withapppath)
		srv.Static("*", *withapppath)
	}
	log.Printf("[INFO] Http server started on [::]:%s", port)
	return srv.Start(":8080")
}
