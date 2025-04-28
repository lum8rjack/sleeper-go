package sleeper

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUserByUsername(t *testing.T) {
	// Create mock user data
	mockUser := User{
		UserID:      "123",
		Username:    "testuser",
		DisplayName: "Test User",
		Avatar:      "https://example.com/avatar.jpg",
		IsBot:       false,
	}

	// Convert mock data to JSON
	mockData, err := json.Marshal(mockUser)
	if err != nil {
		t.Fatalf("Failed to marshal mock data: %v", err)
	}

	// Create test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/user/testuser" {
			t.Errorf("Expected path /v1/user/testuser, got %s", r.URL.Path)
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
	user, err := client.GetUserByUsername("testuser")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if user.Username != "testuser" {
		t.Errorf("Expected username 'testuser', got '%s'", user.Username)
	}
	if user.DisplayName != "Test User" {
		t.Errorf("Expected display name 'Test User', got '%s'", user.DisplayName)
	}
	if user.UserID != "123" {
		t.Errorf("Expected user ID '123', got '%s'", user.UserID)
	}
}

func TestGetUserByID(t *testing.T) {
	// Create mock user data
	mockUser := User{
		UserID:      "123",
		Username:    "testuser",
		DisplayName: "Test User",
		Avatar:      "https://example.com/avatar.jpg",
		IsBot:       false,
	}

	// Convert mock data to JSON
	mockData, err := json.Marshal(mockUser)
	if err != nil {
		t.Fatalf("Failed to marshal mock data: %v", err)
	}

	// Create test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/user/123" {
			t.Errorf("Expected path /v1/user/123, got %s", r.URL.Path)
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
	user, err := client.GetUserByID("123")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if user.UserID != "123" {
		t.Errorf("Expected user ID '123', got '%s'", user.UserID)
	}
	if user.Username != "testuser" {
		t.Errorf("Expected username 'testuser', got '%s'", user.Username)
	}
}

func TestGetUserNotFound(t *testing.T) {
	// Create test server that returns 404
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer ts.Close()

	// Create client with test server URL
	client := NewClientWithOptions(ClientOptions{
		BaseURL: ts.URL,
	})

	// Test user not found by username
	_, err := client.GetUserByUsername("nonexistent")
	if err == nil {
		t.Error("Expected error for nonexistent user, got nil")
	}

	// Test user not found by ID
	_, err = client.GetUserByID("nonexistent")
	if err == nil {
		t.Error("Expected error for nonexistent user ID, got nil")
	}
}

func TestGetUserInvalidResponse(t *testing.T) {
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
	_, err := client.GetUserByUsername("testuser")
	if err == nil {
		t.Error("Expected error for invalid JSON, got nil")
	}
}
