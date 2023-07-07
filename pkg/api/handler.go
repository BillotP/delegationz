package api

import (
	"context"
	"database/sql"
	"delegationz/pkg/importer"
	"delegationz/pkg/repository"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	sentryecho "github.com/getsentry/sentry-go/echo"

	"time"

	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

const genesis = 2018

// DelegationsHandler return delegation items with optionnal query param filters `year=YYYY&limit=xx&page=xx`
func DelegationsHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		year := c.QueryParam("year")
		limit := c.QueryParam("limit")
		page := c.QueryParam("page")
		// A (delegator) address is formed of a prefix (tz1 tz2 tz3 tz4 and KT1 for contracts)
		//  followed by a Base58 encoded hash and terminated by a 4-byte checksum.
		delegator := c.QueryParam("delegator")

		twoHoursAgo := time.Now().Add(-2 * time.Hour)
		currYear := int64(twoHoursAgo.Year())
		yearParam, errYear := strconv.ParseInt(year, 10, 64)
		limitParam, errLimit := strconv.ParseInt(limit, 10, 64)
		offsetParam, errOffset := strconv.ParseInt(page, 10, 64)

		mods := []qm.QueryMod{qm.OrderBy("timestamp desc")}
		// Year parsing
		if errYear == nil && yearParam >= int64(genesis) && yearParam <= currYear {
			mods = append(mods, qm.Where("EXTRACT(YEAR FROM timestamp) = ?", yearParam))
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
		// Delegator parsing (should write a proper method)
		if delegator != "" && len(delegator) == 36 {
			mods = append(mods, qm.Where("delegator = ?", delegator))
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
		return c.JSON(http.StatusOK, ItemsResponse{Data: out})
	}
}

func DelegatorsHandler(db *sql.DB) echo.HandlerFunc {
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
		mods := []qm.QueryMod{qm.OrderBy("first_seen desc")}
		// Year parsing
		if errYear == nil && yearParam >= int64(genesis) && yearParam <= currYear {
			mods = append(mods, qm.Where("EXTRACT(YEAR FROM first_seen) = ?", yearParam))
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
		dd, err := repository.Delegators(mods...).All(context.Background(), db)
		if err != nil {
			log.Printf("[ERROR] Failed to get /delegators: %+v", err)
			if hub := sentryecho.GetHubFromContext(c); hub != nil {
				hub.CaptureException(err)
			}
			return c.JSON(ErrSomethingBad.Code, ErrSomethingBad)
		}
		return c.JSON(http.StatusOK, ItemsResponse{Data: dd})
	}
}

func isTop(kind string) bool {
	return strings.HasPrefix(kind, "TOP")
}
func DelegatorsStatsHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		mods := []qm.QueryMod{}
		kind := c.QueryParam("kind")
		year := c.QueryParam("year")
		// timestamp := c.QueryParam("timestamp")
		delegator := c.QueryParam("delegator")

		twoHoursAgo := time.Now().Add(-2 * time.Hour)
		currYear := int64(twoHoursAgo.Year())
		yearParam, errYear := strconv.ParseInt(year, 10, 64)
		// Kind parsing
		if kind != "" {
			if err := repository.StatKind(kind).IsValid(); err != nil {
				errMsg := fmt.Sprintf("[WARN] Invalid stat kind submitted: %d", yearParam)
				log.Println(errMsg)
				if hub := sentryecho.GetHubFromContext(c); hub != nil {
					hub.CaptureMessage(errMsg)
				}
				return c.JSON(ErrBadParam("").Code, ErrBadParam("kind"))
			}
		} else {
			kind = string(repository.StatKindTOP10VALIDATORS)
		}
		mods = append(mods, qm.Where("kind = ?", kind))

		// Year parsing
		if errYear == nil && yearParam >= int64(genesis) && yearParam <= currYear {
			mods = append(mods, qm.Where("EXTRACT(YEAR FROM timestamp) = ?", yearParam))
		} else if errYear == nil && (yearParam < int64(genesis) || yearParam > currYear) {
			errMsg := fmt.Sprintf("[WARN] Invalid year submitted: %d", yearParam)
			log.Println(errMsg)
			if hub := sentryecho.GetHubFromContext(c); hub != nil {
				hub.CaptureMessage(errMsg)
			}
			return c.JSON(ErrBadParam("").Code, ErrBadParam("year"))
		} else {
			year = fmt.Sprintf("%d", genesis)
		}

		dd, err := repository.DelegationsStats(mods...).All(context.Background(), db)
		if err != nil {
			log.Printf("[ERROR] Failed to get /delegators: %+v", err)
			if hub := sentryecho.GetHubFromContext(c); hub != nil {
				hub.CaptureException(err)
			}
			return c.JSON(ErrSomethingBad.Code, ErrSomethingBad)
		}
		// No result yet , so lets update! (slow but should append only once)
		if len(dd) == 0 || (yearParam == currYear && isTop(kind)) {
			kind := repository.StatKind(kind)
			switch kind {
			case repository.StatKindTOP10VALIDATORS, repository.StatKindTOP100VALIDATORS:
				err := importer.SaveTopKindStats(db, year, kind, true)
				if err != nil {
					log.Printf("[ERROR] Failed to update stats: %+v", err)
					if hub := sentryecho.GetHubFromContext(c); hub != nil {
						hub.CaptureException(err)
					}
					return c.JSON(ErrSomethingBad.Code, ErrSomethingBad)
				}
				return c.NoContent(http.StatusCreated)

			case repository.StatKindYEARLYVOLUME, repository.StatKindMONTHLYVOLUME:
				err := importer.SaveVolumeStats(db, kind, year, delegator, true)
				if err != nil {
					log.Printf("[ERROR] Failed to update stats: %+v", err)
					if hub := sentryecho.GetHubFromContext(c); hub != nil {
						hub.CaptureException(err)
					}
					return c.JSON(ErrSomethingBad.Code, ErrSomethingBad)
				}
				return c.NoContent(http.StatusCreated)
			}

		}
		return c.JSON(http.StatusOK, ItemsResponse{Data: dd})
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
