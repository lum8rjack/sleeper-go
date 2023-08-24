package sleeper

import (
	"encoding/json"
	"fmt"
)

type League struct {
	Avatar                string      `json:"avatar"`
	BracketID             interface{} `json:"bracket_id"`
	CompanyID             interface{} `json:"company_id"`
	DisplayOrder          int         `json:"display_order"`
	DraftID               string      `json:"draft_id"`
	GroupID               interface{} `json:"group_id"`
	LastAuthorAvatar      interface{} `json:"last_author_avatar"`
	LastAuthorDisplayName string      `json:"last_author_display_name"`
	LastAuthorID          string      `json:"last_author_id"`
	LastAuthorIsBot       bool        `json:"last_author_is_bot"`
	LastMessageAttachment interface{} `json:"last_message_attachment"`
	LastMessageID         string      `json:"last_message_id"`
	LastMessageTextMap    interface{} `json:"last_message_text_map"`
	LastMessageTime       int64       `json:"last_message_time"`
	LastPinnedMessageID   interface{} `json:"last_pinned_message_id"`
	LastReadID            string      `json:"last_read_id"`
	LastTransactionID     interface{} `json:"last_transaction_id"`
	LeagueID              string      `json:"league_id"`
	LoserBracketID        interface{} `json:"loser_bracket_id"`
	Metadata              interface{} `json:"metadata"`
	Name                  string      `json:"name"`
	PreviousLeagueID      string      `json:"previous_league_id"`
	RosterPositions       []string    `json:"roster_positions"`
	ScoringSettings       struct {
		BlkKick      float32 `json:"blk_kick"`
		DefStFf      float32 `json:"def_st_ff"`
		DefStFumRec  float32 `json:"def_st_fum_rec"`
		DefStTd      float32 `json:"def_st_td"`
		DefTd        float32 `json:"def_td"`
		Ff           float32 `json:"ff"`
		Fgm019       float32 `json:"fgm_0_19"`
		Fgm2029      float32 `json:"fgm_20_29"`
		Fgm3039      float32 `json:"fgm_30_39"`
		Fgm4049      float32 `json:"fgm_40_49"`
		Fgm50P       float32 `json:"fgm_50p"`
		Fgmiss       float32 `json:"fgmiss"`
		Fum          float32 `json:"fum"`
		FumLost      float32 `json:"fum_lost"`
		FumRec       float32 `json:"fum_rec"`
		FumRecTd     float32 `json:"fum_rec_td"`
		Int          float32 `json:"int"`
		Pass2Pt      float32 `json:"pass_2pt"`
		PassInt      float32 `json:"pass_int"`
		PassTd       float32 `json:"pass_td"`
		PassYd       float32 `json:"pass_yd"`
		PtsAllow0    float32 `json:"pts_allow_0"`
		PtsAllow16   float32 `json:"pts_allow_1_6"`
		PtsAllow713  float32 `json:"pts_allow_7_13"`
		PtsAllow1420 float32 `json:"pts_allow_14_20"`
		PtsAllow2127 float32 `json:"pts_allow_21_27"`
		PtsAllow2834 float32 `json:"pts_allow_28_34"`
		PtsAllow35P  float32 `json:"pts_allow_35p"`
		Rec          float32 `json:"rec"`
		Rec2Pt       float32 `json:"rec_2pt"`
		RecTd        float32 `json:"rec_td"`
		RecYd        float32 `json:"rec_yd"`
		Rush2Pt      float32 `json:"rush_2pt"`
		RushTd       float32 `json:"rush_td"`
		RushYd       float32 `json:"rush_yd"`
		Sack         float32 `json:"sack"`
		Safe         float32 `json:"safe"`
		StFf         float32 `json:"st_ff"`
		StFumRec     float32 `json:"st_fum_rec"`
		StTd         float32 `json:"st_td"`
		Xpm          float32 `json:"xpm"`
		Xpmiss       float32 `json:"xpmiss"`
	} `json:"scoring_settings"`
	Season     string `json:"season"`
	SeasonType string `json:"season_type"`
	Settings   struct {
		BenchLock                int `json:"bench_lock"`
		CapacityOverride         int `json:"capacity_override"`
		CommissionerDirectInvite int `json:"commissioner_direct_invite"`
		DailyWaivers             int `json:"daily_waivers"`
		DailyWaiversHour         int `json:"daily_waivers_hour"`
		DisableAdds              int `json:"disable_adds"`
		DraftRounds              int `json:"draft_rounds"`
		LeagueAverageMatch       int `json:"league_average_match"`
		Leg                      int `json:"leg"`
		MaxKeepers               int `json:"max_keepers"`
		NumTeams                 int `json:"num_teams"`
		OffseasonAdds            int `json:"offseason_adds"`
		PickTrading              int `json:"pick_trading"`
		PlayoffRoundType         int `json:"playoff_round_type"`
		PlayoffSeedType          int `json:"playoff_seed_type"`
		PlayoffTeams             int `json:"playoff_teams"`
		PlayoffType              int `json:"playoff_type"`
		PlayoffWeekStart         int `json:"playoff_week_start"`
		ReserveAllowCov          int `json:"reserve_allow_cov"`
		ReserveAllowDnr          int `json:"reserve_allow_dnr"`
		ReserveAllowDoubtful     int `json:"reserve_allow_doubtful"`
		ReserveAllowNa           int `json:"reserve_allow_na"`
		ReserveAllowOut          int `json:"reserve_allow_out"`
		ReserveAllowSus          int `json:"reserve_allow_sus"`
		ReserveSlots             int `json:"reserve_slots"`
		TaxiAllowVets            int `json:"taxi_allow_vets"`
		TaxiDeadline             int `json:"taxi_deadline"`
		TaxiSlots                int `json:"taxi_slots"`
		TaxiYears                int `json:"taxi_years"`
		TradeDeadline            int `json:"trade_deadline"`
		TradeReviewDays          int `json:"trade_review_days"`
		Type                     int `json:"type"`
		WaiverBidMin             int `json:"waiver_bid_min"`
		WaiverBudget             int `json:"waiver_budget"`
		WaiverClearDays          int `json:"waiver_clear_days"`
		WaiverDayOfWeek          int `json:"waiver_day_of_week"`
		WaiverType               int `json:"waiver_type"`
	} `json:"settings"`
	Shard        int    `json:"shard"`
	Sport        string `json:"sport"`
	Status       string `json:"status"`
	TotalRosters int    `json:"total_rosters"`
}

