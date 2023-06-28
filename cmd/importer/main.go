package main

import (
	"context"
	"database/sql"
	"delegationz/pkg/repository"
	"delegationz/pkg/services/tzkt"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func main() {
	dbURL := "postgres://postgres:supersecret@localhost:5432/dev"
	// Configure the PostgreSQL connection
	connConfig, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		log.Fatal("Error parsing connection config:", err)
	}

	// Create the database connection
	db, err := sql.Open("pgx", connConfig.ConnString())
	if err != nil {
		log.Fatal("Error opening database connection:", err)
	}
	defer db.Close()
	// Query delegations using SQLBoiler
	ctx := context.Background()
	delegations, err := repository.Delegations(qm.Limit(10)).All(ctx, db)
	if err != nil {
		log.Fatal("Error querying delegations:", err)
	}

	log.Printf("[DEBUG] %d delegations in database", len(delegations))

	// Process the retrieved delegation data
	for _, delegation := range delegations {
		fmt.Println("Delegation ID:", delegation.ID)
		fmt.Println("Timestamp:", delegation.Timestamp)
		fmt.Println("Amount:", delegation.Amount)
		fmt.Println("Delegator:", delegation.Delegator)
		fmt.Println("Block:", delegation.Block)
		fmt.Println("---------------------------")
	}
	tzktClient := tzkt.NewTzktClient()
	firstValDate := "2018-06-30T19:30:27Z"

	ts, _ := time.Parse(time.RFC3339, firstValDate)
	te := ts.AddDate(0, 0, 3)
	flters := &tzkt.Filters{
		TimestampGe: ts.Format(time.RFC3339),
		TimestampLt: te.Format(time.RFC3339),
		TypeEq:      "applied",
	}

	var val *tzkt.DelegationItems
	lmts := tzkt.Pagination{
		Offset: 0,
		Limit:  600,
	}
	for val == nil || (val != nil && val.HasMore) {
		rr, err := tzktClient.Delegations(flters, &lmts)
		if err != nil {
			log.Fatal(err)
		}

		// for _, el := range val.Items {
		// 	log.Printf("Address %v delegated %v Tzs to %s @ %+v\n", el.Sender.Address, el.Amount, el.NewDelegate.Address, el.Timestamp.Format(time.RFC3339))
		// }
		if val == nil {
			val = rr
		} else {
			val.Items = append(val.Items, rr.Items...)
			val.HasMore = rr.HasMore
		}
		lmts.Offset++
		time.Sleep(1 * time.Second)
		log.Printf("[DEBUG] %v delegations from API (more: %v)", len(val.Items), rr.HasMore)
	}
	log.Printf("[DEBUG] %v delegations from API (filters: %+v)", len(val.Items), flters)
}
