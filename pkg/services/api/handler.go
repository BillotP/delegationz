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

// DelegationsHandler handle `/xtz/delegations?year=optionnalYearformatYYYY`
func DelegationsHandler(tzktClient *tzkt.TzktClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		year := c.QueryParam("year")
		twoHoursAgo := time.Now().Add(-2 * time.Hour)
		currYear := int64(twoHoursAgo.Year())
		v, ok := strconv.ParseInt(year, 10, 64)
		genesis := 2017
		if ok == nil && v > int64(genesis) && v <= currYear {
			twoHoursAgo = time.Date(int(v), time.January, 1, 0, 0, 0, 0, time.UTC)
		}
		originFilters := &tzkt.Filters{
			TimestampGe: twoHoursAgo.Format(time.RFC3339),
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
