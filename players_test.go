package sleeper

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestGetAllPlayers(t *testing.T) {
	// Create mock player data
	mockPlayers := Players{
		"1": Player{
			PlayerID: "1",
			FullName: "Test Player 1",
			Team:     "TEST",
		},
		"2": Player{
			PlayerID: "2",
			FullName: "Test Player 2",
			Team:     "TEST",
		},
	}

	// Convert mock data to JSON
	mockData, err := json.Marshal(mockPlayers)
	if err != nil {
		t.Fatalf("Failed to marshal mock data: %v", err)
	}

	// Create test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(mockData)
	}))
	defer ts.Close()

	// Create client with test server URL
	client := NewClientWithOptions(ClientOptions{
		BaseURL: ts.URL,
	})

	// Test successful request
	players, err := client.GetAllPlayers("nfl")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(players) != 2 {
		t.Errorf("Expected 2 players, got %d", len(players))
	}

	if players["1"].FullName != "Test Player 1" {
		t.Errorf("Expected player name 'Test Player 1', got '%s'", players["1"].FullName)
	}
}

func TestSaveAllPlayers(t *testing.T) {
	// Create mock player data
	mockPlayers := Players{
		"1": Player{
			PlayerID: "1",
			FullName: "Test Player",
			Team:     "TEST",
		},
	}

	// Convert mock data to JSON
	mockData, err := json.Marshal(mockPlayers)
	if err != nil {
		t.Fatalf("Failed to marshal mock data: %v", err)
	}

	// Create test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(mockData)
	}))
	defer ts.Close()

	// Create client with test server URL
	client := NewClientWithOptions(ClientOptions{
		BaseURL: ts.URL,
	})

	// Create temporary file
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "players.json")

	// Test saving players
	success, err := client.SaveAllPlayers("nfl", filePath)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !success {
		t.Error("Expected save to be successful")
	}

	// Verify file contents
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		t.Errorf("Failed to read saved file: %v", err)
	}

	var savedPlayers Players
	err = json.Unmarshal(fileData, &savedPlayers)
	if err != nil {
		t.Errorf("Failed to unmarshal saved data: %v", err)
	}

	if len(savedPlayers) != 1 {
		t.Errorf("Expected 1 player in saved file, got %d", len(savedPlayers))
	}
}

func TestGetAllPlayersFromFile(t *testing.T) {
	// Create mock player data
	mockPlayers := Players{
		"1": Player{
			PlayerID: "1",
			FullName: "Test Player",
			Team:     "TEST",
		},
	}

	// Convert mock data to JSON
	mockData, err := json.Marshal(mockPlayers)
	if err != nil {
		t.Fatalf("Failed to marshal mock data: %v", err)
	}

	// Create temporary file
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "players.json")

	// Write mock data to file
	err = os.WriteFile(filePath, mockData, 0644)
	if err != nil {
		t.Fatalf("Failed to write mock data to file: %v", err)
	}

	// Test reading players from file
	players, err := GetAllPlayers(filePath)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(players) != 1 {
		t.Errorf("Expected 1 player, got %d", len(players))
	}

	if players["1"].FullName != "Test Player" {
		t.Errorf("Expected player name 'Test Player', got '%s'", players["1"].FullName)
	}
}

func TestGetTrendingPlayers(t *testing.T) {
	// Create mock trending players data
	mockTrendingPlayers := []TrendingPlayer{
		{
			Count:    100,
			PlayerID: "1",
		},
		{
			Count:    50,
			PlayerID: "2",
		},
	}

	// Convert mock data to JSON
	mockData, err := json.Marshal(mockTrendingPlayers)
	if err != nil {
		t.Fatalf("Failed to marshal mock data: %v", err)
	}

	// Create test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(mockData)
	}))
	defer ts.Close()

	// Create client with test server URL
	client := NewClientWithOptions(ClientOptions{
		BaseURL: ts.URL,
	})

	// Test getting trending players
	trendingPlayers, err := client.GetTrendingPlayers("nfl", "add")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(trendingPlayers) != 2 {
		t.Errorf("Expected 2 trending players, got %d", len(trendingPlayers))
	}

	if trendingPlayers[0].Count != 100 {
		t.Errorf("Expected count 100, got %d", trendingPlayers[0].Count)
	}
}

func TestGetTrendingPlayersParams(t *testing.T) {
	// Create mock trending players data
	mockTrendingPlayers := []TrendingPlayer{
		{
			Count:    100,
			PlayerID: "1",
		},
	}

	// Convert mock data to JSON
	mockData, err := json.Marshal(mockTrendingPlayers)
	if err != nil {
		t.Fatalf("Failed to marshal mock data: %v", err)
	}

	// Create test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(mockData)
	}))
	defer ts.Close()

	// Create client with test server URL
	client := NewClientWithOptions(ClientOptions{
		BaseURL: ts.URL,
	})

	// Test getting trending players with parameters
	trendingPlayers, err := client.GetTrendingPlayersParams("nfl", "add", 24, 10)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(trendingPlayers) != 1 {
		t.Errorf("Expected 1 trending player, got %d", len(trendingPlayers))
	}

	if trendingPlayers[0].Count != 100 {
		t.Errorf("Expected count 100, got %d", trendingPlayers[0].Count)
	}
}
