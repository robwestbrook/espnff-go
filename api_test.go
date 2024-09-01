package espnffgo

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

func TestESPNConnection(t *testing.T) {
	t.Run("create HTTP client", func(t *testing.T) {
		var x interface{} = *createHTTPClient()
		_, ok := x.(http.Client)
		if !ok {
			t.Errorf("Did not create a valid HTTP client")
		}
	})
	t.Run("create HTTP request", func(t *testing.T) {
		envs, err := loadEnv()
		if err != nil {
			t.Errorf("Got unexpected error: %q", err)
		}
		got, err := createHTTPRequest("view=mTeam", envs)
		if err != nil {
			t.Errorf("Error creating HTTP request")
		}
		cookie, err := got.Cookie("espn_s2")
		if err != nil {
			t.Errorf("error reading cookie: espn_s2")
		}

		if cookie.Value != envs.EspnS2 {
			t.Errorf("got espn_s2 cookie\n %q\n want %q\n", cookie, envs.EspnS2)
		}

		cookie, err = got.Cookie("SWID")
		if err != nil {
			t.Errorf("error reading cookie: SWID")
		}
		if cookie.Value != envs.EspnSWID {
			t.Errorf("got espn_s2 cookie\n %q\n want %q\n", cookie, envs.EspnSWID)
		}
	})
	t.Run("make HTTP request", func(t *testing.T) {
		vars, err := loadEnv()
		if err != nil {
			t.Errorf("Got unexpected error: %q", err)
		}
		client := createHTTPClient()
		req, err := createHTTPRequest("?view=mTeam", vars)
		if err != nil {
			t.Errorf("Error creating HTTP request: %q", err)
		}

		res, err := makeHTTPRequest(client, *req)
		if err != nil {
			t.Errorf("Error making HTTP request: %q", err)
		}

		got := res.StatusCode
		want := http.StatusOK

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}

		var data interface{}
		e, _ := io.ReadAll(res.Body)
		err = json.Unmarshal(e, &data)
		if err != nil {
			t.Errorf("Error decoding JSON: %q", err)
		}
	})

	t.Run("create new connection", func(t *testing.T) {
		envs, err := loadEnv()
		if err != nil {
			t.Errorf("Got unexpected error: %q", err)
		}
		var connection *ESPNConnection
		connection, err = connection.CreateESPNConnection()
		if err != nil {
			t.Fatalf("Error creating connection: %q", err)
		}
		if connection.Envs.LeagueID != envs.LeagueID {
			t.Errorf("LeagueID: expected:\n %q\n got %q\n", envs.LeagueID, connection.Envs.LeagueID)
		}
		if connection.Envs.EspnS2 != envs.EspnS2 {
			t.Errorf("LeagueID: expected:\n %q\n got %q\n", envs.EspnS2, connection.Envs.EspnS2)
		}
		if connection.Envs.EspnSWID != envs.EspnSWID {
			t.Errorf("LeagueID: expected:\n %q\n got %q\n", envs.EspnSWID, connection.Envs.EspnS2)
		}
		var x interface{} = connection.Client
		_, ok := x.(*http.Client)
		if !ok {
			t.Errorf("Did not create a valid HTTP client")
		}
	})
}

