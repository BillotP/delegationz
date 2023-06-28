package tzkt

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"reflect"
)

func (c *TzktClient) Delegations(filters *Filters, page *Pagination) (*DelegationItems, error) {
	var delegations []*DelegationItem
	queryParams := url.Values{}
	endpoint := "operations/delegations"
	if filters != nil {
		filtersValue := reflect.ValueOf(filters).Elem()
		filtersType := filtersValue.Type()

		for i := 0; i < filtersValue.NumField(); i++ {
			field := filtersValue.Field(i)
			if field.String() != "" {
				queryParams.Set(filterParams[filtersType.Field(i).Name], field.String())
			}
		}

	}
	if page != nil {
		paginationValue := reflect.ValueOf(page).Elem()
		paginationType := paginationValue.Type()

		for i := 0; i < paginationValue.NumField(); i++ {
			field := paginationValue.Field(i)
			if field.String() != "" {
				vv := fmt.Sprintf("%d", field.Int())
				queryParams.Set(paginationParams[paginationType.Field(i).Name], vv)
				if paginationType.Field(i).Name == "Limit" {
					defaultLimit = int(field.Int())
				}
			}
		}
	}
	endpoint += "?" + queryParams.Encode()
	log.Printf("[DEBUG] Will GET %s", endpoint)
	err := c.do(http.MethodGet, endpoint, &delegations)
	if err != nil {
		return nil, err
	}
	res := &DelegationItems{
		delegations,
		len(delegations) == defaultLimit,
	}
	log.Printf("[DEBUG] GOT %+v\n", res)
	return res, nil
}
