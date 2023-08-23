package sleeper

import (
	"encoding/json"
	"fmt"
)

type Draft struct {
	Created         int64       `json:"created"`
	Creators        interface{} `json:"creators"`
	DraftID         string      `json:"draft_id"`
	DraftOrder      interface{} `json:"draft_order"`
	LastMessageTime int64       `json:"last_message_time"`
	LastMessageID   string      `json:"last_message_id"`
	LastPicked      int64       `json:"last_picked"`
	LeagueID        string      `json:"league_id"`
	Metadata        struct {
		Description string `json:"description"`
		Name        string `json:"name"`
		ScoringType string `json:"scoring_type"`
	} `json:"metadata"`
	Settings struct {
		AlphaSort             int `json:"alpha_sort"`
		AutopauseEnabled      int `json:"autopause_enabled"`
		AutopauseEndTime      int `json:"autopause_end_time"`
		AutopauseStartTime    int `json:"autopause_start_time"`
		Autostart             int `json:"autostart"`
		CPUAutopick           int `json:"cpu_autopick"`
		EnforcePositionLimits int `json:"enforce_position_limits"`
		NominationTimer       int `json:"nomination_timer"`
		PickTimer             int `json:"pick_timer"`
		PlayerType            int `json:"player_type"`
		ReversalRound         int `json:"reversal_round"`
		Rounds                int `json:"rounds"`
		SlotsBn               int `json:"slots_bn"`
		SlotsDef              int `json:"slots_def"`
		SlotsFlex             int `json:"slots_flex"`
		SlotsK                int `json:"slots_k"`
		SlotsQb               int `json:"slots_qb"`
		SlotsRb               int `json:"slots_rb"`
		SlotsTe               int `json:"slots_te"`
		SlotsWr               int `json:"slots_wr"`
		Teams                 int `json:"teams"`
	} `json:"settings"`
	Season         string      `json:"season"`
	SeasonType     string      `json:"season_type"`
	SlotToRosterID interface{} `json:"slot_to_roster_id"`
	Sport          string      `json:"sport"`
	StartTime      int64       `json:"start_time"`
	Status         string      `json:"status"`
	Type           string      `json:"type"`
}

type DraftPlayer struct {
	Round    int    `json:"round"`
	RosterID int    `json:"roster_id"`
	PlayerID string `json:"player_id"`
	PickedBy string `json:"picked_by"`
	PickNo   int    `json:"pick_no"`
	Metadata struct {
		YearsExp     string `json:"years_exp"`
		Team         string `json:"team"`
		Status       string `json:"status"`
		Sport        string `json:"sport"`
		Position     string `json:"position"`
		PlayerID     string `json:"player_id"`
		Number       string `json:"number"`
		NewsUpdated  string `json:"news_updated"`
		LastName     string `json:"last_name"`
		InjuryStatus string `json:"injury_status"`
		FirstName    string `json:"first_name"`
	} `json:"metadata"`
	IsKeeper  interface{} `json:"is_keeper"`
	DraftSlot int         `json:"draft_slot"`
	DraftID   string      `json:"draft_id"`
}

type TradedPick struct {
	OwnerID         int    `json:"owner_id"`
	PreviousOwnerID int    `json:"previous_owner_id"`
	Round           int    `json:"round"`
	RosterID        int    `json:"roster_id"`
	Season          string `json:"season"`
}

// Get all drafts by a user.
// (GET `https://api.sleeper.app/v1/user/<user_id>/drafts/<sport>/<season>`)
func (c *Client) GetDraftsForUser(user_id string, sport string, season int) ([]Draft, error) {
	drafts := []Draft{}

	url := fmt.Sprintf("%s/user/%s/drafts/%s/%d", c.sleeperURL, user_id, sport, season)

	data, err := c.getRequest(url)
	if err != nil {
		return drafts, err
	}

	err = json.Unmarshal(data, &drafts)

	return drafts, err
}

// Get all drafts for a league.
// (GET `https://api.sleeper.app/v1/league/<league_id>/drafts`)
func (c *Client) GetDraftsForLeague(league_id string) ([]Draft, error) {
	drafts := []Draft{}

	url := fmt.Sprintf("%s/league/%s/drafts", c.sleeperURL, league_id)

	data, err := c.getRequest(url)
	if err != nil {
		return drafts, err
	}

	err = json.Unmarshal(data, &drafts)

	return drafts, err
}

// Get a specific draft.
// (GET `https://api.sleeper.app/v1/draft/<draft_id>`)
func (c *Client) GetDraft(draft_id string) (Draft, error) {
	draft := Draft{}

	url := fmt.Sprintf("%s/draft/%s", c.sleeperURL, draft_id)

	data, err := c.getRequest(url)
	if err != nil {
		return draft, err
	}

	err = json.Unmarshal(data, &draft)

	return draft, err
}

// Get all picks in a draft.
// (GET `https://api.sleeper.app/v1/draft/<draft_id>/picks`)
func (c *Client) GetAllDraftPicks(draft_id string) ([]DraftPlayer, error) {
	draftPlayers := []DraftPlayer{}

	url := fmt.Sprintf("%s/draft/%s/picks", c.sleeperURL, draft_id)

	data, err := c.getRequest(url)
	if err != nil {
		return draftPlayers, err
	}

	err = json.Unmarshal(data, &draftPlayers)

	return draftPlayers, err
}

// Get all traded picks in a draft.
// (GET `https://api.sleeper.app/v1/draft/<draft_id>/traded_picks`)
func (c *Client) GetDraftTradedPicks(draft_id string) ([]TradedPick, error) {
	tradedPicks := []TradedPick{}

	url := fmt.Sprintf("%s/draft/%s/traded_picks", c.sleeperURL, draft_id)

	data, err := c.getRequest(url)
	if err != nil {
		return tradedPicks, err
	}

	err = json.Unmarshal(data, &tradedPicks)

	return tradedPicks, err
}
