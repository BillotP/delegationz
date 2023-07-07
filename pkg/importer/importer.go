package importer

import (
	"context"
	"database/sql"
	"delegationz/pkg/tzkt"
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
			break
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
	datas := toBulkInsert(delegation)
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("[ERROR] Failed to begin db transaction : %s\n", err)
		return err
	}
	rr, err := tx.ExecContext(ctx, `
		INSERT INTO delegators (address, first_seen) 
		(SELECT * FROM UNNEST($1::varchar[], $2::timestamp[]))
		ON CONFLICT (address) DO NOTHING;`, datas.delegators, datas.timestamps)
	if err != nil {
		log.Printf("[ERROR] Failed to save : %s\n", err)
		return err
	}
	cnt, _ := rr.RowsAffected()
	if verbose {
		log.Printf("[INFO] %d delegators saved", cnt)
	}
	rr, err = tx.ExecContext(ctx, `
		INSERT INTO delegations
		(id, timestamp, amount, delegator, block_hash, block_level)
		(SELECT  * FROM UNNEST($1::bigint[], $2::timestamp[], $3::bigint[], $4::varchar[], $5::text[], $6::bigint[]))
		ON CONFLICT (id) DO UPDATE SET delegator = EXCLUDED.delegator, amount = EXCLUDED.amount;
		`, datas.ids, datas.timestamps, datas.amounts, datas.delegators, datas.block_hashes, datas.block_heights)
	if err != nil {
		log.Printf("[ERROR] Failed to save : %s\n", err)
		return err
	}
	cnt, _ = rr.RowsAffected()
	if verbose {
		log.Printf("[INFO] %d delegations saved", cnt)
	}
	if err = tx.Commit(); err != nil {
		log.Printf("[ERROR] Failed to commit tx : %s\n", err)
		return err
	}
	return nil
}
