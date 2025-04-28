package sleeper

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	client := NewClient()
	if client.sleeperURL != sleeperBaseURL {
		t.Errorf("Expected base URL %s, got %s", sleeperBaseURL, client.sleeperURL)
	}
	if client.httpClient.Timeout != time.Minute {
		t.Errorf("Expected timeout %v, got %v", time.Minute, client.httpClient.Timeout)
	}
}

func TestNewClientWithOptions(t *testing.T) {
	tests := []struct {
		name     string
		opts     ClientOptions
		expected Client
	}{
		{
			name: "with custom timeout",
			opts: ClientOptions{
				Timeout: 30 * time.Second,
			},
			expected: Client{
				httpClient: &http.Client{
					Timeout: 30 * time.Second,
				},
				sleeperURL: sleeperBaseURL,
			},
		},
		{
			name: "with custom base URL",
			opts: ClientOptions{
				BaseURL: "http://custom-url.com",
			},
			expected: Client{
				httpClient: &http.Client{
					Timeout: time.Minute,
				},
				sleeperURL: "http://custom-url.com",
			},
		},
		{
			name: "with custom rate limit",
			opts: ClientOptions{
				RateLimit: 5,
			},
			expected: Client{
				httpClient: &http.Client{
					Timeout: time.Minute,
				},
				sleeperURL: sleeperBaseURL,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewClientWithOptions(tt.opts)
			if tt.opts.Timeout > 0 && client.httpClient.Timeout != tt.opts.Timeout {
				t.Errorf("Expected timeout %v, got %v", tt.opts.Timeout, client.httpClient.Timeout)
			}
			if tt.opts.BaseURL != "" && client.sleeperURL != tt.opts.BaseURL {
				t.Errorf("Expected base URL %s, got %s", tt.opts.BaseURL, client.sleeperURL)
			}
			if tt.opts.RateLimit > 0 {
				// Test rate limiting by making multiple requests
				start := time.Now()
				for i := 0; i < 10; i++ {
					client.limiter.Wait(context.Background())
				}
				elapsed := time.Since(start)
				expectedMinTime := time.Duration(9) * time.Second / time.Duration(tt.opts.RateLimit)
				if elapsed < expectedMinTime {
					t.Errorf("Rate limiting not working as expected. Expected minimum time %v, got %v", expectedMinTime, elapsed)
				}
			}
		})
	}
}

func TestGetRequest(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"test": "data"}`))
	}))
	defer ts.Close()

	// Create client with test server URL
	client := NewClientWithOptions(ClientOptions{
		BaseURL: ts.URL,
	})

	// Test successful request
	data, err := client.getRequest(client.sleeperURL + "/test")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if string(data) != `{"test": "data"}` {
		t.Errorf("Expected response %s, got %s", `{"test": "data"}`, string(data))
	}

	// Test rate limiting
	start := time.Now()
	for i := 0; i < 10; i++ {
		_, err := client.getRequest(client.sleeperURL + "/test")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	}
	elapsed := time.Since(start)
	expectedMinTime := time.Duration(9) * time.Second / time.Duration(1000/60)
	if elapsed < expectedMinTime {
		t.Errorf("Rate limiting not working as expected. Expected minimum time %v, got %v", expectedMinTime, elapsed)
	}
}

func TestGetRequestError(t *testing.T) {
	// Create a test server that returns 404
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer ts.Close()

	// Create client with test server URL
	client := NewClientWithOptions(ClientOptions{
		BaseURL: ts.URL,
	})

	// Test error response
	_, err := client.getRequest("/nonexistent")
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
