package api

import (
	"database/sql"
	"delegationz/pkg/services/frontend"
	"log"
	"os"
	"strings"
	"time"

	"github.com/getsentry/sentry-go"
	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Serve(port string, withfrontend bool, dbClient *sql.DB) error {
	srv := echo.New()
	srv.HideBanner = true
	srv.HidePort = true
	// Create a custom logger with console format
	logger := log.New(os.Stdout, "[Echo] ", log.LstdFlags)
	// To initialize Sentry's handler, you need to initialize Sentry itself beforehand
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://615b5145294c40e7af965940325d89fb@app.glitchtip.com/3513",
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
		// The sampling rate for profiling is relative to TracesSampleRate:
		ProfilesSampleRate: 1.0,
		EnableTracing:      true,
		AttachStacktrace:   true,
	}); err != nil {
		log.Printf("[WARN] Sentry initialization failed: %v\n", err)
	}
	defer sentry.Flush(2 * time.Second)
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
	srv.GET("/xtz/sync", SyncHandler(dbClient))
	srv.GET("/xtz/delegations", DelegationsHandler(dbClient))
	if withfrontend {
		log.Printf("[INFO] Serving SPA frontend for default route")
		srv.Use(middleware.StaticWithConfig(middleware.StaticConfig{
			Filesystem: frontend.BuildHTTPFS(),
			HTML5:      true,
			Skipper: func(c echo.Context) bool {
				return strings.HasPrefix(c.Request().URL.Path, "/__")
			},
		}))
	}

	// Once it's done, you can attach the handler as one of your middleware
	srv.Use(sentryecho.New(sentryecho.Options{
		Repanic: true,
	}))
	log.Printf("[INFO] Http server started on [::]:%s", port)
	return srv.Start(":8080")
}
