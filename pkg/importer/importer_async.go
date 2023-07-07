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

func RunWithDBWorkers(db *sql.DB, pageSize int, watch, fromstart, verbose bool) {

	cli := tzkt.NewTzktClient()

	delegationChannel := make(chan *tzkt.DelegationItems)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// TODO : Check this magic number (maybe optimise with )
	numGoroutines := 5

	for i := 0; i < numGoroutines; i++ {
		go saveDelegationsAsync(db, ctx, delegationChannel, verbose)
	}
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
			// Arbitrary delay between each call once synced
			time.Sleep(2 * time.Second)
		}
		if len(resp.Items) > 0 {
			delegationChannel <- resp
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
	close(delegationChannel)
	log.Println("[INFO] Received stop signal. Exiting...")
}

func saveDelegationsAsync(db *sql.DB, ctx context.Context, delegationChannel <-chan *tzkt.DelegationItems, verbose bool) {

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
			go saveDelegations(db, ctx, delegation, verbose)
		}

	}
}
