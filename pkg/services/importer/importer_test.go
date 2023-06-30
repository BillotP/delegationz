package importer

import (
	"context"
	"database/sql"
	"delegationz/pkg/services/tzkt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSaveDelegations_ContextCancellation(t *testing.T) {
	// Create a mock database connection
	db := &sql.DB{}

	// Create a cancellation context with a cancel function
	ctx, cancel := context.WithCancel(context.Background())

	// Create a delegation channel
	delegationChannel := make(chan *tzkt.DelegationItems)

	// Set up a flag to track if the sleep instruction is reached
	sleepReached := false

	// Start the saveDelegations goroutine
	go func() {
		saveDelegations(db, ctx, delegationChannel, true)
		sleepReached = true
	}()

	// Cancel the context after a short delay
	go func() {
		time.Sleep(100 * time.Millisecond)
		cancel()
	}()

	// Wait for a short duration to allow time for the goroutine to exit
	time.Sleep(300 * time.Millisecond)

	// Assert that the sleep instruction was reached and the goroutine exited
	assert.True(t, sleepReached)
}
