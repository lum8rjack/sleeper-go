// Package sleeper provides a client for interacting with the Sleeper API.
//
// The Sleeper API is a read-only API that provides data about NFL games, players, and teams.
// Be mindful of the frequency of calls. A general rule is to stay under 1000 API calls per minute, otherwise, you risk being IP-blocked.
//
// The API is documented at https://docs.sleeper.app/api.
package sleeper

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

const (
	sleeperBaseURL string = "https://api.sleeper.app"
)

var (
	errorCodes = map[int]string{
		400: "Bad Request",
		404: "Not Found",
		429: "Too Many Requests",
		500: "Internal Server Error",
		503: "Service Unavailable",
	}
)

// Client for interacting with the read-only Sleeper API.
type Client struct {
	httpClient *http.Client
	sleeperURL string
	limiter    *rate.Limiter
}

// ClientOption is a function that modifies a Client.
type ClientOptions struct {
	BaseURL   string
	Timeout   time.Duration
	RateLimit float64 // Rate limit in requests per second
}

// Create a new Sleeper Client.
func NewClient() Client {
	client := Client{
		httpClient: &http.Client{
			Timeout: time.Minute,
		},
		sleeperURL: sleeperBaseURL,
		limiter:    rate.NewLimiter(rate.Limit(1000/60), 1), // Default: 1000 req/min, burst 1
	}
	return client
}

// NewClientWithOptions creates a new Sleeper Client with the given options.
func NewClientWithOptions(opts ClientOptions) Client {
	client := Client{
		httpClient: &http.Client{
			Timeout: time.Minute,
		},
		sleeperURL: sleeperBaseURL,
		limiter:    rate.NewLimiter(rate.Limit(1000/60), 1), // Default: 1000 req/min, burst 1
	}

	// Set the timeout for the HTTP client
	if opts.Timeout > 0 {
		client.httpClient.Timeout = opts.Timeout
	}

	// Set the base URL for the Sleeper API
	if opts.BaseURL != "" {
		client.sleeperURL = opts.BaseURL
	}

	// Set the rate limit for the Sleeper API
	if opts.RateLimit > 0 {
		client.limiter = rate.NewLimiter(rate.Limit(opts.RateLimit), 1)
	}

	return client
}

// Get the base URL for the Sleeper API.
func (c *Client) BaseURL() string {
	return c.sleeperURL
}

// Send a basic HTTP GET request.
func (c *Client) getRequest(url string) ([]byte, error) {
	// Wait for rate limiter
	if err := c.limiter.Wait(context.Background()); err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		val, ok := errorCodes[resp.StatusCode]
		if ok {
			e := fmt.Sprintf("web request error: %d %s\n", resp.StatusCode, val)
			return nil, errors.New(e)
		} else {
			e := fmt.Sprintf("web request error: %d\n", resp.StatusCode)
			return nil, errors.New(e)
		}
	}

	return io.ReadAll(resp.Body)
}
