package sleeper

import (
	"encoding/json"
	"fmt"
	"os"
)

type TrendingPlayer struct {
	Count    int    `json:"count"`
	PlayerID string `json:"player_id"`
}

type Players map[string]Player

type Player struct {
	Active                bool        `json:"active"`
	Age                   int         `json:"age"`
	BirthCity             interface{} `json:"birth_city"`
	BirthCountry          string      `json:"birth_country"`
	BirthDate             string      `json:"birth_date"`
	BirthState            interface{} `json:"birth_state"`
	College               string      `json:"college"`
	DepthChartOrder       int         `json:"depth_chart_order"`
	DepthChartPosition    string      `json:"depth_chart_position"`
	EspnID                int         `json:"espn_id"`
	FantasyDataID         int         `json:"fantasy_data_id"`
	FantasyPositions      []string    `json:"fantasy_positions"`
	FirstName             string      `json:"first_name"`
	FullName              string      `json:"full_name"`
	GsisID                string      `json:"gsis_id"`
	Hashtag               string      `json:"hashtag"`
	Height                string      `json:"height"`
	HighSchool            string      `json:"high_school"`
	InjuryBodyPart        string      `json:"injury_body_part"`
	InjuryNotes           string      `json:"injury_notes"`
	InjuryStartDate       interface{} `json:"injury_start_date"`
	InjuryStatus          string      `json:"injury_status"`
	LastName              string      `json:"last_name"`
	Metadata              interface{} `json:"metadata"`
	NewsUpdated           int64       `json:"news_updated"`
	Number                int         `json:"number"`
	PandascoreID          interface{} `json:"pandascore_id"`
	PlayerID              string      `json:"player_id"`
	Position              string      `json:"position"`
	PracticeDescription   string      `json:"practice_description"`
	PracticeParticipation string      `json:"practice_participation"`
	RotowireID            int         `json:"rotowire_id"`
	RotoworldID           int         `json:"rotoworld_id"`
	SearchFirstName       string      `json:"search_first_name"`
	SearchFullName        string      `json:"search_full_name"`
	SearchLastName        string      `json:"search_last_name"`
	SearchRank            int         `json:"search_rank"`
	Sport                 string      `json:"sport"`
	SportradarID          string      `json:"sportradar_id"`
	StatsID               string      `json:"stats_id"`
	Status                string      `json:"status"`
	SwishID               int         `json:"swish_id"`
	Team                  string      `json:"team"`
	Weight                string      `json:"weight"`
	YahooID               int         `json:"yahoo_id"`
	YearsExp              int         `json:"years_exp"`
}

// Get all players.
// (GET `https://api.sleeper.app/v1/players/<sport>`)
//
// You should save this information on your own servers as this is not intended to be called every time you need to look up players due to the filesize being close to 5MB in size.
// You do not need to call this endpoint more than once per day.
func (c *Client) GetAllPlayers(sport string) (Players, error) {
	players := Players{}

	url := fmt.Sprintf("%s/players/%s", c.sleeperURL, sport)

	data, err := c.getRequest(url)
	if err != nil {
		return players, err
	}

	err = json.Unmarshal(data, &players)

	return players, err
}

// Get all players and save to a file.
// (GET `https://api.sleeper.app/v1/players/<sport>`)
//
// You should save this information on your own servers as this is not intended to be called every time you need to look up players due to the filesize being close to 5MB in size.
// You do not need to call this endpoint more than once per day.
func (c *Client) SaveAllPlayers(sport string, file string) (bool, error) {
	url := fmt.Sprintf("%s/players/%s", c.sleeperURL, sport)

	data, err := c.getRequest(url)
	if err != nil {
		return false, err
	}

	f, err := os.Create(file)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Get all players from a file.
func GetAllPlayers(file string) (Players, error) {
	players := Players{}

	data, err := os.ReadFile(file)
	if err != nil {
		return players, err
	}

	err = json.Unmarshal(data, &players)

	return players, err
}

// Get a list of trending players based on adds or drops in the past 24 hours. Trending type is add or drop.
// (GET `https://api.sleeper.app/v1/players/<sport>/trending/<type>`)
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
// (GET `https://api.sleeper.app/v1/players/<sport>/trending/<type>?lookback_hours=<hours>&limit=<int>`)
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
