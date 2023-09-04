# sleeper-go
Go library for the [Sleeper](https://sleeper.com/) fantasy sports free read-only API. According to their [documentation](https://docs.sleeper.com), stay under 1000 API calls per minute, otherwise, you risk being IP-blocked.

## Players

The players API is not intended to be called every time you need to look up players due to the large file size. It should not be called more than once per day.

With this in mind, I have create methods to save the players information to disk and read the saved file from disk. 

```go
// Get all players - use this sparingly
func (c *Client) GetAllPlayers(sport string) (Players, error)

// Get all players and save the details to a file
func (c *Client) SaveAllPlayers(sport string, file string) (bool, error)

// Get all players from a saved file
func GetAllPlayers(file string) (Players, error)
```

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

## Sleeper API Implementation Status

* [x] User
    * [x] User object - `GET https://api.sleeper.app/v1/user/<username>`
    * [x] User object - `GET https://api.sleeper.app/v1//user/<user_id>`
* [x] Avatars
    * [x] Full size URL - `GET https://sleepercdn.com/avatars/<avatar_id>`
    * [x] Thumbnail URL - `GET https://sleepercdn.com/avatars/thumbs/<avatar_id>`
* [x] Leagues
    * [x] Get all leagues for user - `GET https://api.sleeper.app/v1/user/<user_id>/leagues/<sport>/<season>`
    * [x] Get a specific league - `GET https://api.sleeper.app/v1/league/<league_id>`
    * [x] Get rosters in a league - `GET https://api.sleeper.app/v1/league/<league_id>/rosters`
    * [x] Get users in a league - `GET https://api.sleeper.app/v1/league/<league_id>/users`
    * [x] Get matchups in a league - `GET https://api.sleeper.app/v1/league/<league_id>/matchups/<week>`
    * [x] Get the playoff bracket (winners bracket) - `GET https://api.sleeper.app/v1/league/<league_id>/winners_bracket`
    * [x] Get the playoff bracket (loses bracket) - `GET https://api.sleeper.app/v1/league/<league_id>/loses_bracket`
    * [x] Get transactions - `GET https://api.sleeper.app/v1/league/<league_id>/transactions/<round>`
    * [x] Get traded picks - `GET https://api.sleeper.app/v1/league/<league_id>/traded_picks`
    * [x] Get sport state - `GET https://api.sleeper.app/v1/state/<sport>`
* [x] Drafts
    * [x] Get all drafts for user - `GET https://api.sleeper.app/v1/user/<user_id>/drafts/<sport>/<season>`
    * [x] Get all drafts for a league - `GET https://api.sleeper.app/v1/league/<league_id>/drafts`
    * [x] Get a specific draft - `GET https://api.sleeper.app/v1/draft/<draft_id>`
    * [x] Get all picks in a draft - `GET https://api.sleeper.app/v1/draft/<draft_id>/picks`
    * [x] Get traded picks in a draft - `GET https://api.sleeper.app/v1/draft/<draft_id>/traded_picks`
* [x] Players
    * [x] Get all players - `GET https://api.sleeper.app/v1/players/<sport>`
    * [x] Trending players - `GET https://api.sleeper.app/v1/players/<sport>/trending/<type>?lookback_hours=<hours>&limit=<int>`


## License

MIT
