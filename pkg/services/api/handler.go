package api

import (
	"context"
	"database/sql"
	"delegationz/pkg/repository"
	"fmt"
	"log"
	"net/http"
	"strconv"

	sentryecho "github.com/getsentry/sentry-go/echo"

	"time"

	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// DelegationsHandler return delegation items with optionnal query param filters `year=YYYY&limit=xx&page=xx`
func DelegationsHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		year := c.QueryParam("year")
		limit := c.QueryParam("limit")
		page := c.QueryParam("page")

		twoHoursAgo := time.Now().Add(-2 * time.Hour)
		currYear := int64(twoHoursAgo.Year())
		yearParam, errYear := strconv.ParseInt(year, 10, 64)
		limitParam, errLimit := strconv.ParseInt(limit, 10, 64)
		offsetParam, errOffset := strconv.ParseInt(page, 10, 64)

		genesis := 2018
		mods := []qm.QueryMod{qm.OrderBy("timestamp desc")}
		// Year parsing
		if errYear == nil && yearParam >= int64(genesis) && yearParam <= currYear {
			mods = append(mods, qm.Where("EXTRACT(YEAR FROM timestamp) = $1", yearParam))
		} else if errYear == nil && (yearParam < int64(genesis) || yearParam > currYear) {
			errMsg := fmt.Sprintf("[WARN] Invalid year submitted: %d", yearParam)
			log.Println(errMsg)
			if hub := sentryecho.GetHubFromContext(c); hub != nil {
				hub.CaptureMessage(errMsg)
			}
			return c.JSON(ErrBadParam("").Code, ErrBadParam("year"))
		}
		// Limit parsing
		if errLimit == nil && limitParam > 0 && limitParam < 200 {
			mods = append(mods, qm.Limit(int(limitParam)))
		} else {
			mods = append(mods, qm.Limit(int(100)))
		}
		// Offset parsing
		if errOffset == nil && offsetParam >= 0 {
			if limitParam == 0 {
				limitParam = 100
			}
			mods = append(mods, qm.Offset(int(offsetParam*limitParam)))
		} else {
			mods = append(mods, qm.Offset(0))
		}
		// DB Fetching (should be targetting a read only replica)
		dd, err := repository.Delegations(mods...).All(context.Background(), db)
		if err != nil {
			log.Printf("[ERROR] Failed to get /delegations: %+v", err)
			if hub := sentryecho.GetHubFromContext(c); hub != nil {
				hub.CaptureException(err)
			}
			return c.JSON(ErrSomethingBad.Code, ErrSomethingBad)
		}
		// Serialization loop : TODO : avoid it
		var out = []DelegationItem{}
		for _, el := range dd {
			if el == nil {
				continue
			}
			out = append(out, DelegationItem{
				Timestamp: el.Timestamp,
				Amount:    fmt.Sprintf("%d", el.Amount),
				Delegator: el.Delegator,
				Block:     fmt.Sprintf("%d", el.BlockLevel),
			})
		}
		return c.JSON(http.StatusOK, DelegationsResponse{Data: out})
	}
}

func SyncHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		lastSyncedItem, err := repository.Delegations(
			qm.OrderBy("id desc"),
		).One(c.Request().Context(), db)
		if err != nil {
			log.Printf("[ERROR] Failed to get last delegation : %v", err)
			if hub := sentryecho.GetHubFromContext(c); hub != nil {
				hub.CaptureException(err)
			}
			return c.JSON(ErrSomethingBad.Code, ErrSomethingBad)
		}
		return c.JSON(http.StatusOK, lastSyncedItem)
	}
}

func HealthHandler(c echo.Context) error { return c.NoContent(http.StatusOK) }
