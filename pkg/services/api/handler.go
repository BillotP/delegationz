package api

import (
	"context"
	"database/sql"
	"delegationz/pkg/models"
	"delegationz/pkg/repository"
	"delegationz/pkg/services/tzkt"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// QuickDelegationsHandler return delegation items with optionnal query param filter `year=YYYY`
func QuickDelegationsHandler(tzktClient *tzkt.TzktClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		year := c.QueryParam("year")
		twoHoursAgo := time.Now().Add(-2 * time.Hour)
		currYear := int64(twoHoursAgo.Year())
		originFilters := &tzkt.Filters{
			TimestampGe: twoHoursAgo.Format(time.RFC3339),
		}
		v, ok := strconv.ParseInt(year, 10, 64)
		genesis := 2017
		if ok == nil && v > int64(genesis) && v <= currYear {
			lastDayOfYear := time.Date(int(v+1), time.January, 1, 0, 0, 0, 0, time.UTC)
			lastDayOfYear = lastDayOfYear.AddDate(0, 0, -1)
			originFilters.TimestampGe = lastDayOfYear.Format(time.RFC3339)
		} else if ok == nil && (v < int64(genesis) || v > currYear) {
			log.Printf("[WARN] Invalid year submitted: %d", v)
			return c.JSON(ErrBadParam("").Code, ErrBadParam("year"))
		}
		dd, err := tzktClient.Delegations(originFilters, &tzkt.Pagination{
			Limit: 2,
		})
		if err != nil {
			log.Printf("[ERROR] Failed to get /delegations: %+v", err)
			return c.JSON(ErrSomethingBad.Code, ErrSomethingBad)
		}
		var out = []models.DelegationItem{}
		for _, el := range dd.Items {
			if el == nil {
				continue
			}
			out = append(out, *models.NewDelegationItemFromApi(el))
		}
		// Most recent first
		sort.Sort(models.ByTimestamp(out))
		return c.JSON(http.StatusOK, models.DelegationsResponse{Data: out})
	}
}

// DelegationsHandler return delegation items with optionnal query param filter `year=YYYY`
func DelegationsHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		year := c.QueryParam("year")
		twoHoursAgo := time.Now().Add(-2 * time.Hour)
		currYear := int64(twoHoursAgo.Year())
		v, ok := strconv.ParseInt(year, 10, 64)
		genesis := 2017
		mods := []qm.QueryMod{
			qm.OrderBy("timestamp"),
			qm.Limit(100),
		}
		if ok == nil && v > int64(genesis) && v <= currYear {
			// lastDayOfYear := time.Date(int(v+1), time.January, 1, 0, 0, 0, 0, time.UTC)
			// lastDayOfYear = lastDayOfYear.AddDate(0, 0, -1)
			// originFilters.TimestampGe = lastDayOfYear.Format(time.RFC3339)
			// CREATE INDEX idx_table_datetime_year ON delegations (date_part('year', timestamp));
			mods = append(mods, qm.Where("EXTRACT(YEAR FROM timestamp) = $1", v))
		} else if ok == nil && (v < int64(genesis) || v > currYear) {
			log.Printf("[WARN] Invalid year submitted: %d", v)
			return c.JSON(ErrBadParam("").Code, ErrBadParam("year"))
		}
		dd, err := repository.Delegations(mods...).All(context.Background(), db)
		if err != nil {
			log.Printf("[ERROR] Failed to get /delegations: %+v", err)
			return c.JSON(ErrSomethingBad.Code, ErrSomethingBad)
		}
		// Note : Serializiation should be done database side and result stored
		// in cache to avoid the below operation.
		var out = []models.DelegationItem{}
		for _, el := range dd {
			if el == nil {
				continue
			}
			out = append(out, models.DelegationItem{
				Timestamp: el.Timestamp,
				Amount:    fmt.Sprintf("%d", el.Amount),
				Delegator: el.Delegator,
				Block:     fmt.Sprintf("%d", el.BlockLevel),
			})
		}
		return c.JSON(http.StatusOK, models.DelegationsResponse{Data: out})
	}
}
