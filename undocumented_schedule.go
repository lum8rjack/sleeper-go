package sleeper

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
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

	// Sleeper only has data from 2009 to present
	if year < 2009 || year > time.Now().Year() {
		return schedule, errors.New("invalid year - must be between 2008 and current")
	}

	url := fmt.Sprintf("%s/schedule/nfl/%s/%d", c.sleeperURL, reg, year)

	data, err := c.getRequest(url)
	if err != nil {
		return schedule, err
	}

	err = json.Unmarshal(data, &schedule)

	return schedule, err
}
