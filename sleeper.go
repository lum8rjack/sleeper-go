package sleeper

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	sleeperBaseURL string = "https://api.sleeper.app/v1"
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
}

// Create a new Sleeper Client.
func NewClient() Client {
	client := Client{
		httpClient: &http.Client{
			Timeout: time.Minute,
		},
		sleeperURL: sleeperBaseURL,
	}
	return client
}

// Send a basic HTTP GET request.
func (c *Client) getRequest(url string) ([]byte, error) {
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
