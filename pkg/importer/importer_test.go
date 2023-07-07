package importer

import (
	"delegationz/pkg/tzkt"
	"os"
	"os/signal"
	"regexp"
	"syscall"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

type MockTzktClient struct{}

func (m *MockTzktClient) Delegations(filters *tzkt.Filters, pagination *tzkt.Pagination) (*tzkt.DelegationItems, error) {
	return &tzkt.DelegationItems{
		Items:   []*tzkt.DelegationItem{{ID: 1}},
		HasMore: false,
	}, nil
}
func TestRun(t *testing.T) {
	// Create a mock database using go-sqlmock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	// Create a mock TzktClient
	mockTzktClient := &MockTzktClient{}

	// Prepare the expected database mock responses
	rows := sqlmock.NewRows([]string{"MAX(id)"}).AddRow(1)
	mock.ExpectQuery(regexp.QuoteMeta("select MAX(id) from delegations")).WillReturnRows(rows)

	// Set up the Ctrl+C signal handling
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		time.Sleep(2 * time.Second) // Wait for the Run function to start
		stop <- os.Interrupt        // Send the Ctrl+C signal
	}()

	// Load importer
	importr := New(db, mockTzktClient,
		WithPageSize(10), WithWatch(false),
		WithReset(false), WithVerbose(true))

	// Call the Run function
	err = importr.Run()

	// Perform assertions or checks based on the expected behavior
	if err != nil {
		t.Errorf("Run returned an unexpected error: %v", err)
	}

	// Ensure all the expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Failed to meet database expectations: %v", err)
	}
}
