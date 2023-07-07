package importer

import (
	"delegationz/pkg/tzkt"
	"time"
)

type bulkdelegatorsItems struct {
	timestamps []time.Time
	delegators []string
}

type bulkdelegationsItems struct {
	ids           []int64
	amounts       []int64
	block_hashes  []string
	block_heights []int
	bulkdelegatorsItems
}

func toBulkInsert(delegation *tzkt.DelegationItems) *bulkdelegationsItems {
	if delegation == nil || len(delegation.Items) == 0 {
		return nil
	}
	ids := make([]int64, len(delegation.Items))
	timestamps := make([]time.Time, len(delegation.Items))
	amounts := make([]int64, len(delegation.Items))
	delegators := make([]string, len(delegation.Items))
	block_hashes := make([]string, len(delegation.Items))
	block_heights := make([]int, len(delegation.Items))
	for i, row := range delegation.Items {
		ids[i] = row.ID
		timestamps[i] = row.Timestamp
		amounts[i] = row.Amount
		delegators[i] = row.Sender.Address
		block_hashes[i] = row.Block
		block_heights[i] = row.Level
	}
	return &bulkdelegationsItems{
		ids,
		amounts,
		block_hashes,
		block_heights,
		bulkdelegatorsItems{timestamps,
			delegators},
	}
}
