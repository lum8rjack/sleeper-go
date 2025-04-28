package sleeper

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetNflScheduleRegularSeason(t *testing.T) {
	// Create mock schedule data
	mockSchedule := NflSchedule{
		{
			Status: "scheduled",
			Date:   "2023-09-10",
			Home:   "SF",
			Week:   1,
			GameID: "game1",
			Away:   "PIT",
		},
		{
			Status: "scheduled",
			Date:   "2023-09-11",
			Home:   "NYJ",
			Week:   1,
			GameID: "game2",
			Away:   "BUF",
		},
	}

	// Convert mock data to JSON
	mockData, err := json.Marshal(mockSchedule)
	if err != nil {
		t.Fatalf("Failed to marshal mock data: %v", err)
	}

	// Create test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/schedule/nfl/regular/2023" {
			t.Errorf("Expected path /schedule/nfl/regular/2023, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mockData)
	}))
	defer ts.Close()

	// Create client with test server URL
	client := NewClientWithOptions(ClientOptions{
		BaseURL: ts.URL,
	})

	// Test successful request
	schedule, err := client.GetNflSchedule(2023, false)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify schedule data
	if len(schedule) != 2 {
		t.Errorf("Expected 2 games, got %d", len(schedule))
	}
	if schedule[0].Home != "SF" {
		t.Errorf("Expected home team 'SF', got '%s'", schedule[0].Home)
	}
	if schedule[0].Away != "PIT" {
		t.Errorf("Expected away team 'PIT', got '%s'", schedule[0].Away)
	}
	if schedule[0].Week != 1 {
		t.Errorf("Expected week 1, got %d", schedule[0].Week)
	}
}

func TestGetNflSchedulePostseason(t *testing.T) {
	// Create mock schedule data
	mockSchedule := NflSchedule{
		{
			Status: "scheduled",
			Date:   "2023-01-14",
			Home:   "KC",
			Week:   1,
			GameID: "game1",
			Away:   "MIA",
		},
		{
			Status: "scheduled",
			Date:   "2023-01-15",
			Home:   "BUF",
			Week:   1,
			GameID: "game2",
			Away:   "PIT",
		},
	}

	// Convert mock data to JSON
	mockData, err := json.Marshal(mockSchedule)
	if err != nil {
		t.Fatalf("Failed to marshal mock data: %v", err)
	}

	// Create test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/schedule/nfl/post/2023" {
			t.Errorf("Expected path /schedule/nfl/post/2023, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mockData)
	}))
	defer ts.Close()

	// Create client with test server URL
	client := NewClientWithOptions(ClientOptions{
		BaseURL: ts.URL,
	})

	// Test successful request
	schedule, err := client.GetNflSchedule(2023, true)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify schedule data
	if len(schedule) != 2 {
		t.Errorf("Expected 2 games, got %d", len(schedule))
	}
	if schedule[0].Home != "KC" {
		t.Errorf("Expected home team 'KC', got '%s'", schedule[0].Home)
	}
	if schedule[0].Away != "MIA" {
		t.Errorf("Expected away team 'MIA', got '%s'", schedule[0].Away)
	}
}

func TestGetNflScheduleInvalidYear(t *testing.T) {
	// Create client
	client := NewClient()

	// Test year before 2009
	_, err := client.GetNflSchedule(2008, false)
	if err == nil {
		t.Error("Expected error for year before 2009, got nil")
	}

	// Test year in future
	futureYear := time.Now().Year() + 1
	_, err = client.GetNflSchedule(futureYear, false)
	if err == nil {
		t.Error("Expected error for future year, got nil")
	}
}

func TestGetNflScheduleError(t *testing.T) {
	// Create test server that returns 404
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer ts.Close()

	// Create client with test server URL
	client := NewClientWithOptions(ClientOptions{
		BaseURL: ts.URL,
	})

	// Test error response
	_, err := client.GetNflSchedule(2023, false)
	if err == nil {
		t.Error("Expected error for 404 response, got nil")
	}
}

func TestGetNflScheduleInvalidResponse(t *testing.T) {
	// Create test server that returns invalid JSON
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("invalid json"))
	}))
	defer ts.Close()

	// Create client with test server URL
	client := NewClientWithOptions(ClientOptions{
		BaseURL: ts.URL,
	})

	// Test invalid JSON response
	_, err := client.GetNflSchedule(2023, false)
	if err == nil {
		t.Error("Expected error for invalid JSON, got nil")
	}
}

func TestGetNflScheduleEmptyResponse(t *testing.T) {
	// Create test server that returns empty schedule
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("[]"))
	}))
	defer ts.Close()

	// Create client with test server URL
	client := NewClientWithOptions(ClientOptions{
		BaseURL: ts.URL,
	})

	// Test empty response
	schedule, err := client.GetNflSchedule(2023, false)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify empty schedule
	if len(schedule) != 0 {
		t.Errorf("Expected empty schedule, got %d games", len(schedule))
	}
}
