package sleeper

import (
	"encoding/json"
	"fmt"
)

type TrendingPlayer struct {
	PlayerID string `json:"player_id"`
	Count    int    `json:"count"`
}

// Get a list of trending players based on adds or drops in the past 24 hours. Trending type is add or drop.
// GET https://api.sleeper.app/v1/players/<sport>/trending/<type>
func (c *Client) GetTrendingPlayers(sport string, trending_type string) ([]TrendingPlayer, error) {
	trendingPlayer := []TrendingPlayer{}

	url := fmt.Sprintf("%s/players/%s/trending/%s", c.sleeperURL, sport, trending_type)

	data, err := c.getRequest(url)
	if err != nil {
		return trendingPlayer, err
	}

	err = json.Unmarshal(data, &trendingPlayer)

	return trendingPlayer, err
}

// Get a list of trending players based on adds or drops in the past X hours. Trending type is add or drop.
// GET https://api.sleeper.app/v1/players/<sport>/trending/<type>?lookback_hours=<hours>&limit=<int>
func (c *Client) GetTrendingPlayersParams(sport string, trending_type string, hours int, limit int) ([]TrendingPlayer, error) {
	trendingPlayer := []TrendingPlayer{}

	url := fmt.Sprintf("%s/players/%s/trending/%s?loopback_hours=%d&limit=%d", c.sleeperURL, sport, trending_type, hours, limit)

	data, err := c.getRequest(url)
	if err != nil {
		return trendingPlayer, err
	}

	err = json.Unmarshal(data, &trendingPlayer)

	return trendingPlayer, err
}
