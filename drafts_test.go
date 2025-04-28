package sleeper

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetDraftsForUser(t *testing.T) {
	// Create mock draft data
	mockDrafts := []Draft{
		{
			DraftID:  "123",
			LeagueID: "456",
			Sport:    "nfl",
			Season:   "2023",
			Status:   "complete",
			Settings: struct {
				AlphaSort             int `json:"alpha_sort"`
				AutopauseEnabled      int `json:"autopause_enabled"`
				AutopauseEndTime      int `json:"autopause_end_time"`
				AutopauseStartTime    int `json:"autopause_start_time"`
				Autostart             int `json:"autostart"`
				CPUAutopick           int `json:"cpu_autopick"`
				EnforcePositionLimits int `json:"enforce_position_limits"`
				NominationTimer       int `json:"nomination_timer"`
				PickTimer             int `json:"pick_timer"`
				PlayerType            int `json:"player_type"`
				ReversalRound         int `json:"reversal_round"`
				Rounds                int `json:"rounds"`
				SlotsBn               int `json:"slots_bn"`
				SlotsDef              int `json:"slots_def"`
				SlotsFlex             int `json:"slots_flex"`
				SlotsK                int `json:"slots_k"`
				SlotsQb               int `json:"slots_qb"`
				SlotsRb               int `json:"slots_rb"`
				SlotsTe               int `json:"slots_te"`
				SlotsWr               int `json:"slots_wr"`
				Teams                 int `json:"teams"`
			}{
				Rounds: 15,
				Teams:  12,
			},
		},
	}

	// Convert mock data to JSON
	mockData, err := json.Marshal(mockDrafts)
	if err != nil {
		t.Fatalf("Failed to marshal mock data: %v", err)
	}

	// Create test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/user/123/drafts/nfl/2023" {
			t.Errorf("Expected path /v1/user/123/drafts/nfl/2023, got %s", r.URL.Path)
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
	drafts, err := client.GetDraftsForUser("123", "nfl", 2023)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(drafts) != 1 {
		t.Errorf("Expected 1 draft, got %d", len(drafts))
	}
	if drafts[0].DraftID != "123" {
		t.Errorf("Expected draft ID '123', got '%s'", drafts[0].DraftID)
	}
	if drafts[0].Settings.Rounds != 15 {
		t.Errorf("Expected 15 rounds, got %d", drafts[0].Settings.Rounds)
	}
}

func TestGetDraftsForLeague(t *testing.T) {
	// Create mock draft data
	mockDrafts := []Draft{
		{
			DraftID:  "123",
			LeagueID: "456",
			Sport:    "nfl",
			Season:   "2023",
			Status:   "complete",
		},
	}

	// Convert mock data to JSON
	mockData, err := json.Marshal(mockDrafts)
	if err != nil {
		t.Fatalf("Failed to marshal mock data: %v", err)
	}

	// Create test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/league/456/drafts" {
			t.Errorf("Expected path /v1/league/456/drafts, got %s", r.URL.Path)
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
	drafts, err := client.GetDraftsForLeague("456")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(drafts) != 1 {
		t.Errorf("Expected 1 draft, got %d", len(drafts))
	}
	if drafts[0].LeagueID != "456" {
		t.Errorf("Expected league ID '456', got '%s'", drafts[0].LeagueID)
	}
}

func TestGetDraft(t *testing.T) {
	// Create mock draft data
	mockDraft := Draft{
		DraftID:  "123",
		LeagueID: "456",
		Sport:    "nfl",
		Season:   "2023",
		Status:   "complete",
		Metadata: struct {
			Description string `json:"description"`
			Name        string `json:"name"`
			ScoringType string `json:"scoring_type"`
		}{
			Name: "Test Draft",
		},
	}

	// Convert mock data to JSON
	mockData, err := json.Marshal(mockDraft)
	if err != nil {
		t.Fatalf("Failed to marshal mock data: %v", err)
	}

	// Create test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/draft/123" {
			t.Errorf("Expected path /v1/draft/123, got %s", r.URL.Path)
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
	draft, err := client.GetDraft("123")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if draft.DraftID != "123" {
		t.Errorf("Expected draft ID '123', got '%s'", draft.DraftID)
	}
	if draft.Metadata.Name != "Test Draft" {
		t.Errorf("Expected draft name 'Test Draft', got '%s'", draft.Metadata.Name)
	}
}

func TestGetAllDraftPicks(t *testing.T) {
	// Create mock draft picks data
	mockPicks := []DraftPlayer{
		{
			Round:    1,
			RosterID: 1,
			PlayerID: "player1",
			PickNo:   1,
			Metadata: struct {
				YearsExp     string `json:"years_exp"`
				Team         string `json:"team"`
				Status       string `json:"status"`
				Sport        string `json:"sport"`
				Position     string `json:"position"`
				PlayerID     string `json:"player_id"`
				Number       string `json:"number"`
				NewsUpdated  string `json:"news_updated"`
				LastName     string `json:"last_name"`
				InjuryStatus string `json:"injury_status"`
				FirstName    string `json:"first_name"`
			}{
				Position: "QB",
				Team:     "SF",
			},
		},
	}

	// Convert mock data to JSON
	mockData, err := json.Marshal(mockPicks)
	if err != nil {
		t.Fatalf("Failed to marshal mock data: %v", err)
	}

	// Create test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/draft/123/picks" {
			t.Errorf("Expected path /v1/draft/123/picks, got %s", r.URL.Path)
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
	picks, err := client.GetAllDraftPicks("123")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(picks) != 1 {
		t.Errorf("Expected 1 pick, got %d", len(picks))
	}
	if picks[0].Round != 1 {
		t.Errorf("Expected round 1, got %d", picks[0].Round)
	}
	if picks[0].Metadata.Position != "QB" {
		t.Errorf("Expected position 'QB', got '%s'", picks[0].Metadata.Position)
	}
}

func TestGetDraftTradedPicks(t *testing.T) {
	// Create mock traded picks data
	mockTradedPicks := []TradedPick{
		{
			OwnerID:         1,
			PreviousOwnerID: 2,
			Round:           1,
			RosterID:        1,
			Season:          "2023",
		},
	}

	// Convert mock data to JSON
	mockData, err := json.Marshal(mockTradedPicks)
	if err != nil {
		t.Fatalf("Failed to marshal mock data: %v", err)
	}

	// Create test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/draft/123/traded_picks" {
			t.Errorf("Expected path /v1/draft/123/traded_picks, got %s", r.URL.Path)
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
	picks, err := client.GetDraftTradedPicks("123")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(picks) != 1 {
		t.Errorf("Expected 1 traded pick, got %d", len(picks))
	}
	if picks[0].Round != 1 {
		t.Errorf("Expected round 1, got %d", picks[0].Round)
	}
	if picks[0].PreviousOwnerID != 2 {
		t.Errorf("Expected previous owner ID 2, got %d", picks[0].PreviousOwnerID)
	}
}

func TestGetDraftError(t *testing.T) {
	// Create test server that returns 404
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer ts.Close()

	// Create client with test server URL
	client := NewClientWithOptions(ClientOptions{
		BaseURL: ts.URL,
	})

	// Test draft not found
	_, err := client.GetDraft("nonexistent")
	if err == nil {
		t.Error("Expected error for nonexistent draft, got nil")
	}
}

func TestGetDraftInvalidResponse(t *testing.T) {
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
	_, err := client.GetDraft("123")
	if err == nil {
		t.Error("Expected error for invalid JSON, got nil")
	}
}