type Roster struct {
	CoOwners interface{} `json:"co_owners"`
	Keepers  interface{} `json:"keepers"`
	LeagueID string      `json:"league_id"`
	Metadata struct {
		AllowPnInactiveStarters       string `json:"allow_pn_inactive_starters"`
		AllowPnPlayerInjuryStatus     string `json:"allow_pn_player_injury_status"`
		AllowPnScoring                string `json:"allow_pn_scoring"`
		RestrictPnScoringStartersOnly string `json:"restrict_pn_scoring_starters_only"`
	} `json:"metadata"`
	OwnerID   string      `json:"owner_id"`
	PlayerMap interface{} `json:"player_map"`
	Players   []string    `json:"players"`
	Reserve   interface{} `json:"reserve"`
	RosterID  int         `json:"roster_id"`
	Settings  struct {
		Fpts             int `json:"fpts"`
		Losses           int `json:"losses"`
		Ties             int `json:"ties"`
		TotalMoves       int `json:"total_moves"`
		WaiverBudgetUsed int `json:"waiver_budget_used"`
		WaiverPosition   int `json:"waiver_position"`
		Wins             int `json:"wins"`
	} `json:"settings"`
	Starters []string    `json:"starters"`
	Taxi     interface{} `json:"taxi"`
}

type LeagueUser struct {
	Avatar      string      `json:"avatar"`
	DisplayName string      `json:"display_name"`
	IsBot       bool        `json:"is_bot"`
	IsOwner     interface{} `json:"is_owner"`
	LeagueID    string      `json:"league_id"`
	Metadata    struct {
		AllowPn                 string `json:"allow_pn"`
		AllowSms                string `json:"allow_sms"`
		Avatar                  string `json:"avatar"`
		MascotMessage           string `json:"mascot_message"`
		MentionPn               string `json:"mention_pn"`
		PlayerLikePn            string `json:"player_like_pn"`
		PlayerNicknameUpdate    string `json:"player_nickname_update"`
		TeamName                string `json:"team_name"`
		TeamNameUpdate          string `json:"team_name_update"`
		TradeBlockPn            string `json:"trade_block_pn"`
		TransactionCommissioner string `json:"transaction_commissioner"`
		TransactionFreeAgent    string `json:"transaction_free_agent"`
		TransactionTrade        string `json:"transaction_trade"`
		TransactionWaiver       string `json:"transaction_waiver"`
		UserMessagePn           string `json:"user_message_pn"`
	} `json:"metadata,omitempty"`
	Settings interface{} `json:"settings"`
	UserID   string      `json:"user_id"`
}

