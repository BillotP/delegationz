package main

import (
	"context"
	"delegationz/pkg/db"
	"delegationz/pkg/services/tzkt"
	"log"
	"sync"
	"time"
)

var dbURL = "postgres://postgres:supersecret@localhost:5432/dev"

func main() {
	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Load tzkt api client
	cli := tzkt.NewTzktClient()

	// Create a channel to communicate delegation objects between goroutines
	delegationChannel := make(chan *tzkt.DelegationItems)

	// Set the number of goroutines to use
	numGoroutines := 5

	// Launch goroutines to fetch and save delegations
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go saveDelegations(delegationChannel, &wg)
	}

	// Fetch the paginated items concurrently
	pageSize := 600
	pagination := tzkt.Pagination{
		Limit:  pageSize,
		Offset: 0,
	}
	flters := tzkt.Filters{
		TypeEq: "applied",
	}
	for i := 0; i < numGoroutines; i++ {
		resp, err := cli.Delegations(&flters, &pagination)
		if err != nil {
			log.Printf("[ERROR] error making API request:%v\n", err)
			break
		}

		// Send items to the channel for goroutines to process and save
		delegationChannel <- resp

		// Check if there are more pages
		if !resp.HasMore {
			break
		}

		pagination.Offset++
	}
	// Close the delegation channel to indicate that no more delegations will be sent
	close(delegationChannel)

	// Wait for all goroutines to finish
	wg.Wait()

}

func saveDelegations(delegationChannel <-chan *tzkt.DelegationItems, wg *sync.WaitGroup) {
	defer wg.Done()

	for delegation := range delegationChannel {
		// Perform the logic to save the delegation to the database or perform any other operations
		// For simplicity, we will just print the delegation here
		ctx := context.Background()
		log.Printf("[DEBUG] Will save %v delegation\n", len(delegation.Items))
		if len(delegation.Items) == 0 {
			continue
		}

		ids := make([]int64, len(delegation.Items))
		timestamps := make([]time.Time, len(delegation.Items))
		amounts := make([]int64, len(delegation.Items))
		delegators := make([]string, len(delegation.Items))
		blocks := make([]string, len(delegation.Items))

		for i, row := range delegation.Items {
			ids[i] = row.ID
			timestamps[i] = row.Timestamp
			amounts[i] = row.Amount
			delegators[i] = row.NewDelegate.Address
			blocks[i] = row.Block
		}
		rr, err := db.Get(dbURL).ExecContext(ctx, `
		INSERT INTO delegations
		(id, timestamp, amount, delegator, block)
		(SELECT  * FROM UNNEST($1::bigint[], $2::timestamp[], $3::bigint[], $4::varchar[], $5::text[]))
		ON CONFLICT (id) DO UPDATE SET delegator = EXCLUDED.delegator
		`, ids, timestamps, amounts, delegators, blocks)
		if err != nil {
			log.Printf("[ERROR] Failed to save : %s\n", err)
			return
		}
		cnt, _ := rr.RowsAffected()
		log.Printf("[INFO] %d delegations saved", cnt)
	}
}
