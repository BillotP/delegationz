package main

import (
	"context"
	"database/sql"
	"delegationz/pkg/db"
	"delegationz/pkg/services/tzkt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var dbURL = "postgres://postgres:supersecret@localhost:5432/dev"

func getLastID() *int64 {
	var lID int64
	err := db.Get(dbURL).QueryRow(`select MAX(id) from delegations`).Scan(&lID)
	if err == sql.ErrNoRows {
		return nil
	}
	return &lID
}

func main() {
	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Load tzkt api client
	cli := tzkt.NewTzktClient()

	// Create a channel to communicate delegation objects between goroutines
	delegationChannel := make(chan *tzkt.DelegationItems)

	// Create a context that can be cancelled on stop signal
	ctx, cancel := context.WithCancel(context.Background())

	// Set up the Ctrl+C handler to gracefully stop the program
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Set the number of goroutines to use
	numGoroutines := 30

	// Launch goroutines to fetch and save delegations
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go saveDelegations(ctx, delegationChannel, &wg)
	}

	// Fetch the paginated items concurrently
	pageSize := 800
	flters := tzkt.Filters{}
	pagination := tzkt.Pagination{
		Limit: pageSize,
	}
	// Fetch the last saved id
	if lastID := getLastID(); lastID != nil {
		v := *lastID
		pagination.OffsetCr = int(v)
	}
fetchLoop:
	for {
		resp, err := cli.Delegations(&flters, &pagination)
		if err != nil {
			log.Printf("[ERROR] error making API request:%v\n", err)
			break
		}

		// Send items to the channel for goroutines to process and save
		delegationChannel <- resp

		// Check if there are more pages
		if !resp.HasMore || resp.Items == nil {
			break fetchLoop
		}
		pagination.OffsetCr = int(resp.Items[len(resp.Items)-1].ID)
		// Check if the program should stop
		select {
		case <-stop:
			log.Println("Received stop signal. Stopping...")
			cancel() // Cancel the context to signal goroutines to stop
			break fetchLoop
		default:
			// Continue fetching
		}
	}
	// Close the delegation channel to indicate that no more delegations will be sent
	close(delegationChannel)

	// Wait for all goroutines to finish
	wg.Wait()

}

func saveDelegations(ctx context.Context, delegationChannel <-chan *tzkt.DelegationItems, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			// Context cancelled, stop processing delegations
			return
		case delegation, ok := <-delegationChannel:
			if !ok {
				// Channel closed, no more delegations to process
				return
			}
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
			block_hashes := make([]string, len(delegation.Items))
			block_heights := make([]int, len(delegation.Items))

			for i, row := range delegation.Items {
				ids[i] = row.ID
				timestamps[i] = row.Timestamp
				amounts[i] = row.Amount
				delegators[i] = row.NewDelegate.Address
				block_hashes[i] = row.Block
				block_heights[i] = row.Level
			}
			rr, err := db.Get(dbURL).ExecContext(ctx, `
			INSERT INTO delegations
			(id, timestamp, amount, delegator, block_hash, block_level)
			(SELECT  * FROM UNNEST($1::bigint[], $2::timestamp[], $3::bigint[], $4::varchar[], $5::text[], $6::bigint[]))
			ON CONFLICT (id) DO UPDATE SET delegator = EXCLUDED.delegator
			`, ids, timestamps, amounts, delegators, block_hashes, block_heights)
			if err != nil {
				log.Printf("[ERROR] Failed to save : %s\n", err)
				return
			}
			cnt, _ := rr.RowsAffected()
			log.Printf("[INFO] %d delegations saved", cnt)
		}
	}
}
