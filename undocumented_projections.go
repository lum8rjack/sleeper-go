package sleeper

import (
	"encoding/json"
	"fmt"
)

type Projections []struct {
	Date  string `json:"date"`
	Stats struct {
		AdpDdPpr       float64 `json:"adp_dd_ppr"`
		BlkKick        float64 `json:"blk_kick"`
		BonusRecWr     float64 `json:"bonus_rec_wr"`
		CmpPct         float64 `json:"cmp_pct"`
		DefFumTd       float64 `json:"def_fum_td"`
		DefKrYd        float64 `json:"def_kr_yd"`
		DefPrYd        float64 `json:"def_pr_yd"`
		DefTd          float64 `json:"def_td"`
		Ff             float64 `json:"ff"`
		Fum            float64 `json:"fum"`
		FumLost        float64 `json:"fum_lost"`
		FumRec         float64 `json:"fum_rec"`
		Gp             float64 `json:"gp"`
		Int            float64 `json:"int"`
		Pass2Pt        float64 `json:"pass_2pt"`
		PassAtt        float64 `json:"pass_att"`
		PassCmp        float64 `json:"pass_cmp"`
		PassCmp40P     float64 `json:"pass_cmp_40p"`
		PassFd         float64 `json:"pass_fd"`
		PassInc        float64 `json:"pass_inc"`
		PassInt        float64 `json:"pass_int"`
		PassIntTd      float64 `json:"pass_int_td"`
		PassSack       float64 `json:"pass_sack"`
		PassTd         float64 `json:"pass_td"`
		PassTd40P      float64 `json:"pass_td_40p"`
		PassYd         float64 `json:"pass_yd"`
		PosAdpDdPpr    float64 `json:"pos_adp_dd_ppr"`
		Pr             float64 `json:"pr"`
		PrYd           float64 `json:"pr_yd"`
		PtsAllow       float64 `json:"pts_allow"`
		PtsAllow2127   float64 `json:"pts_allow_21_27"`
		PtsHalfPpr     float64 `json:"pts_half_ppr"`
		PtsPpr         float64 `json:"pts_ppr"`
		PtsStd         float64 `json:"pts_std"`
		Rec            float64 `json:"rec"`
		Rec04          float64 `json:"rec_0_4"`
		Rec1019        float64 `json:"rec_10_19"`
		Rec2029        float64 `json:"rec_20_29"`
		Rec3039        float64 `json:"rec_30_39"`
		Rec40P         float64 `json:"rec_40p"`
		Rec59          float64 `json:"rec_5_9"`
		RecFd          float64 `json:"rec_fd"`
		RecTd          float64 `json:"rec_td"`
		RecTgt         float64 `json:"rec_tgt"`
		RecYd          float64 `json:"rec_yd"`
		RushAtt        float64 `json:"rush_att"`
		RushFd         float64 `json:"rush_fd"`
		RushTd         float64 `json:"rush_td"`
		RushYd         float64 `json:"rush_yd"`
		Sack           float64 `json:"sack"`
		Safe           float64 `json:"safe"`
		TklLoss        float64 `json:"tkl_loss"`
		YdsAllow       float64 `json:"yds_allow"`
		YdsAllow300349 float64 `json:"yds_allow_300_349"`
	} `json:"stats,omitempty"`
	Category   string `json:"category"`
	Week       int    `json:"week"`
	Season     string `json:"season"`
	SeasonType string `json:"season_type"`
	Sport      string `json:"sport"`
	PlayerID   string `json:"player_id"`
	GameID     string `json:"game_id"`
	Team       string `json:"team"`
	Company    string `json:"company"`
	Opponent   string `json:"opponent"`
	Player     Player `json:"player"`
}

// Get NFL player score projections for a specific season and week.
// (GET `https://api.sleeper.app/projections/nfl/<season>/<week>?season_type=regular&position[]=FLEX&position[]=K&position[]=QB&position[]=RB&position[]=TE&position[]=WR&position[]=DEF`)
func (c *Client) GetNflProjections(season int, week int) (Projections, error) {
	projections := Projections{}

	url := fmt.Sprintf("%s/projections/nfl/%d/%d?eason_type=regular&position[]=FLEX&position[]=K&position[]=QB&position[]=RB&position[]=TE&position[]=WR&position[]=DEF", sleeperUndocumentedURL, season, week)

	data, err := c.getRequest(url)
	if err != nil {
		return projections, err
	}

	err = json.Unmarshal(data, &projections)

	return projections, err
}
