package sleeper

import (
	"encoding/json"
	"fmt"
)

type League struct {
	TotalRosters int    `json:"total_rosters"`
	Status       string `json:"status"`
	Sport        string `json:"sport"`
	Shard        int    `json:"shard"`
	Settings     struct {
		ReserveAllowCov          int `json:"reserve_allow_cov"`
		ReserveSlots             int `json:"reserve_slots"`
		Leg                      int `json:"leg"`
		OffseasonAdds            int `json:"offseason_adds"`
		BenchLock                int `json:"bench_lock"`
		TradeReviewDays          int `json:"trade_review_days"`
		LeagueAverageMatch       int `json:"league_average_match"`
		WaiverType               int `json:"waiver_type"`
		MaxKeepers               int `json:"max_keepers"`
		Type                     int `json:"type"`
		PickTrading              int `json:"pick_trading"`
		DailyWaivers             int `json:"daily_waivers"`
		TaxiYears                int `json:"taxi_years"`
		TradeDeadline            int `json:"trade_deadline"`
		ReserveAllowSus          int `json:"reserve_allow_sus"`
		ReserveAllowOut          int `json:"reserve_allow_out"`
		PlayoffRoundType         int `json:"playoff_round_type"`
		WaiverDayOfWeek          int `json:"waiver_day_of_week"`
		TaxiAllowVets            int `json:"taxi_allow_vets"`
		ReserveAllowDnr          int `json:"reserve_allow_dnr"`
		CommissionerDirectInvite int `json:"commissioner_direct_invite"`
		ReserveAllowDoubtful     int `json:"reserve_allow_doubtful"`
		WaiverClearDays          int `json:"waiver_clear_days"`
		PlayoffWeekStart         int `json:"playoff_week_start"`
		TaxiSlots                int `json:"taxi_slots"`
		PlayoffType              int `json:"playoff_type"`
		DailyWaiversHour         int `json:"daily_waivers_hour"`
		NumTeams                 int `json:"num_teams"`
		PlayoffTeams             int `json:"playoff_teams"`
		PlayoffSeedType          int `json:"playoff_seed_type"`
		ReserveAllowNa           int `json:"reserve_allow_na"`
		DraftRounds              int `json:"draft_rounds"`
		TaxiDeadline             int `json:"taxi_deadline"`
		WaiverBidMin             int `json:"waiver_bid_min"`
		CapacityOverride         int `json:"capacity_override"`
		DisableAdds              int `json:"disable_adds"`
		WaiverBudget             int `json:"waiver_budget"`
	} `json:"settings"`
	SeasonType      string `json:"season_type"`
	Season          string `json:"season"`
	ScoringSettings struct {
		StFf         float32 `json:"st_ff"`
		PtsAllow713  float32 `json:"pts_allow_7_13"`
		DefStFf      float32 `json:"def_st_ff"`
		RecYd        float32 `json:"rec_yd"`
		FumRecTd     float32 `json:"fum_rec_td"`
		PtsAllow35P  float32 `json:"pts_allow_35p"`
		PtsAllow2834 float32 `json:"pts_allow_28_34"`
		Fum          float32 `json:"fum"`
		RushYd       float32 `json:"rush_yd"`
		PassTd       float32 `json:"pass_td"`
		BlkKick      float32 `json:"blk_kick"`
		PassYd       float32 `json:"pass_yd"`
		Safe         float32 `json:"safe"`
		DefTd        float32 `json:"def_td"`
		Fgm50P       float32 `json:"fgm_50p"`
		DefStTd      float32 `json:"def_st_td"`
		FumRec       float32 `json:"fum_rec"`
		Rush2Pt      float32 `json:"rush_2pt"`
		Xpm          float32 `json:"xpm"`
		PtsAllow2127 float32 `json:"pts_allow_21_27"`
		Fgm2029      float32 `json:"fgm_20_29"`
		PtsAllow16   float32 `json:"pts_allow_1_6"`
		FumLost      float32 `json:"fum_lost"`
		DefStFumRec  float32 `json:"def_st_fum_rec"`
		Int          float32 `json:"int"`
		Fgm019       float32 `json:"fgm_0_19"`
		PtsAllow1420 float32 `json:"pts_allow_14_20"`
		Rec          float32 `json:"rec"`
		Ff           float32 `json:"ff"`
		Fgmiss       float32 `json:"fgmiss"`
		StFumRec     float32 `json:"st_fum_rec"`
		Rec2Pt       float32 `json:"rec_2pt"`
		RushTd       float32 `json:"rush_td"`
		Xpmiss       float32 `json:"xpmiss"`
		Fgm3039      float32 `json:"fgm_30_39"`
		RecTd        float32 `json:"rec_td"`
		StTd         float32 `json:"st_td"`
		Pass2Pt      float32 `json:"pass_2pt"`
		PtsAllow0    float32 `json:"pts_allow_0"`
		PassInt      float32 `json:"pass_int"`
		Fgm4049      float32 `json:"fgm_40_49"`
		Sack         float32 `json:"sack"`
	} `json:"scoring_settings"`
	RosterPositions       []string    `json:"roster_positions"`
	PreviousLeagueID      string      `json:"previous_league_id"`
	Name                  string      `json:"name"`
	Metadata              interface{} `json:"metadata"`
	LoserBracketID        interface{} `json:"loser_bracket_id"`
	LeagueID              string      `json:"league_id"`
	LastTransactionID     interface{} `json:"last_transaction_id"`
	LastReadID            string      `json:"last_read_id"`
	LastPinnedMessageID   interface{} `json:"last_pinned_message_id"`
	LastMessageTime       int64       `json:"last_message_time"`
	LastMessageTextMap    interface{} `json:"last_message_text_map"`
	LastMessageID         string      `json:"last_message_id"`
	LastMessageAttachment interface{} `json:"last_message_attachment"`
	LastAuthorIsBot       bool        `json:"last_author_is_bot"`
	LastAuthorID          string      `json:"last_author_id"`
	LastAuthorDisplayName string      `json:"last_author_display_name"`
	LastAuthorAvatar      interface{} `json:"last_author_avatar"`
	GroupID               interface{} `json:"group_id"`
	DraftID               string      `json:"draft_id"`
	DisplayOrder          int         `json:"display_order"`
	CompanyID             interface{} `json:"company_id"`
	BracketID             interface{} `json:"bracket_id"`
	Avatar                string      `json:"avatar"`
}

