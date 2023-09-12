package sleeper

type customTeamInfo struct {
	DisplayName string
	Losses      int
	MatchupID   int
	OwnerID     string
	Points      float32
	RosterID    int
	Teamname    string
	Week        int
	Wins        int
}

type TeamMatchup struct {
	Teamname1   string
	Teamname2   string
	Team1Losses int
	Team2Losses int
	Team1Wins   int
	Team2Wins   int
}

// Get matchup information for the specified week.
func (c *Client) GetTeamMatchups(league_id string, week int) ([]TeamMatchup, error) {
	var matchups []TeamMatchup

	teaminfo, err := c.getFantasyInfo(league_id, week)
	if err != nil {
		return matchups, err
	}

	allmatchups := make(map[int]TeamMatchup)

	for _, team := range teaminfo {
		if _, ok := allmatchups[team.MatchupID]; ok {
			newteam := allmatchups[team.MatchupID]
			newteam.Teamname2 = team.Teamname
			newteam.Team2Wins = team.Wins
			newteam.Team2Losses = team.Losses
			allmatchups[team.MatchupID] = newteam

		} else {
			mu := TeamMatchup{
				Teamname1:   team.Teamname,
				Team1Wins:   team.Wins,
				Team1Losses: team.Losses,
			}
			allmatchups[team.MatchupID] = mu
		}
	}

	for _, m := range allmatchups {
		matchups = append(matchups, m)
	}

	clear(allmatchups)

	return matchups, nil
}

type Scoreboard struct {
	Teamname1 string  `json:"teamname_1"`
	Teamname2 string  `json:"teamname_2"`
	Points1   float32 `json:"points_1"`
	Points2   float32 `json:"points_2"`
}

// Get the scoreboard for each game for the specified week.
func (c *Client) GetScoreboards(league_id string, week int) ([]Scoreboard, error) {
	var scoreboards []Scoreboard

	teaminfo, err := c.getFantasyInfo(league_id, week)
	if err != nil {
		return scoreboards, err
	}

	allscoreboards := make(map[int]Scoreboard)

	for _, team := range teaminfo {
		if _, ok := allscoreboards[team.MatchupID]; ok {
			newteam := allscoreboards[team.MatchupID]
			newteam.Teamname2 = team.Teamname
			newteam.Points2 = team.Points
			allscoreboards[team.MatchupID] = newteam

		} else {
			sb := Scoreboard{
				Teamname1: team.Teamname,
				Points1:   team.Points,
			}
			allscoreboards[team.MatchupID] = sb
		}

	}

	for _, ac := range allscoreboards {
		scoreboards = append(scoreboards, ac)
	}

	clear(allscoreboards)

	return scoreboards, nil
}

// Sends multiple API requests to get information for matchups, records, and scoreboard in order to correlate the data into one structure
func (c *Client) getFantasyInfo(league_id string, week int) ([]customTeamInfo, error) {
	var customInfo []customTeamInfo
	matchupWeek := week

	if week <= 0 {
		league, err := c.GetLeague(league_id)
		if err != nil {
			return customInfo, err
		}

		sportstate, err := c.GetSportState(league.Sport)
		if err != nil {
			return customInfo, err
		}

		matchupWeek = sportstate.Week
		if sportstate.SeasonType == "pre" {
			matchupWeek = 1
		}

	}

	// Get the matchups in the league
	matchups, err := c.GetMatchups(league_id, matchupWeek)
	if err != nil {
		return customInfo, err
	}

	// Get the rosters in the league
	rosters, err := c.GetRosters(league_id)
	if err != nil {
		return customInfo, err
	}

	// Get the users in the league
	users, err := c.GetLeagueUsers(league_id)
	if err != nil {
		return customInfo, err
	}

	// Combine the information from all of the data returned to have one full struct of user information
	// Loop through the users first
	for _, user := range users {
		newuser := customTeamInfo{
			DisplayName: user.DisplayName,
			OwnerID:     user.UserID,
			Week:        matchupWeek,
			Teamname:    user.Metadata.TeamName,
		}

		if user.Metadata.TeamName == "" {
			newuser.Teamname = "Team " + user.DisplayName
		}

		// Loop through rosters
		for _, roster := range rosters {
			if roster.OwnerID == newuser.OwnerID {
				newuser.Wins = roster.Settings.Wins
				newuser.Losses = roster.Settings.Losses
				newuser.RosterID = roster.RosterID
				break
			}
		}

		// Loop through matchups
		for _, matchup := range matchups {
			if matchup.RosterID == newuser.RosterID {
				newuser.MatchupID = matchup.MatchupID
				newuser.Points = matchup.Points
				break
			}
		}

		customInfo = append(customInfo, newuser)
	}

	return customInfo, nil
}
