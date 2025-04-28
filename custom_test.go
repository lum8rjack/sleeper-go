package sleeper

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTeamMatchups(t *testing.T) {
	// Create mock data for getFantasyInfo
	mockUsers := []LeagueUser{
		{
			UserID:      "1",
			DisplayName: "User 1",
			Metadata: struct {
				AllowPn                 string `json:"allow_pn"`
				AllowSms                string `json:"allow_sms"`
				Avatar                  string `json:"avatar"`
				MascotMessage           string `json:"mascot_message"`
				MentionPn               string `json:"mention_pn"`
				PlayerLikePn            string `json:"player_like_pn"`
				PlayerNicknameUpdate    string `json:"player_nickname_update"`
				TeamName                string `json:"team_name"`
				TeamNameUpdate          string `json:"team_name_update"`
				TradeBlockPn            string `json:"trade_block_pn"`
				TransactionCommissioner string `json:"transaction_commissioner"`
				TransactionFreeAgent    string `json:"transaction_free_agent"`
				TransactionTrade        string `json:"transaction_trade"`
				TransactionWaiver       string `json:"transaction_waiver"`
				UserMessagePn           string `json:"user_message_pn"`
			}{
				TeamName: "Team 1",
			},
		},
		{
			UserID:      "2",
			DisplayName: "User 2",
			Metadata: struct {
				AllowPn                 string `json:"allow_pn"`
				AllowSms                string `json:"allow_sms"`
				Avatar                  string `json:"avatar"`
				MascotMessage           string `json:"mascot_message"`
				MentionPn               string `json:"mention_pn"`
				PlayerLikePn            string `json:"player_like_pn"`
				PlayerNicknameUpdate    string `json:"player_nickname_update"`
				TeamName                string `json:"team_name"`
				TeamNameUpdate          string `json:"team_name_update"`
				TradeBlockPn            string `json:"trade_block_pn"`
				TransactionCommissioner string `json:"transaction_commissioner"`
				TransactionFreeAgent    string `json:"transaction_free_agent"`
				TransactionTrade        string `json:"transaction_trade"`
				TransactionWaiver       string `json:"transaction_waiver"`
				UserMessagePn           string `json:"user_message_pn"`
			}{
				TeamName: "Team 2",
			},
		},
	}

	mockRosters := []Roster{
		{
			RosterID: 1,
			OwnerID:  "1",
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
		{
			RosterID: 2,
			OwnerID:  "2",
			Settings: struct {
				Fpts             int `json:"fpts"`
				Losses           int `json:"losses"`
				Ties             int `json:"ties"`
				TotalMoves       int `json:"total_moves"`
				WaiverBudgetUsed int `json:"waiver_budget_used"`
				WaiverPosition   int `json:"waiver_position"`
				Wins             int `json:"wins"`
			}{
				Wins:   3,
				Losses: 4,
			},
		},
	}

	mockMatchups := []Matchup{
		{
			MatchupID: 1,
			RosterID:  1,
			Points:    100.5,
		},
		{
			MatchupID: 1,
			RosterID:  2,
			Points:    90.5,
		},
	}

	// Create test server for league endpoint
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/v1/league/123":
			league := League{
				Sport: "nfl",
			}
			data, _ := json.Marshal(league)
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		case "/v1/state/nfl":
			state := SportState{
				Week:        1,
				SeasonType:  "regular",
				DisplayWeek: 1,
			}
			data, _ := json.Marshal(state)
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		case "/v1/league/123/users":
			data, _ := json.Marshal(mockUsers)
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		case "/v1/league/123/rosters":
			data, _ := json.Marshal(mockRosters)
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		case "/v1/league/123/matchups/1":
			data, _ := json.Marshal(mockMatchups)
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	// Create client with test server URL
	client := NewClientWithOptions(ClientOptions{
		BaseURL: ts.URL,
	})

	// Test successful request
	matchups, err := client.GetTeamMatchups("123", 1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(matchups) != 1 {
		t.Errorf("Expected 1 matchup, got %d", len(matchups))
	}
	if matchups[0].Teamname1 != "Team 1" {
		t.Errorf("Expected team 1 name 'Team 1', got '%s'", matchups[0].Teamname1)
	}
	if matchups[0].Teamname2 != "Team 2" {
		t.Errorf("Expected team 2 name 'Team 2', got '%s'", matchups[0].Teamname2)
	}
	if matchups[0].Team1Wins != 5 {
		t.Errorf("Expected team 1 wins 5, got %d", matchups[0].Team1Wins)
	}
	if matchups[0].Team2Wins != 3 {
		t.Errorf("Expected team 2 wins 3, got %d", matchups[0].Team2Wins)
	}
}

func TestGetScoreboards(t *testing.T) {
	// Create mock data for getFantasyInfo
	mockUsers := []LeagueUser{
		{
			UserID:      "1",
			DisplayName: "User 1",
			Metadata: struct {
				AllowPn                 string `json:"allow_pn"`
				AllowSms                string `json:"allow_sms"`
				Avatar                  string `json:"avatar"`
				MascotMessage           string `json:"mascot_message"`
				MentionPn               string `json:"mention_pn"`
				PlayerLikePn            string `json:"player_like_pn"`
				PlayerNicknameUpdate    string `json:"player_nickname_update"`
				TeamName                string `json:"team_name"`
				TeamNameUpdate          string `json:"team_name_update"`
				TradeBlockPn            string `json:"trade_block_pn"`
				TransactionCommissioner string `json:"transaction_commissioner"`
				TransactionFreeAgent    string `json:"transaction_free_agent"`
				TransactionTrade        string `json:"transaction_trade"`
				TransactionWaiver       string `json:"transaction_waiver"`
				UserMessagePn           string `json:"user_message_pn"`
			}{
				TeamName: "Team 1",
			},
		},
		{
			UserID:      "2",
			DisplayName: "User 2",
			Metadata: struct {
				AllowPn                 string `json:"allow_pn"`
				AllowSms                string `json:"allow_sms"`
				Avatar                  string `json:"avatar"`
				MascotMessage           string `json:"mascot_message"`
				MentionPn               string `json:"mention_pn"`
				PlayerLikePn            string `json:"player_like_pn"`
				PlayerNicknameUpdate    string `json:"player_nickname_update"`
				TeamName                string `json:"team_name"`
				TeamNameUpdate          string `json:"team_name_update"`
				TradeBlockPn            string `json:"trade_block_pn"`
				TransactionCommissioner string `json:"transaction_commissioner"`
				TransactionFreeAgent    string `json:"transaction_free_agent"`
				TransactionTrade        string `json:"transaction_trade"`
				TransactionWaiver       string `json:"transaction_waiver"`
				UserMessagePn           string `json:"user_message_pn"`
			}{
				TeamName: "Team 2",
			},
		},
	}

	mockRosters := []Roster{
		{
			RosterID: 1,
			OwnerID:  "1",
		},
		{
			RosterID: 2,
			OwnerID:  "2",
		},
	}

	mockMatchups := []Matchup{
		{
			MatchupID: 1,
			RosterID:  1,
			Points:    100.5,
		},
		{
			MatchupID: 1,
			RosterID:  2,
			Points:    90.5,
		},
	}

	// Create test server for league endpoint
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/v1/league/123":
			league := League{
				Sport: "nfl",
			}
			data, _ := json.Marshal(league)
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		case "/v1/state/nfl":
			state := SportState{
				Week:        1,
				SeasonType:  "regular",
				DisplayWeek: 1,
			}
			data, _ := json.Marshal(state)
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		case "/v1/league/123/users":
			data, _ := json.Marshal(mockUsers)
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		case "/v1/league/123/rosters":
			data, _ := json.Marshal(mockRosters)
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		case "/v1/league/123/matchups/1":
			data, _ := json.Marshal(mockMatchups)
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	// Create client with test server URL
	client := NewClientWithOptions(ClientOptions{
		BaseURL: ts.URL,
	})

	// Test successful request
	scoreboards, err := client.GetScoreboards("123", 1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(scoreboards) != 1 {
		t.Errorf("Expected 1 scoreboard, got %d", len(scoreboards))
	}
	if scoreboards[0].Teamname1 != "Team 1" {
		t.Errorf("Expected team 1 name 'Team 1', got '%s'", scoreboards[0].Teamname1)
	}
	if scoreboards[0].Teamname2 != "Team 2" {
		t.Errorf("Expected team 2 name 'Team 2', got '%s'", scoreboards[0].Teamname2)
	}
	if scoreboards[0].Points1 != 100.5 {
		t.Errorf("Expected team 1 points 100.5, got %f", scoreboards[0].Points1)
	}
	if scoreboards[0].Points2 != 90.5 {
		t.Errorf("Expected team 2 points 90.5, got %f", scoreboards[0].Points2)
	}
}

func TestGetFantasyInfoError(t *testing.T) {
	// Create test server that returns 404 for league
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer ts.Close()

	// Create client with test server URL
	client := NewClientWithOptions(ClientOptions{
		BaseURL: ts.URL,
	})

	// Test league not found
	_, err := client.GetTeamMatchups("nonexistent", 1)
	if err == nil {
		t.Error("Expected error for nonexistent league, got nil")
	}
}

func TestGetFantasyInfoInvalidResponse(t *testing.T) {
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
	_, err := client.GetTeamMatchups("123", 1)
	if err == nil {
		t.Error("Expected error for invalid JSON, got nil")
	}
}

func TestGetFantasyInfoDefaultTeamName(t *testing.T) {
	// Create mock data with empty team name
	mockUsers := []LeagueUser{
		{
			UserID:      "1",
			DisplayName: "User 1",
			Metadata: struct {
				AllowPn                 string `json:"allow_pn"`
				AllowSms                string `json:"allow_sms"`
				Avatar                  string `json:"avatar"`
				MascotMessage           string `json:"mascot_message"`
				MentionPn               string `json:"mention_pn"`
				PlayerLikePn            string `json:"player_like_pn"`
				PlayerNicknameUpdate    string `json:"player_nickname_update"`
				TeamName                string `json:"team_name"`
				TeamNameUpdate          string `json:"team_name_update"`
				TradeBlockPn            string `json:"trade_block_pn"`
				TransactionCommissioner string `json:"transaction_commissioner"`
				TransactionFreeAgent    string `json:"transaction_free_agent"`
				TransactionTrade        string `json:"transaction_trade"`
				TransactionWaiver       string `json:"transaction_waiver"`
				UserMessagePn           string `json:"user_message_pn"`
			}{},
		},
	}

	mockRosters := []Roster{
		{
			RosterID: 1,
			OwnerID:  "1",
		},
	}

	mockMatchups := []Matchup{
		{
			MatchupID: 1,
			RosterID:  1,
			Points:    100.5,
		},
	}

	// Create test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/v1/league/123":
			league := League{
				Sport: "nfl",
			}
			data, _ := json.Marshal(league)
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		case "/v1/state/nfl":
			state := SportState{
				Week:        1,
				SeasonType:  "regular",
				DisplayWeek: 1,
			}
			data, _ := json.Marshal(state)
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		case "/v1/league/123/users":
			data, _ := json.Marshal(mockUsers)
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		case "/v1/league/123/rosters":
			data, _ := json.Marshal(mockRosters)
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		case "/v1/league/123/matchups/1":
			data, _ := json.Marshal(mockMatchups)
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	// Create client with test server URL
	client := NewClientWithOptions(ClientOptions{
		BaseURL: ts.URL,
	})

	// Test default team name
	matchups, err := client.GetTeamMatchups("123", 1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if matchups[0].Teamname1 != "Team User 1" {
		t.Errorf("Expected default team name 'Team User 1', got '%s'", matchups[0].Teamname1)
	}
}
