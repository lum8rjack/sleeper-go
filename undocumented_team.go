package sleeper

import (
	"encoding/json"
	"fmt"
)

type TeamDepthChart struct {
	Db   []string `json:"DB"`
	Dl   []string `json:"DL"`
	Fs   []string `json:"FS"`
	K    []string `json:"K"`
	Lb   []string `json:"LB"`
	Lcb  []string `json:"LCB"`
	Lde  []string `json:"LDE"`
	Ldt  []string `json:"LDT"`
	Lolb []string `json:"LOLB"`
	Ls   []string `json:"LS"`
	Mlb  []string `json:"MLB"`
	Nb   []string `json:"NB"`
	Ol   []string `json:"OL"`
	P    []string `json:"P"`
	Qb   []string `json:"QB"`
	Rb   []string `json:"RB"`
	Rcb  []string `json:"RCB"`
	Rde  []string `json:"RDE"`
	Rdt  []string `json:"RDT"`
	Rolb []string `json:"ROLB"`
	Ss   []string `json:"SS"`
	Te   []string `json:"TE"`
	Wr1  []string `json:"WR1"`
	Wr2  []string `json:"WR2"`
	Wr3  []string `json:"WR3"`
}

// Get NFL team depth chart.
// `GET https://api.sleeper.app/players/nfl/<team>/depth_chart`
func (c *Client) GetNflTeamDepthChart(team string) (TeamDepthChart, error) {
	tdc := TeamDepthChart{}

	url := fmt.Sprintf("%s/player/nfl/%s/depth_chart", sleeperUndocumentedURL, team)

	data, err := c.getRequest(url)
	if err != nil {
		return tdc, err
	}

	err = json.Unmarshal(data, &tdc)

	return tdc, err
}
