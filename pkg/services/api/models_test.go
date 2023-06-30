package api

import (
	"delegationz/pkg/services/tzkt"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewDelegationItemFromApi(t *testing.T) {
	itm := &tzkt.DelegationItem{
		Timestamp: time.Now(),
		Amount:    1000,
		NewDelegate: tzkt.NewDelegate{
			Address: "delegator123",
		},
		Level: 500,
	}

	result := NewDelegationItemFromApi(itm)

	expected := &DelegationItem{
		Timestamp: itm.Timestamp,
		Amount:    fmt.Sprintf("%d", itm.Amount),
		Delegator: itm.NewDelegate.Address,
		Block:     fmt.Sprintf("%d", itm.Level),
	}

	assert.Equal(t, expected, result)
}

func TestByTimestamp_Len(t *testing.T) {
	timestamps := ByTimestamp{
		{Timestamp: time.Now()},
		{Timestamp: time.Now()},
		{Timestamp: time.Now()},
	}

	expectedLength := 3

	result := timestamps.Len()

	assert.Equal(t, expectedLength, result)
}

func TestByTimestamp_Less(t *testing.T) {
	timestamps := ByTimestamp{
		{Timestamp: time.Now().Add(time.Hour)},
		{Timestamp: time.Now()},
	}

	result1 := timestamps.Less(0, 1)
	result2 := timestamps.Less(1, 0)

	assert.True(t, result1)
	assert.False(t, result2)
}

func TestByTimestamp_Swap(t *testing.T) {
	dd := time.Now()
	rr := dd.Add(time.Hour)
	timestamps := ByTimestamp{
		{Timestamp: dd},
		{Timestamp: rr},
	}

	expectedTimestamps := ByTimestamp{
		{Timestamp: rr},
		{Timestamp: dd},
	}

	timestamps.Swap(0, 1)

	assert.Equal(t, expectedTimestamps, timestamps)
}
