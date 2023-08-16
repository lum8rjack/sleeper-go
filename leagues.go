package sleeper

import (
	"encoding/json"
	"fmt"
)

const (
	leagueEndpoint string = "/league"
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
		StFf         float64 `json:"st_ff"`
		PtsAllow713  float64 `json:"pts_allow_7_13"`
		DefStFf      float64 `json:"def_st_ff"`
		RecYd        float64 `json:"rec_yd"`
		FumRecTd     float64 `json:"fum_rec_td"`
		PtsAllow35P  int     `json:"pts_allow_35p"`
		PtsAllow2834 int     `json:"pts_allow_28_34"`
		Fum          float64 `json:"fum"`
		RushYd       float64 `json:"rush_yd"`
		PassTd       float64 `json:"pass_td"`
		BlkKick      float64 `json:"blk_kick"`
		PassYd       float64 `json:"pass_yd"`
		Safe         float64 `json:"safe"`
		DefTd        float64 `json:"def_td"`
		Fgm50P       float64 `json:"fgm_50p"`
		DefStTd      float64 `json:"def_st_td"`
		FumRec       float64 `json:"fum_rec"`
		Rush2Pt      float64 `json:"rush_2pt"`
		Xpm          float64 `json:"xpm"`
		PtsAllow2127 float64 `json:"pts_allow_21_27"`
		Fgm2029      float64 `json:"fgm_20_29"`
		PtsAllow16   float64 `json:"pts_allow_1_6"`
		FumLost      int     `json:"fum_lost"`
		DefStFumRec  float64 `json:"def_st_fum_rec"`
		Int          float64 `json:"int"`
		Fgm019       float64 `json:"fgm_0_19"`
		PtsAllow1420 float64 `json:"pts_allow_14_20"`
		Rec          float64 `json:"rec"`
		Ff           float64 `json:"ff"`
		Fgmiss       int     `json:"fgmiss"`
		StFumRec     float64 `json:"st_fum_rec"`
		Rec2Pt       float64 `json:"rec_2pt"`
		RushTd       float64 `json:"rush_td"`
		Xpmiss       int     `json:"xpmiss"`
		Fgm3039      float64 `json:"fgm_30_39"`
		RecTd        float64 `json:"rec_td"`
		StTd         float64 `json:"st_td"`
		Pass2Pt      float64 `json:"pass_2pt"`
		PtsAllow0    float64 `json:"pts_allow_0"`
		PassInt      int     `json:"pass_int"`
		Fgm4049      float64 `json:"fgm_40_49"`
		Sack         float64 `json:"sack"`
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

// Get all leagues for a specific user, sport, and season
func (c *Client) GetAllLeagesForUser(user_id string, sport string, season string) ([]League, error) {
	leagues := []League{}

	//https://api.sleeper.app/v1/user/<user_id>/leagues/<sport>/<season>
	url := fmt.Sprintf("%s/user/%s/leagues/%s/%s", c.sleeperURL, user_id, sport, season)

	data, err := c.getRequest(url)
	if err != nil {
		return leagues, err
	}

	err = json.Unmarshal(data, &leagues)

	return leagues, err
}

// Get a specific league by the league_id
func (c *Client) GetLeague(league_id string) (League, error) {
	league := League{}

	url := fmt.Sprintf("%s%s/%s", c.sleeperURL, leagueEndpoint, league_id)

	data, err := c.getRequest(url)
	if err != nil {
		return league, err
	}

	err = json.Unmarshal(data, &league)

	return league, err
}
