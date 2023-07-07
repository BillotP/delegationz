package tzkt

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTzktClient(t *testing.T) {
	// Call the function
	client := NewTzktClient()

	// Assert the client struct fields
	assert.NotNil(t, client.cli)
	assert.Equal(t, 1, client.loglevel)
	assert.Equal(t, "https://api.tzkt.io/v1", client.baseURL)
	assert.IsType(t, &http.Client{}, client.cli)
}

func TestDo(t *testing.T) {
	// Create a test server to mock the HTTP requests
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the request method and URL
		if r.Method != http.MethodGet {
			t.Errorf("expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/api/endpoint" {
			t.Errorf("expected endpoint /api/endpoint, got %s", r.URL.Path)
		}

		// Write a mock response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"result": "success"}`))
	}))
	defer server.Close()

	// Create a new instance of TzktClient
	client := &TzktClient{
		cli:     server.Client(),
		baseURL: server.URL,
	}

	// Make a test request
	err := client.do(http.MethodGet, "api/endpoint", nil)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestDelegations(t *testing.T) {
	// Create a test server to mock the HTTP requests
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the request method and URL
		if r.Method != http.MethodGet {
			t.Errorf("expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/operations/delegations" {
			t.Errorf("expected endpoint /operations/delegations, got %s", r.URL.Path)
		}

		// Check the query parameters
		expectedQuery := "limit=10&timestamp.gt=2022-01-01&timestamp.lt=2022-12-31"
		if r.URL.RawQuery != expectedQuery {
			t.Errorf("expected query parameters %s, got %s", expectedQuery, r.URL.RawQuery)
		}

		// Write a mock response
		bb, _ := json.Marshal([]*DelegationItem{})
		w.WriteHeader(http.StatusOK)
		w.Write(bb)
	}))
	defer server.Close()

	// Create a new instance of TzktClient
	client := &TzktClient{
		cli:     server.Client(),
		baseURL: server.URL,
	}

	// Create filters and pagination objects
	filters := NewFilters(
		WithTimestampGt("2022-01-01"),
		WithTimestampLt("2022-12-31"),
	)
	pagination := NewPagination(
		WithLimit(10),
		WithOffsetPage(0),
	)

	// Make a test request to the Delegations method
	_, err := client.Delegations(filters, pagination)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
}
