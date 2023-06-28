package main

import (
	"delegationz/pkg/services/tzkt"
	"log"
	"time"
)

func main() {
	tzktClient := tzkt.NewTzktClient()
	twoDaysFromNow := time.Now().AddDate(0, 0, -5).Format(time.RFC3339)
	flters := &tzkt.Filters{
		TimestampGe: twoDaysFromNow,
	}
	val, err := tzktClient.Delegations(flters, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Found %v delegations (more: %v)", len(val.Items), val.HasMore)
	// for _, el := range val.Items {
	// 	log.Printf("Address %v delegated %v Tzs to %s @ %+v\n", el.Sender.Address, el.Amount, el.NewDelegate.Address, el.Timestamp.Format(time.RFC3339))
	// }
}
