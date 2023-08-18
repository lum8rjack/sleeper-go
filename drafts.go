package sleeper

import (
	"encoding/json"
	"fmt"
)

type Draft struct {
	Type           string      `json:"type"`
	Status         string      `json:"status"`
	StartTime      int64       `json:"start_time"`
	Sport          string      `json:"sport"`
	SlotToRosterID interface{} `json:"slot_to_roster_id"`
	Settings       struct {
		Teams                 int `json:"teams"`
		SlotsWr               int `json:"slots_wr"`
		SlotsTe               int `json:"slots_te"`
		SlotsRb               int `json:"slots_rb"`
		SlotsQb               int `json:"slots_qb"`
		SlotsK                int `json:"slots_k"`
		SlotsFlex             int `json:"slots_flex"`
		SlotsDef              int `json:"slots_def"`
		SlotsBn               int `json:"slots_bn"`
		Rounds                int `json:"rounds"`
		ReversalRound         int `json:"reversal_round"`
		PlayerType            int `json:"player_type"`
		PickTimer             int `json:"pick_timer"`
		NominationTimer       int `json:"nomination_timer"`
		EnforcePositionLimits int `json:"enforce_position_limits"`
		CPUAutopick           int `json:"cpu_autopick"`
		Autostart             int `json:"autostart"`
		AutopauseStartTime    int `json:"autopause_start_time"`
		AutopauseEndTime      int `json:"autopause_end_time"`
		AutopauseEnabled      int `json:"autopause_enabled"`
		AlphaSort             int `json:"alpha_sort"`
	} `json:"settings"`
	SeasonType string `json:"season_type"`
	Season     string `json:"season"`
	Metadata   struct {
		ScoringType string `json:"scoring_type"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"metadata"`
	LeagueID        string      `json:"league_id"`
	LastPicked      int64       `json:"last_picked"`
	LastMessageTime int64       `json:"last_message_time"`
	LastMessageID   string      `json:"last_message_id"`
	DraftOrder      interface{} `json:"draft_order"`
	DraftID         string      `json:"draft_id"`
	Creators        interface{} `json:"creators"`
	Created         int64       `json:"created"`
}

// Get all drafts by a user
// GET https://api.sleeper.app/v1/user/<user_id>/drafts/<sport>/<season>
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

// Get all drafts for a league
// GET https://api.sleeper.app/v1/league/<league_id>/drafts
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

// Get a specific draft
// GET https://api.sleeper.app/v1/draft/<draft_id>
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
