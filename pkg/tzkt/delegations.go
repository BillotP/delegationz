package tzkt

import (
	"log"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

func mergeURLValues(maps ...url.Values) url.Values {
	result := make(url.Values)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

func (c *TzktClient) Delegations(filters *Filters, page *Pagination) (*DelegationItems, error) {
	var delegations []*DelegationItem
	queryParams := url.Values{}
	endpoint := "operations/delegations"
	if filters != nil {
		v, err := query.Values(filters)
		if err != nil {
			return nil, err
		}
		queryParams = v
	}
	if page != nil {
		v, err := query.Values(page)
		if err != nil {
			return nil, err
		}
		queryParams = mergeURLValues(queryParams, v)
	}
	endpoint += "?" + queryParams.Encode()
	if c.loglevel > 1 {
		log.Printf("[DEBUG] Will GET %s", endpoint)
	}
	err := c.do(http.MethodGet, endpoint, &delegations)
	if err != nil {
		return nil, err
	}
	res := &DelegationItems{
		delegations,
		len(delegations) == defaultLimit,
	}
	if c.loglevel > 2 {
		log.Printf("[DEBUG] GOT %+v\n", res)
	}
	return res, nil
}
