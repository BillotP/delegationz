package importer

import (
	"context"
	"database/sql"
	"delegationz/pkg/repository"
	"log"
)

func SaveTopKindStats(db *sql.DB, year string, kind repository.StatKind, verbose bool) error {
	limit := 1
	q := `
	INSERT INTO delegations_stats (id, kind, value, timestamp)
	SELECT
	 NEXTVAL('delegations_stats_id_seq'),
	  $1 AS kind,
	  json_agg(json_build_object('delegator', delegator, 'value', totals)) AS value,
	  -- max(lastdate) AS timestamp
	  DATE_TRUNC('year', MAKE_DATE($2::integer, 1, 1)) + INTERVAL '1 year - 1 day' AS timestamp
	FROM (
	  SELECT delegator, SUM(amount) AS totals
	  FROM delegations
	  WHERE EXTRACT(year FROM timestamp) = $2::bigint
	  GROUP BY delegator
	  ORDER BY totals DESC
	  LIMIT $3
	) bulk;`
	switch kind {
	case repository.StatKindTOP10VALIDATORS:
		limit = 10
	case repository.StatKindTOP100VALIDATORS:
		limit = 100
	default:
		return nil
	}
	rr, err := db.ExecContext(context.Background(), q, kind, year, limit)
	if err != nil {
		log.Printf("[ERROR] failed to save stats : %s", err)
		return err
	}
	cnt, _ := rr.RowsAffected()
	if verbose {
		log.Printf("[INFO] Inserted %d stat row for %s Y%s", cnt, kind, year)
	}
	return nil
}

func SaveVolumeStats(db *sql.DB, kind repository.StatKind, year, delegator string, verbose bool) error {
	switch kind {
	case repository.StatKindDAILYVOLUME:
		return nil
	case repository.StatKindWEEKLYVOLUME:
		return nil
	case repository.StatKindMONTHLYVOLUME:
		return nil
	case repository.StatKindYEARLYVOLUME:
		// Should check if delegator exist !
		q := `
		INSERT INTO delegations_stats (id, kind, value, timestamp, delegator_address)
		SELECT
			NEXTVAL('delegations_stats_id_seq') as id,
			$1 AS kind,
			json_build_object('delegator', delegator, 'value', SUM(amount)) AS value,
			DATE_TRUNC('year', MAKE_DATE($2::integer, 1, 1)) + INTERVAL '1 year - 1 day' AS timestamp,
			delegator AS delegator_address
	    FROM 
			delegations
		WHERE EXTRACT(year FROM timestamp) = $2
		AND 
			delegator = $3
		GROUP BY 
			delegator
		ORDER BY 
			SUM(amount) DESC
		LIMIT $4
		`
		rr, err := db.ExecContext(context.Background(), q, kind, year, delegator, 1)
		if err != nil {
			log.Printf("[ERROR] failed to save stats : %s", err)
			return err
		}
		cnt, _ := rr.RowsAffected()
		if verbose {
			log.Printf("[INFO] Inserted %d stat row for %s Y%s %s", cnt, kind, year, delegator)
		}
	}
	return nil
}
