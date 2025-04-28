package sleeper

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllLeagesForUser(t *testing.T) {
	// Create mock league data
	mockLeagues := []League{
		{
			LeagueID: "123",
			Name:     "Test League 1",
			Sport:    "nfl",
			Season:   "2023",
		},
		{
			LeagueID: "456",
			Name:     "Test League 2",
			Sport:    "nfl",
			Season:   "2023",
		},
	}

	// Convert mock data to JSON
	mockData, err := json.Marshal(mockLeagues)
	if err != nil {
		t.Fatalf("Failed to marshal mock data: %v", err)
	}

	// Create test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/user/123/leagues/nfl/2023" {
			t.Errorf("Expected path /v1/user/123/leagues/nfl/2023, got %s", r.URL.Path)
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
	leagues, err := client.GetAllLeagesForUser("123", "nfl", "2023")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(leagues) != 2 {
		t.Errorf("Expected 2 leagues, got %d", len(leagues))
	}
	if leagues[0].LeagueID != "123" {
		t.Errorf("Expected league ID '123', got '%s'", leagues[0].LeagueID)
	}
}

func TestGetLeague(t *testing.T) {
	// Create mock league data
	mockLeague := League{
		LeagueID: "123",
		Name:     "Test League",
		Sport:    "nfl",
		Season:   "2023",
		Settings: struct {
			BenchLock                int `json:"bench_lock"`
			CapacityOverride         int `json:"capacity_override"`
			CommissionerDirectInvite int `json:"commissioner_direct_invite"`
			DailyWaivers             int `json:"daily_waivers"`
			DailyWaiversHour         int `json:"daily_waivers_hour"`
			DisableAdds              int `json:"disable_adds"`
			DraftRounds              int `json:"draft_rounds"`
			LeagueAverageMatch       int `json:"league_average_match"`
			Leg                      int `json:"leg"`
			MaxKeepers               int `json:"max_keepers"`
			NumTeams                 int `json:"num_teams"`
			OffseasonAdds            int `json:"offseason_adds"`
			PickTrading              int `json:"pick_trading"`
			PlayoffRoundType         int `json:"playoff_round_type"`
			PlayoffSeedType          int `json:"playoff_seed_type"`
			PlayoffTeams             int `json:"playoff_teams"`
			PlayoffType              int `json:"playoff_type"`
			PlayoffWeekStart         int `json:"playoff_week_start"`
			ReserveAllowCov          int `json:"reserve_allow_cov"`
			ReserveAllowDnr          int `json:"reserve_allow_dnr"`
			ReserveAllowDoubtful     int `json:"reserve_allow_doubtful"`
			ReserveAllowNa           int `json:"reserve_allow_na"`
			ReserveAllowOut          int `json:"reserve_allow_out"`
			ReserveAllowSus          int `json:"reserve_allow_sus"`
			ReserveSlots             int `json:"reserve_slots"`
			TaxiAllowVets            int `json:"taxi_allow_vets"`
			TaxiDeadline             int `json:"taxi_deadline"`
			TaxiSlots                int `json:"taxi_slots"`
			TaxiYears                int `json:"taxi_years"`
			TradeDeadline            int `json:"trade_deadline"`
			TradeReviewDays          int `json:"trade_review_days"`
			Type                     int `json:"type"`
			WaiverBidMin             int `json:"waiver_bid_min"`
			WaiverBudget             int `json:"waiver_budget"`
			WaiverClearDays          int `json:"waiver_clear_days"`
			WaiverDayOfWeek          int `json:"waiver_day_of_week"`
			WaiverType               int `json:"waiver_type"`
		}{
			NumTeams: 10,
		},
	}

	// Convert mock data to JSON
	mockData, err := json.Marshal(mockLeague)
	if err != nil {
		t.Fatalf("Failed to marshal mock data: %v", err)
	}

	// Create test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/league/123" {
			t.Errorf("Expected path /v1/league/123, got %s", r.URL.Path)
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
	league, err := client.GetLeague("123")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if league.LeagueID != "123" {
		t.Errorf("Expected league ID '123', got '%s'", league.LeagueID)
	}
	if league.Settings.NumTeams != 10 {
		t.Errorf("Expected 10 teams, got %d", league.Settings.NumTeams)
	}
}

