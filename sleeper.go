package sleeper

import (
	"io"
	"net/http"
	"time"
)

const (
	sleeperBaseURL string = "https://api.sleeper.app/v1"
)

// Client for interacting with the read-only Sleeper API
type Client struct {
	httpClient *http.Client
	sleeperURL string
}

// Creates a new Sleeper Client
func NewClient() Client {
	client := Client{
		httpClient: &http.Client{
			Timeout: time.Minute,
		},
		sleeperURL: sleeperBaseURL,
	}
	return client
}

// Send a basic HTTP GET request
func (c *Client) getRequest(url string) ([]byte, error) {
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
