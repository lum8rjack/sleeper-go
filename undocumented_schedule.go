package sleeper

import (
	"encoding/json"
	"fmt"
)

type NflSchedule []struct {
	Status string `json:"status"`
	Date   string `json:"date"`
	Home   string `json:"home"`
	Week   int    `json:"week"`
	GameID string `json:"game_id"`
	Away   string `json:"away"`
}

// Get NFL schedule.
// `GET https://api.sleeper.app/schedule/nfl/<regular or post>/<year>`
func (c *Client) GetNflSchedule(year int, postseason bool) (NflSchedule, error) {
	schedule := NflSchedule{}
	reg := "regular"
	if postseason {
		reg = "post"
	}

	url := fmt.Sprintf("%s/schedule/nfl/%s", sleeperUndocumentedURL, reg)

	data, err := c.getRequest(url)
	if err != nil {
		return schedule, err
	}

	err = json.Unmarshal(data, &schedule)

	return schedule, err
}
