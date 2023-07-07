package api

import (
	"delegationz/pkg/tzkt"
	"fmt"
	"time"
)

type DelegationItem struct {
	Timestamp time.Time `json:"timestamp"`
	Amount    string    `json:"amount"`
	Delegator string    `json:"delegator"`
	Block     string    `json:"block"`
}

type ItemsResponse struct {
	Data interface{} `json:"data"`
}

type ByTimestamp []DelegationItem

func (t ByTimestamp) Len() int           { return len(t) }
func (t ByTimestamp) Less(i, j int) bool { return t[i].Timestamp.After(t[j].Timestamp) }
func (t ByTimestamp) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }

func NewDelegationItemFromApi(itm *tzkt.DelegationItem) *DelegationItem {
	return &DelegationItem{
		Timestamp: itm.Timestamp,
		Amount:    fmt.Sprintf("%d", itm.Amount),
		Delegator: itm.NewDelegate.Address,
		Block:     fmt.Sprintf("%d", itm.Level),
	}
}
