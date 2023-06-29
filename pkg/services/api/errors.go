package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	ErrSomethingBad = echo.NewHTTPError(http.StatusInternalServerError, "Something bad happened, try again later")
)
