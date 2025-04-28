package sleeper

import (
	"encoding/json"
	"fmt"
)

type PlayerResearch struct {
	Owned   float64 `json:"owned"`
	Started float64 `json:"started"`
}

type PlayerStats struct {
	Date  any `json:"date"`
	Stats struct {
		BonusPassCmp25 float64 `json:"bonus_pass_cmp_25"`
		BonusPassYd300 float64 `json:"bonus_pass_yd_300"`
		BonusRecWr     float64 `json:"bonus_rec_wr"`
		BonusSack2P    float64 `json:"bonus_sack_2p"`
		BonusTkl10P    float64 `json:"bonus_tkl_10p"`
		CmpPct         float64 `json:"cmp_pct"`
		DefSnp         float64 `json:"def_snp"`
		Fga            float64 `json:"fga"`
		Fgm2029        float64 `json:"fgm_20_29"`
		Fgm3039        float64 `json:"fgm_30_39"`
		Fgm4049        float64 `json:"fgm_40_49"`
		Fgm50P         float64 `json:"fgm_50p"`
		Fgm            float64 `json:"fgm"`
		Fgmiss3039     float64 `json:"fgmiss_30_39"`
		Fgmiss4049     float64 `json:"fgmiss_40_49"`
		Fgmiss50P      float64 `json:"fgmiss_50p"`
		Fgmiss         float64 `json:"fgmiss"`
		FgmLng         float64 `json:"fgm_lng"`
		FgmPct         float64 `json:"fgm_pct"`
		FgmYds         float64 `json:"fgm_yds"`
		FgmYdsOver30   float64 `json:"fgm_yds_over_30"`
		Fum            float64 `json:"fum"`
		FumLost        float64 `json:"fum_lost"`
		GmsActive      float64 `json:"gms_active"`
		Gp             float64 `json:"gp"`
		Gs             float64 `json:"gs"`
		IdpFf          float64 `json:"idp_ff"`
		IdpFumRec      float64 `json:"idp_fum_rec"`
		IdpFumRetYd    float64 `json:"idp_fum_ret_yd"`
		IdpInt         float64 `json:"idp_int"`
		IdpIntRetYd    float64 `json:"idp_int_ret_yd"`
		IdpPassDef     float64 `json:"idp_pass_def"`
		IdpQbHit       float64 `json:"idp_qb_hit"`
		IdpSack        float64 `json:"idp_sack"`
		IdpSackYd      float64 `json:"idp_sack_yd"`
		IdpTklAst      float64 `json:"idp_tkl_ast"`
		IdpTkl         float64 `json:"idp_tkl"`
		IdpTklLoss     float64 `json:"idp_tkl_loss"`
		IdpTklSolo     float64 `json:"idp_tkl_solo"`
		OffSnp         float64 `json:"off_snp"`
		PassAirYd      float64 `json:"pass_air_yd"`
		PassAtt        float64 `json:"pass_att"`
		PassCmp40P     float64 `json:"pass_cmp_40p"`
		PassCmp        float64 `json:"pass_cmp"`
		PassFd         float64 `json:"pass_fd"`
		PassInc        float64 `json:"pass_inc"`
		PassInt        float64 `json:"pass_int"`
		PassIntTd      float64 `json:"pass_int_td"`
		PassLng        float64 `json:"pass_lng"`
		PassRtg        float64 `json:"pass_rtg"`
		PassRushYd     float64 `json:"pass_rush_yd"`
		PassRzAtt      float64 `json:"pass_rz_att"`
		PassSack       float64 `json:"pass_sack"`
		PassSackYds    float64 `json:"pass_sack_yds"`
		PassTd40P      float64 `json:"pass_td_40p"`
		PassTd         float64 `json:"pass_td"`
		PassTdLng      float64 `json:"pass_td_lng"`
		PassYd         float64 `json:"pass_yd"`
		PassYpa        float64 `json:"pass_ypa"`
		PassYpc        float64 `json:"pass_ypc"`
		Penalty        float64 `json:"penalty"`
		PenaltyYd      float64 `json:"penalty_yd"`
		PosRankHalfPpr int     `json:"pos_rank_half_ppr"`
		PosRankPpr     int     `json:"pos_rank_ppr"`
		PosRankStd     int     `json:"pos_rank_std"`
		PtsHalfPpr     float64 `json:"pts_half_ppr"`
		PtsPpr         float64 `json:"pts_ppr"`
		PtsStd         float64 `json:"pts_std"`
		PuntIn20       float64 `json:"punt_in_20"`
		PuntNetYd      float64 `json:"punt_net_yd"`
		Punts          float64 `json:"punts"`
		PuntTb         float64 `json:"punt_tb"`
		PuntYds        float64 `json:"punt_yds"`
		RankHalfPpr    int     `json:"rank_half_ppr"`
		RankPpr        int     `json:"rank_ppr"`
		RankStd        int     `json:"rank_std"`
		Rec04          float64 `json:"rec_0_4"`
		Rec1019        float64 `json:"rec_10_19"`
		Rec2029        float64 `json:"rec_20_29"`
		Rec3039        float64 `json:"rec_30_39"`
		Rec40P         float64 `json:"rec_40p"`
		Rec59          float64 `json:"rec_5_9"`
		RecAirYd       float64 `json:"rec_air_yd"`
		RecDrop        float64 `json:"rec_drop"`
		RecFd          float64 `json:"rec_fd"`
		Rec            float64 `json:"rec"`
		RecLng         float64 `json:"rec_lng"`
		RecTd40P       float64 `json:"rec_td_40p"`
		RecTd          float64 `json:"rec_td"`
		RecTdLng       float64 `json:"rec_td_lng"`
		RecTgt         float64 `json:"rec_tgt"`
		RecYar         float64 `json:"rec_yar"`
		RecYd          float64 `json:"rec_yd"`
		RecYpr         float64 `json:"rec_ypr"`
		RecYpt         float64 `json:"rec_ypt"`
		RushAtt        float64 `json:"rush_att"`
		RushBtkl       float64 `json:"rush_btkl"`
		RushFd         float64 `json:"rush_fd"`
		RushLng        float64 `json:"rush_lng"`
		RushRecYd      float64 `json:"rush_rec_yd"`
		RushRzAtt      float64 `json:"rush_rz_att"`
		RushTd         float64 `json:"rush_td"`
		RushTdLng      float64 `json:"rush_td_lng"`
		RushTklLoss    float64 `json:"rush_tkl_loss"`
		RushTklLossYd  float64 `json:"rush_tkl_loss_yd"`
		RushYac        float64 `json:"rush_yac"`
		RushYd         float64 `json:"rush_yd"`
		RushYpa        float64 `json:"rush_ypa"`
		SackYd         float64 `json:"sack_yd"`
		Snp            float64 `json:"snp"`
		StSnp          float64 `json:"st_snp"`
		StTklSolo      float64 `json:"st_tkl_solo"`
		TmDefSnp       float64 `json:"tm_def_snp"`
		TmOffSnp       float64 `json:"tm_off_snp"`
		TmStSnp        float64 `json:"tm_st_snp"`
		Xpa            float64 `json:"xpa"`
		Xpm            float64 `json:"xpm"`
	} `json:"stats"`
	Category     string `json:"category"`
	LastModified any    `json:"last_modified"`
	Week         any    `json:"week"`
	Season       string `json:"season"`
	SeasonType   string `json:"season_type"`
	Sport        string `json:"sport"`
	PlayerID     string `json:"player_id"`
	GameID       string `json:"game_id"`
	UpdatedAt    any    `json:"updated_at"`
	Team         string `json:"team"`
	Company      string `json:"company"`
	Opponent     any    `json:"opponent"`
	Player       Player `json:"player"`
}

