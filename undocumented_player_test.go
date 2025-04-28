package sleeper

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetNflPlayer(t *testing.T) {
	// Create mock player data
	mockPlayer := Player{
		PlayerID:  "123",
		FirstName: "John",
		LastName:  "Doe",
		Position:  "QB",
		Team:      "SF",
		Status:    "ACT",
		Height:    "6'2\"",
		Weight:    "220",
		BirthDate: "1990-01-01",
		College:   "Stanford",
		Number:    12,
	}

	// Convert mock data to JSON
	mockData, err := json.Marshal(mockPlayer)
	if err != nil {
		t.Fatalf("Failed to marshal mock data: %v", err)
	}

	// Create test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedPath := "/player/nfl/123"
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
	player, err := client.GetNflPlayer(123)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify player data
	if player.PlayerID != "123" {
		t.Errorf("Expected PlayerID '123', got '%s'", player.PlayerID)
	}
	if player.FirstName != "John" {
		t.Errorf("Expected FirstName 'John', got '%s'", player.FirstName)
	}
	if player.Position != "QB" {
		t.Errorf("Expected Position 'QB', got '%s'", player.Position)
	}
}

func TestGetNflPlayerResearch(t *testing.T) {
	// Create mock research data
	mockResearch := map[string]PlayerResearch{
		"123": {
			Owned:   0.85,
			Started: 0.75,
		},
		"456": {
			Owned:   0.65,
			Started: 0.55,
		},
	}

	// Convert mock data to JSON
	mockData, err := json.Marshal(mockResearch)
	if err != nil {
		t.Fatalf("Failed to marshal mock data: %v", err)
	}

	// Create test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedPath := "/players/nfl/research/regular/2023/1"
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
	research, err := client.GetNflPlayerResearch(2023, 1, false)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify research data
	if len(research) != 2 {
		t.Errorf("Expected 2 players, got %d", len(research))
	}
	if research["123"].Owned != 0.85 {
		t.Errorf("Expected Owned 0.85, got %f", research["123"].Owned)
	}
	if research["123"].Started != 0.75 {
		t.Errorf("Expected Started 0.75, got %f", research["123"].Started)
	}
}

func TestGetNflPlayerSeasonStats(t *testing.T) {
	// Create mock stats data
	mockStats := PlayerStats{
		Date:       "2023-09-10",
		Week:       1,
		Season:     "2023",
		SeasonType: "regular",
		Sport:      "nfl",
		PlayerID:   "123",
		GameID:     "game1",
		Team:       "SF",
		Company:    "sleeper",
		Opponent:   "PIT",
		Stats: struct {
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
		}{
			PassYd:  4000.0,
			PassTd:  30.0,
			PassInt: 10.0,
			RushYd:  200.0,
			RushTd:  2.0,
			PtsPpr:  350.0,
		},
	}

	// Convert mock data to JSON
	mockData, err := json.Marshal(mockStats)
	if err != nil {
		t.Fatalf("Failed to marshal mock data: %v", err)
	}

	// Create test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedPath := "/stats/nfl/player/123"
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
	stats, err := client.GetNflPlayerSeasonStats(123, 2023, false)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify stats data
	if stats.Stats.PassYd != 4000.0 {
		t.Errorf("Expected PassYd 4000.0, got %f", stats.Stats.PassYd)
	}
	if stats.Stats.PassTd != 30.0 {
		t.Errorf("Expected PassTd 30.0, got %f", stats.Stats.PassTd)
	}
	if stats.Stats.PtsPpr != 350.0 {
		t.Errorf("Expected PtsPpr 350.0, got %f", stats.Stats.PtsPpr)
	}
}

func TestGetNflPlayerError(t *testing.T) {
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
	_, err := client.GetNflPlayer(123)
	if err == nil {
		t.Error("Expected error for 404 response, got nil")
	}
}

func TestGetNflPlayerResearchError(t *testing.T) {
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
	_, err := client.GetNflPlayerResearch(2023, 1, false)
	if err == nil {
		t.Error("Expected error for 404 response, got nil")
	}
}

func TestGetNflPlayerSeasonStatsError(t *testing.T) {
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
	_, err := client.GetNflPlayerSeasonStats(123, 2023, false)
	if err == nil {
		t.Error("Expected error for 404 response, got nil")
	}
}

func TestGetNflPlayerInvalidResponse(t *testing.T) {
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
	_, err := client.GetNflPlayer(123)
	if err == nil {
		t.Error("Expected error for invalid JSON, got nil")
	}
}

func TestGetNflPlayerResearchInvalidResponse(t *testing.T) {
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
	_, err := client.GetNflPlayerResearch(2023, 1, false)
	if err == nil {
		t.Error("Expected error for invalid JSON, got nil")
	}
}

func TestGetNflPlayerSeasonStatsInvalidResponse(t *testing.T) {
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
	_, err := client.GetNflPlayerSeasonStats(123, 2023, false)
	if err == nil {
		t.Error("Expected error for invalid JSON, got nil")
	}
}
