package sleeper

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetNflTeamDepthChart(t *testing.T) {
	// Create mock depth chart data
	mockDepthChart := TeamDepthChart{
		Qb:   []string{"player1", "player2"},
		Rb:   []string{"player3", "player4"},
		Wr1:  []string{"player5", "player6"},
		Wr2:  []string{"player7", "player8"},
		Wr3:  []string{"player9", "player10"},
		Te:   []string{"player11", "player12"},
		Ol:   []string{"player13", "player14"},
		Lde:  []string{"player15", "player16"},
		Rde:  []string{"player17", "player18"},
		Ldt:  []string{"player19", "player20"},
		Rdt:  []string{"player21", "player22"},
		Lolb: []string{"player23", "player24"},
		Rolb: []string{"player25", "player26"},
		Mlb:  []string{"player27", "player28"},
		Lcb:  []string{"player29", "player30"},
		Rcb:  []string{"player31", "player32"},
		Fs:   []string{"player33", "player34"},
		Ss:   []string{"player35", "player36"},
		K:    []string{"player37", "player38"},
		P:    []string{"player39", "player40"},
		Ls:   []string{"player41", "player42"},
		Db:   []string{"player43", "player44"},
		Dl:   []string{"player45", "player46"},
		Lb:   []string{"player47", "player48"},
		Nb:   []string{"player49", "player50"},
	}

	// Convert mock data to JSON
	mockData, err := json.Marshal(mockDepthChart)
	if err != nil {
		t.Fatalf("Failed to marshal mock data: %v", err)
	}

	// Create test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/players/nfl/SF/depth_chart" {
			t.Errorf("Expected path /players/nfl/SF/depth_chart, got %s", r.URL.Path)
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
	depthChart, err := client.GetNflTeamDepthChart("SF")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Test QB depth
	if len(depthChart.Qb) != 2 {
		t.Errorf("Expected 2 QBs, got %d", len(depthChart.Qb))
	}
	if depthChart.Qb[0] != "player1" {
		t.Errorf("Expected QB1 'player1', got '%s'", depthChart.Qb[0])
	}

	// Test RB depth
	if len(depthChart.Rb) != 2 {
		t.Errorf("Expected 2 RBs, got %d", len(depthChart.Rb))
	}
	if depthChart.Rb[0] != "player3" {
		t.Errorf("Expected RB1 'player3', got '%s'", depthChart.Rb[0])
	}

	// Test WR depth
	if len(depthChart.Wr1) != 2 {
		t.Errorf("Expected 2 WR1s, got %d", len(depthChart.Wr1))
	}
	if depthChart.Wr1[0] != "player5" {
		t.Errorf("Expected WR1 'player5', got '%s'", depthChart.Wr1[0])
	}

	// Test defensive positions
	if len(depthChart.Lde) != 2 {
		t.Errorf("Expected 2 LDEs, got %d", len(depthChart.Lde))
	}
	if depthChart.Lde[0] != "player15" {
		t.Errorf("Expected LDE1 'player15', got '%s'", depthChart.Lde[0])
	}

	// Test special teams
	if len(depthChart.K) != 2 {
		t.Errorf("Expected 2 Ks, got %d", len(depthChart.K))
	}
	if depthChart.K[0] != "player37" {
		t.Errorf("Expected K1 'player37', got '%s'", depthChart.K[0])
	}
}

func TestGetNflTeamDepthChartError(t *testing.T) {
	// Create test server that returns 404
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer ts.Close()

	// Create client with test server URL
	client := NewClientWithOptions(ClientOptions{
		BaseURL: ts.URL,
	})

	// Test team not found
	_, err := client.GetNflTeamDepthChart("nonexistent")
	if err == nil {
		t.Error("Expected error for nonexistent team, got nil")
	}
}

func TestGetNflTeamDepthChartInvalidResponse(t *testing.T) {
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
	_, err := client.GetNflTeamDepthChart("SF")
	if err == nil {
		t.Error("Expected error for invalid JSON, got nil")
	}
}

func TestGetNflTeamDepthChartEmptyResponse(t *testing.T) {
	// Create test server that returns empty depth chart
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{}"))
	}))
	defer ts.Close()

	// Create client with test server URL
	client := NewClientWithOptions(ClientOptions{
		BaseURL: ts.URL,
	})

	// Test empty response
	depthChart, err := client.GetNflTeamDepthChart("SF")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify all positions are empty
	if len(depthChart.Qb) != 0 {
		t.Errorf("Expected empty QB list, got %d players", len(depthChart.Qb))
	}
	if len(depthChart.Rb) != 0 {
		t.Errorf("Expected empty RB list, got %d players", len(depthChart.Rb))
	}
	if len(depthChart.Wr1) != 0 {
		t.Errorf("Expected empty WR1 list, got %d players", len(depthChart.Wr1))
	}
}