// Get specific NFL player details.
// (GET `https://api.sleeper.app/players/nfl/<player id>)
func (c *Client) GetNflPlayer(playerID int) (Player, error) {
	player := Player{}

	url := fmt.Sprintf("%s/player/nfl/%d", c.sleeperURL, playerID)

	data, err := c.getRequest(url)
	if err != nil {
		return player, err
	}

	err = json.Unmarshal(data, &player)

	return player, err
}

// Get NFL players research
// `GET https://api.sleeper.app/players/nfl/research/<regular or post>/<year>/<week>`
func (c *Client) GetNflPlayerResearch(year int, week int, postseason bool) (map[string]PlayerResearch, error) {
	var results map[string]PlayerResearch
	reg := "regular"
	if postseason {
		reg = "post"
	}

	url := fmt.Sprintf("%s/players/nfl/research/%s/%d/%d", c.sleeperURL, reg, year, week)

	data, err := c.getRequest(url)
	if err != nil {
		return results, err
	}

	err = json.Unmarshal(data, &results)

	return results, err
}

// Get NFL player season stats
// `GET https://api.sleeper.app/stats/nfl/player/<player id>?season_type=<regular or post>&season=<season>`
func (c *Client) GetNflPlayerSeasonStats(playerID int, year int, postseason bool) (PlayerStats, error) {
	stats := PlayerStats{}
	reg := "regular"
	if postseason {
		reg = "post"
	}

	url := fmt.Sprintf("%s/stats/nfl/player/%d?season_type=%s&season=%d", c.sleeperURL, playerID, reg, year)

	data, err := c.getRequest(url)
	if err != nil {
		return stats, err
	}

	err = json.Unmarshal(data, &stats)

	return stats, err
}