type Roster struct {
	Taxi     interface{} `json:"taxi"`
	Starters []string    `json:"starters"`
	Settings struct {
		Wins             int `json:"wins"`
		WaiverPosition   int `json:"waiver_position"`
		WaiverBudgetUsed int `json:"waiver_budget_used"`
		TotalMoves       int `json:"total_moves"`
		Ties             int `json:"ties"`
		Losses           int `json:"losses"`
		Fpts             int `json:"fpts"`
	} `json:"settings"`
	RosterID  int         `json:"roster_id"`
	Reserve   interface{} `json:"reserve"`
	Players   []string    `json:"players"`
	PlayerMap interface{} `json:"player_map"`
	OwnerID   string      `json:"owner_id"`
	Metadata  struct {
		RestrictPnScoringStartersOnly string `json:"restrict_pn_scoring_starters_only"`
		AllowPnScoring                string `json:"allow_pn_scoring"`
		AllowPnPlayerInjuryStatus     string `json:"allow_pn_player_injury_status"`
		AllowPnInactiveStarters       string `json:"allow_pn_inactive_starters"`
	} `json:"metadata"`
	LeagueID string      `json:"league_id"`
	Keepers  interface{} `json:"keepers"`
	CoOwners interface{} `json:"co_owners"`
}

type LeagueUser struct {
	UserID      string      `json:"user_id"`
	Settings    interface{} `json:"settings"`
	LeagueID    string      `json:"league_id"`
	IsOwner     interface{} `json:"is_owner"`
	IsBot       bool        `json:"is_bot"`
	DisplayName string      `json:"display_name"`
	Avatar      string      `json:"avatar"`
	Metadata    []struct {
		UserMessagePn           string `json:"user_message_pn"`
		TransactionWaiver       string `json:"transaction_waiver"`
		TransactionTrade        string `json:"transaction_trade"`
		TransactionFreeAgent    string `json:"transaction_free_agent"`
		TransactionCommissioner string `json:"transaction_commissioner"`
		TradeBlockPn            string `json:"trade_block_pn"`
		TeamNameUpdate          string `json:"team_name_update"`
		TeamName                string `json:"team_name"`
		PlayerNicknameUpdate    string `json:"player_nickname_update"`
		PlayerLikePn            string `json:"player_like_pn"`
		MentionPn               string `json:"mention_pn"`
		MascotMessage           string `json:"mascot_message"`
		Avatar                  string `json:"avatar"`
		AllowSms                string `json:"allow_sms"`
		AllowPn                 string `json:"allow_pn"`
	} `json:"metadata,omitempty"`
}

