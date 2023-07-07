package importer

import (
	"delegationz/pkg/repository"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestSaveTopKindStats(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database connection: %v", err)
	}
	defer db.Close()

	testCases := []struct {
		kind    repository.StatKind
		year    string
		limit   int
		verbose bool
	}{
		{kind: repository.StatKindTOP10VALIDATORS, year: "2022", limit: 10, verbose: true},
		{kind: repository.StatKindTOP100VALIDATORS, year: "2022", limit: 100, verbose: true},
	}

	for _, tc := range testCases {
		mock.ExpectExec("INSERT INTO delegations_stats").
			WithArgs(tc.kind, tc.year, tc.limit).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err = SaveTopKindStats(db, tc.year, tc.kind, tc.verbose)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}

func TestSaveVolumeStats(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database connection: %v", err)
	}
	defer db.Close()

	testCases := []struct {
		kind      repository.StatKind
		year      string
		delegator string
		verbose   bool
	}{
		{kind: repository.StatKindDAILYVOLUME, year: "2022", delegator: "delegator1", verbose: true},
		{kind: repository.StatKindWEEKLYVOLUME, year: "2022", delegator: "delegator2", verbose: true},
		{kind: repository.StatKindMONTHLYVOLUME, year: "2022", delegator: "delegator3", verbose: true},
		{kind: repository.StatKindYEARLYVOLUME, year: "2022", delegator: "delegator4", verbose: true},
	}

	for _, tc := range testCases {
		mock.ExpectExec("INSERT INTO delegations_stats").
			WithArgs(tc.kind, tc.year, tc.delegator, 1).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err = SaveVolumeStats(db, tc.kind, tc.year, tc.delegator, tc.verbose)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}
