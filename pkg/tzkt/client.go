package tzkt

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var defaultLimit = 100

type ITzktClient interface {
	Delegations(filters *Filters, page *Pagination) (*DelegationItems, error)
}
type TzktClient struct {
	cli      *http.Client
	baseURL  string
	loglevel int
}

func NewTzktClient() *TzktClient {
	return &TzktClient{
		cli:      http.DefaultClient,
		loglevel: 1,
		baseURL:  "https://api.tzkt.io/v1",
	}
}

func (c *TzktClient) do(method, endpoint string, holder interface{}) error {
	url := fmt.Sprintf("%s/%s", c.baseURL, endpoint)
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}

	res, err := c.cli.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}
	if holder != nil {
		err = json.NewDecoder(res.Body).Decode(holder)
		if err != nil {
			return err
		}
	}

	return nil
}