type Matchup struct {
	CustomPoints   float32            `json:"custom_points"`
	MatchupID      int                `json:"matchup_id"`
	Players        []string           `json:"players"`
	PlayersPoints  map[string]float32 `json:"players_points"`
	Points         float32            `json:"points"`
	RosterID       int                `json:"roster_id"`
	Starters       []string           `json:"starters"`
	StartersPoints []float32          `json:"starters_points"`
}

type PlayoffRound struct {
	L      int `json:"l"`
	M      int `json:"m"`
	P      int `json:"p,omitempty"`
	R      int `json:"r"`
	T1     int `json:"t1"`
	T1From struct {
		L int `json:"l,omitempty"`
		W int `json:"w,omitempty"`
	} `json:"t1_from,omitempty"`
	T2     int `json:"t2"`
	T2From struct {
		L int `json:"l,omitempty"`
		W int `json:"w,omitempty"`
	} `json:"t2_from,omitempty"`
	W int `json:"w"`
}

type Transaction struct {
	Adds         map[string]int `json:"adds"`
	ConsenterIds []int          `json:"consenter_ids"`
	Created      int64          `json:"created"`
	Creator      string         `json:"creator"`
	DraftPicks   []struct {
		OwnerID         int    `json:"owner_id"`
		PreviousOwnerID int    `json:"previous_owner_id"`
		RosterID        int    `json:"roster_id"`
		Round           int    `json:"round"`
		Season          string `json:"season"`
	} `json:"draft_picks"`
	Drops         map[string]int `json:"drops"`
	Leg           int            `json:"leg"`
	Metadata      interface{}    `json:"metadata"`
	RosterIds     []int          `json:"roster_ids"`
	Settings      interface{}    `json:"settings"`
	Status        string         `json:"status"`
	StatusUpdated int64          `json:"status_updated"`
	TransactionID string         `json:"transaction_id"`
	Type          string         `json:"type"`
	WaiverBudget  []struct {
		Amount   int `json:"amount"`
		Receiver int `json:"receiver"`
		Sender   int `json:"sender"`
	} `json:"waiver_budget"`
}

type SportState struct {
	DisplayWeek        int    `json:"display_week"`
	LeagueCreateSeason string `json:"league_create_season"`
	LeagueSeason       string `json:"league_season"`
	Leg                int    `json:"leg"`
	PreviousSeason     string `json:"previous_season"`
	Season             string `json:"season"`
	SeasonStartDate    string `json:"season_start_date"`
	SeasonType         string `json:"season_type"`
	Week               int    `json:"week"`
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
