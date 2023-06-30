package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestDelegationsHandler(t *testing.T) {
	e := echo.New()
	// Create a mock database connection
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Create a request with query parameters
	req := httptest.NewRequest(http.MethodGet, "/delegations?year=2022&limit=50&page=1", nil)

	// Create a response recorder
	rec := httptest.NewRecorder()

	// Set up the request context
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("year", "limit", "page")
	ctx.SetParamValues("2022", "50", "1")

	// Set up a predictable date
	ddd := time.Date(2019, time.May, 19, 1, 2, 3, 4, time.UTC)
	// Set up the mock expectations
	rows := sqlmock.NewRows(
		[]string{"id", "timestamp", "amount", "delegator", "block_height", "block_hash"}).
		AddRow("1", ddd.AddDate(-1, 0, 0), "100", "delegator1", "1", "oxab").
		AddRow("2", ddd, "200", "delegator2", "2", "oxbc")

	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	// Call the handler function
	err = DelegationsHandler(db)(ctx)

	// Assert that no error occurred
	assert.NoError(t, err)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, rec.Code)

	// Assert the response body
	expectedResponse := `{"data":[{"timestamp":"2018-05-19T01:02:03.000000004Z","amount":"100","delegator":"delegator1","block":"0"},{"timestamp":"2019-05-19T01:02:03.000000004Z","amount":"200","delegator":"delegator2","block":"0"}]}` // Replace [...] with your expected response JSON
	assert.Equal(t, expectedResponse, strings.TrimSpace(rec.Body.String()))

	// Assert the mock expectations
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDelegationsHandler_DateInFuture(t *testing.T) {
	e := echo.New()
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Create a request with query parameters
	req := httptest.NewRequest(http.MethodGet, "/delegations?year=2025&limit=50&page=1", nil)

	// Create a response recorder
	rec := httptest.NewRecorder()

	// Set up the request context
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("year", "limit", "page")
	ctx.SetParamValues(fmt.Sprintf("%d", time.Now().Year()+2), "50", "1")

	// Call the handler function
	err = DelegationsHandler(db)(ctx)

	// Assert that no error occurred
	assert.NoError(t, err)

	// Assert the response status code
	assert.Equal(t, ErrBadParam("").Code, rec.Code)

	// Assert the response body
	expectedResponse := `{"message":"` + ErrBadParam("year").Message.(string) + `"}` // Empty data array since there are no rows in the result
	assert.Equal(t, expectedResponse, strings.TrimSpace(rec.Body.String()))

	// Assert the mock expectations
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestHealthHandler(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)

	// Create a response recorder
	rec := httptest.NewRecorder()

	// Set up the request context
	ctx := e.NewContext(req, rec)

	// Call the handler function
	err := HealthHandler(ctx)

	// Assert that no error occurred
	assert.NoError(t, err)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, rec.Code)

	// Assert the response body
	assert.Empty(t, rec.Body.String())
}