func TestProcessESPNData(t *testing.T) {

	t.Run("process ESPN mTeam data", func(t *testing.T) {
		var team MTeam
		vars, err := loadEnv()
		if err != nil {
			t.Errorf("Got unexpected error: %q", err)
		}
		client := createHTTPClient()
		req, err := createHTTPRequest("?view=mTeam", vars)
		if err != nil {
			t.Errorf("Error creating HTTP request: %q", err)
		}

		res, err := makeHTTPRequest(client, *req)
		if err != nil {
			t.Errorf("Error making HTTP request: %q", err)
		}

		e, err := io.ReadAll(res.Body)
		if err != nil {
			t.Errorf("Error reading response body JSON")
		}
		err = json.Unmarshal(e, &team)
		if err != nil {
			t.Errorf("Error decoding JSON:\n %q\n", err)
		}
	})

	t.Run("process ESPN mMatchup data", func(t *testing.T) {
		var matchup MMatchup
		vars, err := loadEnv()
		if err != nil {
			t.Errorf("Got unexpected error: %q", err)
		}
		client := createHTTPClient()
		req, err := createHTTPRequest("?view=mMatchup", vars)
		if err != nil {
			t.Errorf("Error creating HTTP request: %q", err)
		}

		res, err := makeHTTPRequest(client, *req)
		if err != nil {
			t.Errorf("Error making HTTP request: %q", err)
		}

		e, err := io.ReadAll(res.Body)
		if err != nil {
			t.Errorf("Error reading response body JSON")
		}
		err = json.Unmarshal(e, &matchup)
		if err != nil {
			t.Errorf("Error decoding JSON:\n %q\n", err)
		}
	})

	t.Run("process ESPN mRoster data", func(t *testing.T) {
		var roster MRoster
		vars, err := loadEnv()
		if err != nil {
			t.Errorf("Got unexpected error: %q", err)
		}
		client := createHTTPClient()
		req, err := createHTTPRequest("?view=mRoster", vars)
		if err != nil {
			t.Errorf("Error creating HTTP request: %q", err)
		}

		res, err := makeHTTPRequest(client, *req)
		if err != nil {
			t.Errorf("Error making HTTP request: %q", err)
		}

		e, err := io.ReadAll(res.Body)
		if err != nil {
			t.Errorf("Error reading response body JSON")
		}
		err = json.Unmarshal(e, &roster)
		if err != nil {
			t.Errorf("Error decoding JSON:\n %q\n", err)
		}
	})

	t.Run("process ESPN mSettings data", func(t *testing.T) {
		var settings MSettings
		vars, err := loadEnv()
		if err != nil {
			t.Errorf("Got unexpected error: %q", err)
		}
		client := createHTTPClient()
		req, err := createHTTPRequest("?view=mSettings", vars)
		if err != nil {
			t.Errorf("Error creating HTTP request: %q", err)
		}

		res, err := makeHTTPRequest(client, *req)
		if err != nil {
			t.Errorf("Error making HTTP request: %q", err)
		}

		e, err := io.ReadAll(res.Body)
		if err != nil {
			t.Errorf("Error reading response body JSON")
		}
		err = json.Unmarshal(e, &settings)
		if err != nil {
			t.Errorf("Error decoding JSON:\n %q\n", err)
		}
	})

	t.Run("process ESPN mBoxScore data", func(t *testing.T) {
		var boxscores MBoxscore
		vars, err := loadEnv()
		if err != nil {
			t.Errorf("Got unexpected error: %q", err)
		}
		client := createHTTPClient()
		req, err := createHTTPRequest("?view=mBoxscore", vars)
		if err != nil {
			t.Errorf("Error creating HTTP request: %q", err)
		}

		res, err := makeHTTPRequest(client, *req)
		if err != nil {
			t.Errorf("Error making HTTP request: %q", err)
		}

		e, err := io.ReadAll(res.Body)
		if err != nil {
			t.Errorf("Error reading response body JSON")
		}
		err = json.Unmarshal(e, &boxscores)
		if err != nil {
			t.Errorf("Error decoding JSON:\n %q\n", err)
		}
	})

	t.Run("process ESPN mMatchupScore data", func(t *testing.T) {
		var matchupScore MMAtchupScore
		vars, err := loadEnv()
		if err != nil {
			t.Errorf("Got unexpected error: %q", err)
		}
		client := createHTTPClient()
		req, err := createHTTPRequest("?view=mMatchupScore", vars)
		if err != nil {
			t.Errorf("Error creating HTTP request: %q", err)
		}

		res, err := makeHTTPRequest(client, *req)
		if err != nil {
			t.Errorf("Error making HTTP request: %q", err)
		}

		e, err := io.ReadAll(res.Body)
		if err != nil {
			t.Errorf("Error reading response body JSON")
		}
		err = json.Unmarshal(e, &matchupScore)
		if err != nil {
			t.Errorf("Error decoding JSON:\n %q\n", err)
		}
	})

	t.Run("process ESPN kona_player_info data", func(t *testing.T) {
		var konaPlayerInfo KonaPlayerInfo
		vars, err := loadEnv()
		if err != nil {
			t.Errorf("Got unexpected error: %q", err)
		}
		client := createHTTPClient()
		req, err := createHTTPRequest("?view=kona_player_info", vars)
		if err != nil {
			t.Errorf("Error creating HTTP request: %q", err)
		}

		res, err := makeHTTPRequest(client, *req)
		if err != nil {
			t.Errorf("Error making HTTP request: %q", err)
		}

		e, err := io.ReadAll(res.Body)
		if err != nil {
			t.Errorf("Error reading response body JSON")
		}
		err = json.Unmarshal(e, &konaPlayerInfo)
		if err != nil {
			t.Errorf("Error decoding JSON:\n %q\n", err)
		}
	})

	t.Run("process ESPN mDraftDetail data", func(t *testing.T) {
		var draftDetail MDraftDetail
		vars, err := loadEnv()
		if err != nil {
			t.Errorf("Got unexpected error: %q", err)
		}
		client := createHTTPClient()
		req, err := createHTTPRequest("?view=mDraftDetail", vars)
		if err != nil {
			t.Errorf("Error creating HTTP request: %q", err)
		}

		res, err := makeHTTPRequest(client, *req)
		if err != nil {
			t.Errorf("Error making HTTP request: %q", err)
		}

		e, err := io.ReadAll(res.Body)
		if err != nil {
			t.Errorf("Error reading response body JSON")
		}
		err = json.Unmarshal(e, &draftDetail)
		if err != nil {
			t.Errorf("Error decoding JSON:\n %q\n", err)
		}
	})

	t.Run("process ESPN mTransactions2 data", func(t *testing.T) {
		var transactions MTransactions2
		vars, err := loadEnv()
		if err != nil {
			t.Errorf("Got unexpected error: %q", err)
		}
		client := createHTTPClient()
		req, err := createHTTPRequest("?view=mTransactions2", vars)
		if err != nil {
			t.Errorf("Error creating HTTP request: %q", err)
		}

		res, err := makeHTTPRequest(client, *req)
		if err != nil {
			t.Errorf("Error making HTTP request: %q", err)
		}

		e, err := io.ReadAll(res.Body)
		if err != nil {
			t.Errorf("Error reading response body JSON")
		}
		err = json.Unmarshal(e, &transactions)
		if err != nil {
			t.Errorf("Error decoding JSON:\n %q\n", err)
		}
	})

	t.Run("process ESPN mStatus data", func(t *testing.T) {
		var status MStatus
		vars, err := loadEnv()
		if err != nil {
			t.Errorf("Got unexpected error: %q", err)
		}
		client := createHTTPClient()
		req, err := createHTTPRequest("?view=mStatus", vars)
		if err != nil {
			t.Errorf("Error creating HTTP request: %q", err)
		}

		res, err := makeHTTPRequest(client, *req)
		if err != nil {
			t.Errorf("Error making HTTP request: %q", err)
		}

		e, err := io.ReadAll(res.Body)
		if err != nil {
			t.Errorf("Error reading response body JSON")
		}
		err = json.Unmarshal(e, &status)
		if err != nil {
			t.Errorf("Error decoding JSON:\n %q\n", err)
		}
	})

	// t.Run("process ESPN kona_league_communication data", func(t *testing.T) {
	// 	var comms KonaLeagueCommunication
	// 	vars, err := loadEnv()
	// 	if err != nil {
	// 		t.Errorf("Got unexpected error: %q", err)
	// 	}
	// 	client := createHTTPClient()
	// 	req, err := createHTTPRequest("?view=kona_league_communication", vars)
	// 	if err != nil {
	// 		t.Errorf("Error creating HTTP request: %q", err)
	// 	}

	// 	res, err := makeHTTPRequest(client, *req)
	// 	if err != nil {
	// 		t.Errorf("Error making HTTP request: %q", err)
	// 	}

	// 	e, err := io.ReadAll(res.Body)
	// 	if err != nil {
	// 		t.Errorf("Error reading response body JSON")
	// 	}
	// 	err = json.Unmarshal(e, &comms)
	// 	if err != nil {
	// 		t.Errorf("Error decoding JSON:\n %q\n", err)
	// 	}
	// })
}