type Matchup struct {
	StartersPoints []float32          `json:"starters_points"`
	Starters       []string           `json:"starters"`
	RosterID       int                `json:"roster_id"`
	Players        []string           `json:"players"`
	MatchupID      int                `json:"matchup_id"`
	Points         float32            `json:"points"`
	CustomPoints   float32            `json:"custom_points"`
	PlayersPoints  map[string]float32 `json:"players_points"`
}

type PlayoffRound struct {
	L      int `json:"l"`
	M      int `json:"m"`
	P      int `json:"p,omitempty"`
	R      int `json:"r"`
	T1     int `json:"t1"`
	T2     int `json:"t2"`
	W      int `json:"w"`
	T1From struct {
		L int `json:"l,omitempty"`
		W int `json:"w,omitempty"`
	} `json:"t1_from,omitempty"`
	T2From struct {
		L int `json:"l,omitempty"`
		W int `json:"w,omitempty"`
	} `json:"t2_from,omitempty"`
}

type Transaction struct {
	WaiverBudget []struct {
		Sender   int `json:"sender"`
		Receiver int `json:"receiver"`
		Amount   int `json:"amount"`
	} `json:"waiver_budget"`
	Type          string         `json:"type"`
	TransactionID string         `json:"transaction_id"`
	StatusUpdated int64          `json:"status_updated"`
	Status        string         `json:"status"`
	Settings      interface{}    `json:"settings"`
	RosterIds     []int          `json:"roster_ids"`
	Metadata      interface{}    `json:"metadata"`
	Leg           int            `json:"leg"`
	Drops         map[string]int `json:"drops"`
	DraftPicks    []struct {
		Season          string `json:"season"`
		Round           int    `json:"round"`
		RosterID        int    `json:"roster_id"`
		PreviousOwnerID int    `json:"previous_owner_id"`
		OwnerID         int    `json:"owner_id"`
	} `json:"draft_picks"`
	Creator      string         `json:"creator"`
	Created      int64          `json:"created"`
	ConsenterIds []int          `json:"consenter_ids"`
	Adds         map[string]int `json:"adds"`
}

type SportState struct {
	Week               int    `json:"week"`
	SeasonType         string `json:"season_type"`
	SeasonStartDate    string `json:"season_start_date"`
	Season             string `json:"season"`
	PreviousSeason     string `json:"previous_season"`
	Leg                int    `json:"leg"`
	LeagueSeason       string `json:"league_season"`
	LeagueCreateSeason string `json:"league_create_season"`
	DisplayWeek        int    `json:"display_week"`
}

// Get all leagues for a specific user, sport, and season.
// (GET `https://api.sleeper.app/v1/user/<user_id>/leagues/<sport>/<season>`)
func (c *Client) GetAllLeagesForUser(user_id string, sport string, season string) ([]League, error) {
	leagues := []League{}

	url := fmt.Sprintf("%s/user/%s/leagues/%s/%s", c.sleeperURL, user_id, sport, season)

	data, err := c.getRequest(url)
	if err != nil {
		return leagues, err
	}

	err = json.Unmarshal(data, &leagues)

	return leagues, err
}

// Get a specific league by the league_id.
// (GET `https://api.sleeper.app/v1/league/<league_id>`)
func (c *Client) GetLeague(league_id string) (League, error) {
	league := League{}

	url := fmt.Sprintf("%s/league/%s", c.sleeperURL, league_id)

	data, err := c.getRequest(url)
	if err != nil {
		return league, err
	}

	err = json.Unmarshal(data, &league)

	return league, err
}

