package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	ErrSomethingBad = echo.NewHTTPError(http.StatusInternalServerError, "Something bad happened, try again later")
)

func ErrBadParam(param string) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid parameter [%s]", param))
}
