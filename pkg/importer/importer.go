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

type Importer struct {
	PageSize int
	Watch    bool
	Reset    bool
	Verbose  bool
	db       *sql.DB
	tzktcli  tzkt.ITzktClient
}

func New(db *sql.DB, tzktcli tzkt.ITzktClient, options ...func(*Importer)) *Importer {
	importr := &Importer{
		db:      db,
		tzktcli: tzktcli,
	}
	for _, o := range options {
		o(importr)
	}
	return importr
}

func WithPageSize(pagesize int) func(*Importer) {
	return func(i *Importer) {
		i.PageSize = pagesize
	}
}

func WithWatch(watch bool) func(*Importer) {
	return func(i *Importer) {
		i.Watch = watch
	}
}

func WithReset(reset bool) func(*Importer) {
	return func(i *Importer) {
		i.Reset = reset
	}
}

func WithVerbose(verbose bool) func(*Importer) {
	return func(i *Importer) {
		i.Verbose = verbose
	}
}

func (i *Importer) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	flters := tzkt.NewFilters()
	pagination := tzkt.NewPagination()
	if lastID := getLastID(i.db); lastID != nil {
		v := *lastID
		if !i.Reset {
			pagination = tzkt.NewPagination(tzkt.WithOffsetCursor(int(v)))
		}
	} else if i.Verbose {
		log.Printf("[INFO] No existing index, starting (re)sync")
	}
	for {
		select {
		default:
			resp, err := i.tzktcli.Delegations(flters, pagination)
			if err != nil {
				log.Printf("[ERROR] error making API request:%v\n", err)
				return err
			}
			if i.Verbose {
				log.Printf("[INFO] Fetched %d items", len(resp.Items))
			}

			if !i.Watch && (!resp.HasMore || resp.Items == nil) {
				// If no more items needs to be fetched and importer is not in watch mode, exiting
				return nil
			} else if !resp.HasMore {
				if i.Verbose {
					log.Printf("[INFO] Last events saved, waiting for update polling\n")
				}
				// Arbitrary delay between each call
				time.Sleep(2 * time.Second)
			}
			if len(resp.Items) > 0 {
				if err = saveDelegations(i.db, ctx, resp, i.Verbose); err != nil {
					return err
				}
				cr := int(resp.Items[len(resp.Items)-1].ID)
				pagination = tzkt.NewPagination(tzkt.WithOffsetCursor(cr))
			} else {
				cr := getLastID(i.db)
				pagination.OffsetCr = int(*cr)
			}

		case <-stop:
			log.Println("[INFO] Received stop signal. Exiting...")
			cancel()
			return nil
		}
	}
}

func getLastID(db *sql.DB) *int64 {
	var lID int64
	err := db.QueryRow("select MAX(id) from delegations").Scan(&lID)
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
