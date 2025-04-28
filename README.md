# sleeper-go
Go library for the [Sleeper](https://sleeper.com/) fantasy sports free read-only API. According to their [documentation](https://docs.sleeper.com), stay under 1000 API calls per minute, otherwise, you risk being IP-blocked.

## Usage
Example code that gets the users for the specified league and outputs their display name and team name.

```go
package main

import (
	"fmt"
	"log"

	"github.com/lum8rjack/sleeper-go"
)

func main() {
    leagueID := "123456789012345678" // Fake league Id
	botClient := sleeper.NewClient()

	leagueUsers, err := botClient.GetLeagueUsers(leagueID)
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range leagueUsers {
		fmt.Printf("%s (Team name: %s)\n", user.DisplayName, user.Metadata.TeamName)
	}
}

```

Sample output:
```bash
2KSports (Team name: Game of End Zones)
TOBOT (Team name: Saving Matt Ryan)
...
```

## Custom Methods

I created a few custom methods that require correlating results from multiple API calls in order to receive the data.

### GetMatchups

This method gets the matchups for the specified week and returns each team and their wins and losses.
```go
type TeamMatchup struct {
	Teamname1   string
	Teamname2   string
	Team1Losses int
	Team2Losses int
	Team1Wins   int
	Team2Wins   int
}

func (c *Client) GetTeamMatchups(league_id string, week int) ([]TeamMatchup, error)
```

### GetScoreboards

This method gets the scoreboard for each matchup for the specified week and returns each team and points.
```go
type Scoreboard struct {
	Teamname1 string  `json:"teamname_1"`
	Teamname2 string  `json:"teamname_2"`
	Points1   float32 `json:"points_1"`
	Points2   float32 `json:"points_2"`
}

func (c *Client) GetScoreboards(league_id string, week int) ([]Scoreboard, error)
```

### Players

The players API is not intended to be called every time you need to look up players due to the large file size. It should not be called more than once per day.

With this in mind, I have created additional methods to save the players information to disk and read the saved file from disk. 

```go
// Get all players - use this sparingly
func (c *Client) GetAllPlayers(sport string) (Players, error)

// Get all players and save the details to a file
func (c *Client) SaveAllPlayers(sport string, file string) (bool, error)

// Get all players from a saved file
func GetAllPlayers(file string) (Players, error)
```

## Sleeper API Implementation Status

The following table shows the implementation status of all known Sleeper API endpoints in this library. The table includes both officially documented endpoints from Sleeper's API documentation as well as several undocumented endpoints that were discovered during development. All endpoints, both documented and undocumented, have been fully implemented.


| Category       | Endpoint                                                                 | Description                                                                 | Implemented |
|----------------|--------------------------------------------------------------------------|-----------------------------------------------------------------------------|-------------|
| User           | `GET https://api.sleeper.app/v1/user/<username>`                         | Get user details by username                                                | ✅          |
| User           | `GET https://api.sleeper.app/v1/user/<user_id>`                          | Get user details by user ID                                                 | ✅          |
| Avatars        | `GET https://sleepercdn.com/avatars/<avatar_id>`                         | Get full-size avatar image                                                  | ✅          |
| Avatars        | `GET https://sleepercdn.com/avatars/thumbs/<avatar_id>`                  | Get thumbnail avatar image                                                  | ✅          |
| Leagues        | `GET https://api.sleeper.app/v1/user/<user_id>/leagues/<sport>/<season>` | Get all leagues a user is in for a specific sport and season                | ✅          |
| Leagues        | `GET https://api.sleeper.app/v1/league/<league_id>`                      | Get details of a specific league                                            | ✅          |
| Leagues        | `GET https://api.sleeper.app/v1/league/<league_id>/rosters`              | Get all rosters in a league                                                 | ✅          |
| Leagues        | `GET https://api.sleeper.app/v1/league/<league_id>/users`                | Get all users in a league                                                   | ✅          |
| Leagues        | `GET https://api.sleeper.app/v1/league/<league_id>/matchups/<week>`      | Get all matchups for a specific week in a league                            | ✅          |
| Leagues        | `GET https://api.sleeper.app/v1/league/<league_id>/winners_bracket`      | Get the winners bracket for league playoffs                                 | ✅          |
| Leagues        | `GET https://api.sleeper.app/v1/league/<league_id>/losers_bracket`       | Get the losers bracket for league playoffs                                  | ✅          |
| Leagues        | `GET https://api.sleeper.app/v1/league/<league_id>/transactions/<round>` | Get all transactions for a specific round in a league                       | ✅          |
| Leagues        | `GET https://api.sleeper.app/v1/league/<league_id>/traded_picks`         | Get all traded draft picks in a league                                      | ✅          |
| Leagues        | `GET https://api.sleeper.app/v1/state/<sport>`                           | Get current state of a sport (season, week, etc.)                           | ✅          |
| Drafts         | `GET https://api.sleeper.app/v1/user/<user_id>/drafts/<sport>/<season>`  | Get all drafts a user is in for a specific sport and season                 | ✅          |
| Drafts         | `GET https://api.sleeper.app/v1/league/<league_id>/drafts`               | Get all drafts in a league                                                  | ✅          |
| Drafts         | `GET https://api.sleeper.app/v1/draft/<draft_id>`                        | Get details of a specific draft                                             | ✅          |
| Drafts         | `GET https://api.sleeper.app/v1/draft/<draft_id>/picks`                  | Get all picks in a draft                                                    | ✅          |
| Drafts         | `GET https://api.sleeper.app/v1/draft/<draft_id>/traded_picks`           | Get all traded picks in a draft                                             | ✅          |
| Players        | `GET https://api.sleeper.app/v1/players/<sport>`                         | Get all players for a sport (large response, use sparingly)                 | ✅          |
| Players        | `GET https://api.sleeper.app/v1/players/<sport>/trending/<type>`         | Get trending players with optional lookback hours and limit                 | ✅          |
| Undocumented   | `GET https://api.sleeper.app/projections/nfl/<season>/<week>`            | Get player projections for a specific NFL season and week                   | ✅          |
| Undocumented   | `GET https://api.sleeper.app/stats/nfl/player/<player_id>`               | Get season stats for a specific NFL player                                  | ✅          |
| Undocumented   | `GET https://api.sleeper.app/players/nfl/<player_id>`                    | Get detailed information for a specific NFL player                          | ✅          |
| Undocumented   | `GET https://api.sleeper.app/players/nfl/research/<regular/post>/<year>/<week>` | Get player research data (ownership, start rates) for a specific week | ✅ |
| Undocumented   | `GET https://api.sleeper.app/players/nfl/<team>/depth_chart`             | Get depth chart for a specific NFL team                                     | ✅          |
| Undocumented   | `GET https://api.sleeper.app/schedule/nfl/<regular/post>/<year>`         | Get NFL schedule for a specific season (regular or postseason)              | ✅          |


## License

MIT
