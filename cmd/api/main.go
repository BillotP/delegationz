package main

import (
	"delegationz/pkg/services/tzkt"
	"fmt"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/labstack/echo/v4"
)

type delegationItem struct {
	Timestamp time.Time `json:"timestamp"`
	Amount    string    `json:"amount"`
	Delegator string    `json:"delegator"`
	Block     string    `json:"block"`
}

type delegationsResponse struct {
	Data []delegationItem `json:"data"`
}

type ByTimestamp []delegationItem

func (t ByTimestamp) Len() int           { return len(t) }
func (t ByTimestamp) Less(i, j int) bool { return t[i].Timestamp.After(t[j].Timestamp) }
func (t ByTimestamp) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }

func main() {
	srv := echo.New()
	tzktClient := tzkt.NewTzktClient()
	srv.GET("/xtz/delegations", delegationsHandler(tzktClient))
	log.Fatal(srv.Start(":8080"))
}

func delegationsHandler(tzktClient *tzkt.TzktClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		twoHoursAgo := time.Now().Add(-2 * time.Hour)
		dd, err := tzktClient.Delegations(&tzkt.Filters{
			TimestampGe: twoHoursAgo.Format(time.RFC3339),
		}, &tzkt.Pagination{
			Limit: 2,
		})
		if err != nil {
			log.Printf("[ERROR] %+v", err)
			// Create an HTTPError with the desired status code and error message
			httpErr := echo.NewHTTPError(http.StatusInternalServerError, "An error occurred")
			// Return the HTTPError as a JSON response
			return c.JSON(httpErr.Code, httpErr)
		}
		var out = []delegationItem{}
		for _, el := range dd.Items {
			if el == nil {
				continue
			}
			out = append(out, delegationItem{
				Timestamp: el.Timestamp,
				Amount:    fmt.Sprintf("%d", el.Amount),
				Delegator: el.NewDelegate.Address,
				Block:     el.Block,
			})
		}
		// Most recent first
		sort.Sort(ByTimestamp(out))
		return c.JSON(http.StatusOK, delegationsResponse{Data: out})
	}
}
