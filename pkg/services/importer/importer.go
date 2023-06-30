package importer

import (
	"context"
	"database/sql"
	"delegationz/pkg/services/tzkt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(db *sql.DB, pageSize int, watch, fromstart, verbose bool) {

	cli := tzkt.NewTzktClient()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	flters := tzkt.Filters{}
	pagination := tzkt.Pagination{
		Limit: pageSize,
	}

	if lastID := getLastID(db); lastID != nil {
		v := *lastID
		if !fromstart {
			pagination.OffsetCr = int(v)
		}
	} else {
		log.Printf("[INFO] No existing index, starting (re)sync")
	}
fetchLoop:
	for {
		resp, err := cli.Delegations(&flters, &pagination)
		if err != nil {
			log.Printf("[ERROR] error making API request:%v\n", err)
			break
		}
		if verbose {
			log.Printf("[INFO] Fetched %d items", len(resp.Items))
		}

		if !watch && (!resp.HasMore || resp.Items == nil) {
			// If no more items needs to be fetched and importer is not in watch mode, exiting
			break fetchLoop
		} else if !resp.HasMore {
			if verbose {
				log.Printf("[INFO] Last events saved, waiting for update polling\n")
			}
			// Arbitrary delay between each call
			time.Sleep(2 * time.Second)
		}
		if len(resp.Items) > 0 {
			saveDelegations(db, ctx, resp, verbose)
			pagination.OffsetCr = int(resp.Items[len(resp.Items)-1].ID)
		} else {
			pagination.OffsetCr = int(*getLastID(db))
		}

		select {
		case <-stop:
			cancel()
			break fetchLoop
		default:

		}
	}
	log.Println("[INFO] Received stop signal. Exiting...")
}

func getLastID(db *sql.DB) *int64 {
	var lID int64
	err := db.QueryRow(`select MAX(id) from delegations`).Scan(&lID)
	if err == sql.ErrNoRows {
		return nil
	}
	return &lID
}

func saveDelegations(db *sql.DB, ctx context.Context, delegation *tzkt.DelegationItems, verbose bool) error {

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

	rr, err := db.ExecContext(ctx, `
		INSERT INTO delegations
		(id, timestamp, amount, delegator, block_hash, block_level)
		(SELECT  * FROM UNNEST($1::bigint[], $2::timestamp[], $3::bigint[], $4::varchar[], $5::text[], $6::bigint[]))
		ON CONFLICT (id) DO UPDATE SET delegator = EXCLUDED.delegator
		`, ids, timestamps, amounts, delegators, block_hashes, block_heights)
	if err != nil {
		log.Printf("[ERROR] Failed to save : %s\n", err)
		return err
	}
	cnt, _ := rr.RowsAffected()
	if verbose {
		log.Printf("[INFO] %d delegations saved", cnt)
	}
	return nil
}
