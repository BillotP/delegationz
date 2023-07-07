package statistics

import (
	"database/sql"
	"fmt"
	"time"
)

func UpdateTop(dbclient *sql.DB, frontier time.Time) error {
	q := `
	select delegator, sum(amount) as totals from delegations group by delegator order by totals desc limit 10;
	`
	fmt.Printf("%s\n", q)
	return nil
}
