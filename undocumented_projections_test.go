package sleeper

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetNflProjections(t *testing.T) {
	// Create mock projections data
	mockProjections := Projections{
		{
			Date:       "2023-09-10",
			Week:       1,
			Season:     "2023",
			SeasonType: "regular",
			Sport:      "nfl",
			PlayerID:   "player1",
			GameID:     "game1",
			Team:       "SF",
			Company:    "sleeper",
			Opponent:   "PIT",
			Category:   "projection",
			Stats: struct {
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
			}{
				PassYd:  250.5,
				PassTd:  2.0,
				PassInt: 0.5,
				RushYd:  25.0,
				RushTd:  0.5,
				Rec:     5.0,
				RecYd:   50.0,
				RecTd:   0.5,
				PtsPpr:  25.5,
			},
		},
		{
			Date:       "2023-09-10",
			Week:       1,
			Season:     "2023",
			SeasonType: "regular",
			Sport:      "nfl",
			PlayerID:   "player2",
			GameID:     "game1",
			Team:       "PIT",
			Company:    "sleeper",
			Opponent:   "SF",
			Category:   "projection",
			Stats: struct {
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
			}{
				RushYd: 100.0,
				RushTd: 1.0,
				Rec:    3.0,
				RecYd:  30.0,
				RecTd:  0.5,
				PtsPpr: 20.5,
			},
		},
	}

	// Convert mock data to JSON
	mockData, err := json.Marshal(mockProjections)
	if err != nil {
		t.Fatalf("Failed to marshal mock data: %v", err)
	}

	// Create test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedPath := "/projections/nfl/2023/1"
		if r.URL.Path != expectedPath {
			t.Errorf("Expected path %s, got %s", expectedPath, r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mockData)
	}))
	defer ts.Close()

	// Create client with test server URL
	client := NewClientWithOptions(ClientOptions{
		BaseURL: ts.URL,
	})

	// Test successful request
	projections, err := client.GetNflProjections(2023, 1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify projections data
	if len(projections) != 2 {
		t.Errorf("Expected 2 projections, got %d", len(projections))
	}

	// Test first player's stats
	if projections[0].Stats.PassYd != 250.5 {
		t.Errorf("Expected PassYd 250.5, got %f", projections[0].Stats.PassYd)
	}
	if projections[0].Stats.PassTd != 2.0 {
		t.Errorf("Expected PassTd 2.0, got %f", projections[0].Stats.PassTd)
	}
	if projections[0].Stats.PtsPpr != 25.5 {
		t.Errorf("Expected PtsPpr 25.5, got %f", projections[0].Stats.PtsPpr)
	}

	// Test second player's stats
	if projections[1].Stats.RushYd != 100.0 {
		t.Errorf("Expected RushYd 100.0, got %f", projections[1].Stats.RushYd)
	}
	if projections[1].Stats.RushTd != 1.0 {
		t.Errorf("Expected RushTd 1.0, got %f", projections[1].Stats.RushTd)
	}
	if projections[1].Stats.PtsPpr != 20.5 {
		t.Errorf("Expected PtsPpr 20.5, got %f", projections[1].Stats.PtsPpr)
	}
}

func TestGetNflProjectionsError(t *testing.T) {
	// Create test server that returns 404
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer ts.Close()

	// Create client with test server URL
	client := NewClientWithOptions(ClientOptions{
		BaseURL: ts.URL,
	})

	// Test error response
	_, err := client.GetNflProjections(2023, 1)
	if err == nil {
		t.Error("Expected error for 404 response, got nil")
	}
}

func TestGetNflProjectionsInvalidResponse(t *testing.T) {
	// Create test server that returns invalid JSON
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("invalid json"))
	}))
	defer ts.Close()

	// Create client with test server URL
	client := NewClientWithOptions(ClientOptions{
		BaseURL: ts.URL,
	})

	// Test invalid JSON response
	_, err := client.GetNflProjections(2023, 1)
	if err == nil {
		t.Error("Expected error for invalid JSON, got nil")
	}
}

func TestGetNflProjectionsEmptyResponse(t *testing.T) {
	// Create test server that returns empty projections
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("[]"))
	}))
	defer ts.Close()

	// Create client with test server URL
	client := NewClientWithOptions(ClientOptions{
		BaseURL: ts.URL,
	})

	// Test empty response
	projections, err := client.GetNflProjections(2023, 1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify empty projections
	if len(projections) != 0 {
		t.Errorf("Expected empty projections, got %d projections", len(projections))
	}
}