// Get all rosters in a league.
// (GET `https://api.sleeper.app/v1/league/<league_id>/rosters`)
func (c *Client) GetRosters(league_id string) ([]Roster, error) {
	rosters := []Roster{}

	url := fmt.Sprintf("%s/league/%s/rosters", c.sleeperURL, league_id)

	data, err := c.getRequest(url)
	if err != nil {
		return rosters, err
	}

	err = json.Unmarshal(data, &rosters)
	return rosters, err
}

// Get all users in a league.
// (GET `https://api.sleeper.app/v1/league/<league_id>/users`)
func (c *Client) GetLeagueUsers(league_id string) ([]LeagueUser, error) {
	leagueUsers := []LeagueUser{}

	url := fmt.Sprintf("%s/league/%s/users", c.sleeperURL, league_id)

	data, err := c.getRequest(url)
	if err != nil {
		return leagueUsers, err
	}

	err = json.Unmarshal(data, &leagueUsers)

	return leagueUsers, err
}

// Get all matchups in a league for a given week. Each object in the list represents one team. The two teams with the same matchup_id match up against each other.
// (GET `https://api.sleeper.app/v1/league/<league_id>/matchups/<week>`)
func (c *Client) GetMatchups(league_id string, week int) ([]Matchup, error) {
	matchups := []Matchup{}

	url := fmt.Sprintf("%s/league/%s/matchups/%d", c.sleeperURL, league_id, week)

	data, err := c.getRequest(url)
	if err != nil {
		return matchups, err
	}

	err = json.Unmarshal(data, &matchups)
	return matchups, err
}

// Get the playoff winners bracket for a league for 4, 6, and 8 team playoffs.
// (GET `https://api.sleeper.app/v1/league/<league_id>/winners_bracket`)
func (c *Client) GetPlayoffsWinnersBracket(league_id string) ([]PlayoffRound, error) {
	playoffRounds := []PlayoffRound{}

	url := fmt.Sprintf("%s/league/%s/winners_bracket", c.sleeperURL, league_id)

	data, err := c.getRequest(url)
	if err != nil {
		return playoffRounds, err
	}

	err = json.Unmarshal(data, &playoffRounds)
	return playoffRounds, err
}

// Get the playoff losers bracket for a league for 4, 6, and 8 team playoffs.
// (GET `https://api.sleeper.app/v1/league/<league_id>/losers_bracket`)
func (c *Client) GetPlayoffsLosersBracket(league_id string) ([]PlayoffRound, error) {
	playoffRounds := []PlayoffRound{}

	url := fmt.Sprintf("%s/league/%s/losers_bracket", c.sleeperURL, league_id)

	data, err := c.getRequest(url)
	if err != nil {
		return playoffRounds, err
	}

	err = json.Unmarshal(data, &playoffRounds)
	return playoffRounds, err
}

// Get all free agent transactions, waivers and trades.
// (GET `https://api.sleeper.app/v1/league/<league_id>/transactions/<round>`)
func (c *Client) GetTransactions(league_id string, round int) ([]Transaction, error) {
	transactions := []Transaction{}

	url := fmt.Sprintf("%s/league/%s/transactions/%d", c.sleeperURL, league_id, round)

	data, err := c.getRequest(url)
	if err != nil {
		return transactions, err
	}

	err = json.Unmarshal(data, &transactions)
	return transactions, err
}

// Get all traded picks in a league, including future picks.
// (GET `https://api.sleeper.app/v1/league/<league_id>/traded_picks`)
func (c *Client) GetLeagueTradedPicks(league_id string) ([]TradedPick, error) {
	tradedPicks := []TradedPick{}

	url := fmt.Sprintf("%s/league/%s/traded_picks", c.sleeperURL, league_id)

	data, err := c.getRequest(url)
	if err != nil {
		return tradedPicks, err
	}

	err = json.Unmarshal(data, &tradedPicks)

	return tradedPicks, err
}

// Get information about the current state for any sport.
// (GET `https://api.sleeper.app/v1/state/<sport>`)
func (c *Client) GetSportState(sport string) (SportState, error) {
	sportstate := SportState{}

	url := fmt.Sprintf("%s/state/%s", c.sleeperURL, sport)

	data, err := c.getRequest(url)
	if err != nil {
		return sportstate, err
	}

	err = json.Unmarshal(data, &sportstate)
	return sportstate, err
}
