package api

import (
	"delegationz/pkg/models"
	"delegationz/pkg/services/tzkt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

// DelegationsHandler return delegation items with optionnal query param filter `year=YYYY`
func DelegationsHandler(tzktClient *tzkt.TzktClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		year := c.QueryParam("year")
		twoHoursAgo := time.Now().Add(-2 * time.Hour)
		currYear := int64(twoHoursAgo.Year())
		// TODO: fetch from saved datas
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