func TestGetRosters(t *testing.T) {
	// Create mock roster data
	mockRosters := []Roster{
		{
			RosterID: 1,
			OwnerID:  "123",
			Players:  []string{"player1", "player2"},
			Settings: struct {
				Fpts             int `json:"fpts"`
				Losses           int `json:"losses"`
				Ties             int `json:"ties"`
				TotalMoves       int `json:"total_moves"`
				WaiverBudgetUsed int `json:"waiver_budget_used"`
				WaiverPosition   int `json:"waiver_position"`
				Wins             int `json:"wins"`
			}{
				Wins:   5,
				Losses: 2,
			},
		},
	}

	// Convert mock data to JSON
	mockData, err := json.Marshal(mockRosters)
	if err != nil {
		t.Fatalf("Failed to marshal mock data: %v", err)
	}

	// Create test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/league/123/rosters" {
			t.Errorf("Expected path /v1/league/123/rosters, got %s", r.URL.Path)
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
	rosters, err := client.GetRosters("123")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(rosters) != 1 {
		t.Errorf("Expected 1 roster, got %d", len(rosters))
	}
	if rosters[0].RosterID != 1 {
		t.Errorf("Expected roster ID 1, got %d", rosters[0].RosterID)
	}
}

func TestGetLeagueUsers(t *testing.T) {
	// Create mock league users data
	mockUsers := []LeagueUser{
		{
			UserID:      "123",
			DisplayName: "Test User",
			Avatar:      "https://example.com/avatar.jpg",
			IsBot:       false,
		},
	}

	// Convert mock data to JSON
	mockData, err := json.Marshal(mockUsers)
	if err != nil {
		t.Fatalf("Failed to marshal mock data: %v", err)
	}

	// Create test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/league/123/users" {
			t.Errorf("Expected path /v1/league/123/users, got %s", r.URL.Path)
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
	users, err := client.GetLeagueUsers("123")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(users) != 1 {
		t.Errorf("Expected 1 user, got %d", len(users))
	}
	if users[0].UserID != "123" {
		t.Errorf("Expected user ID '123', got '%s'", users[0].UserID)
	}
}

func TestGetMatchups(t *testing.T) {
	// Create mock matchup data
	mockMatchups := []Matchup{
		{
			MatchupID: 1,
			RosterID:  1,
			Points:    100.5,
			Starters:  []string{"player1", "player2"},
		},
	}

	// Convert mock data to JSON
	mockData, err := json.Marshal(mockMatchups)
	if err != nil {
		t.Fatalf("Failed to marshal mock data: %v", err)
	}

	// Create test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/league/123/matchups/1" {
			t.Errorf("Expected path /v1/league/123/matchups/1, got %s", r.URL.Path)
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
	matchups, err := client.GetMatchups("123", 1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(matchups) != 1 {
		t.Errorf("Expected 1 matchup, got %d", len(matchups))
	}
	if matchups[0].MatchupID != 1 {
		t.Errorf("Expected matchup ID 1, got %d", matchups[0].MatchupID)
	}
}

func TestGetSportState(t *testing.T) {
	// Create mock sport state data
	mockState := SportState{
		DisplayWeek: 1,
		Season:      "2023",
		SeasonType:  "regular",
		Week:        1,
	}

	// Convert mock data to JSON
	mockData, err := json.Marshal(mockState)
	if err != nil {
		t.Fatalf("Failed to marshal mock data: %v", err)
	}

	// Create test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/state/nfl" {
			t.Errorf("Expected path /v1/state/nfl, got %s", r.URL.Path)
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
	state, err := client.GetSportState("nfl")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if state.DisplayWeek != 1 {
		t.Errorf("Expected display week 1, got %d", state.DisplayWeek)
	}
	if state.Season != "2023" {
		t.Errorf("Expected season '2023', got '%s'", state.Season)
	}
}

func TestGetLeagueError(t *testing.T) {
	// Create test server that returns 404
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer ts.Close()

	// Create client with test server URL
	client := NewClientWithOptions(ClientOptions{
		BaseURL: ts.URL,
	})

	// Test league not found
	_, err := client.GetLeague("nonexistent")
	if err == nil {
		t.Error("Expected error for nonexistent league, got nil")
	}
}

func TestGetLeagueInvalidResponse(t *testing.T) {
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
	_, err := client.GetLeague("123")
	if err == nil {
		t.Error("Expected error for invalid JSON, got nil")
	}
}
