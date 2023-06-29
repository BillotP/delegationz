package importer

import (
	"context"
	"database/sql"
	"delegationz/pkg/services/tzkt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func Run(db *sql.DB, pageSize int) {
	var wg sync.WaitGroup

	cli := tzkt.NewTzktClient()

	delegationChannel := make(chan *tzkt.DelegationItems)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// TODO : Check this magic numbers
	numGoroutines := 5

	for i := 0; i < numGoroutines; i++ {
		go saveDelegations(db, ctx, delegationChannel)
	}

	flters := tzkt.Filters{}
	pagination := tzkt.Pagination{
		Limit: pageSize,
	}

	if lastID := getLastID(db); lastID != nil {
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
		log.Printf("[INFO] Fetched %d items", len(resp.Items))
		delegationChannel <- resp

		if !resp.HasMore || resp.Items == nil {
			break fetchLoop
		}
		pagination.OffsetCr = int(resp.Items[len(resp.Items)-1].ID)

		select {
		case <-stop:
			log.Println("[INFO] Received stop signal. Stopping...")
			cancel()
			break fetchLoop
		default:

		}
	}

	close(delegationChannel)

	wg.Wait()
}

func getLastID(db *sql.DB) *int64 {
	var lID int64
	err := db.QueryRow(`select MAX(id) from delegations`).Scan(&lID)
	if err == sql.ErrNoRows {
		return nil
	}
	return &lID
}

func saveDelegations(db *sql.DB, ctx context.Context, delegationChannel <-chan *tzkt.DelegationItems) {

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
			go func() {
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

				rr, err := db.ExecContext(ctx, `
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
			}()
		}

	}
}
