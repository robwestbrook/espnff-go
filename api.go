package espnffgo

import (
	"errors"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// Define ESPN endpoint and segment constants
const (
	BASE_SEASON_URL  = "https://lm-api-reads.fantasy.espn.com/apis/v3/games/ffl/seasons/"
	BASE_HISTORY_URL = "https://lm-api-reads.fantasy.espn.com/apis/v3/games/ffl/leagueHistory/"
	SEGMENT          = "/segments/0/leagues/"
)

var (
	ErrCreatingConnection = errors.New(
		"error creating connection to ESPN API",
	)
	ErrCreatingNewConnection = errors.New(
		"error creating a new connection to ESPN API",
	)
	ErrLoadingEnvironmentVariables = errors.New(
		"error loading environmental variables for .env file",
	)
	ErrCreatingHTTPRequest = errors.New(
		"error creating an HTTP request to ESPN API",
	)
	ErrMakingHTTPRequest = errors.New(
		"error making HTTP request to ESPN API",
	)
)

// Create an Environment Variables struct to hold values
// from .env file.
type EnvVars struct {
	LeagueID string
	EspnS2   string
	EspnSWID string
}

type ESPNConnection struct {
	Envs   EnvVars
	Client *http.Client
}

func (e *ESPNConnection) CreateESPNConnection() (*ESPNConnection, error) {
	c, err := NewESPNConnection()
	if err != nil {
		return nil, ErrCreatingConnection
	}
	return c, nil
}

func NewESPNConnection() (*ESPNConnection, error) {
	var connection ESPNConnection
	envs, err := loadEnv()
	if err != nil {
		return nil, ErrCreatingNewConnection
	}
	connection.Envs = *envs

	client := createHTTPClient()
	connection.Client = client

	return &connection, nil
}

// LoadEnv loads environmet variables
func loadEnv() (*EnvVars, error) {
	e := EnvVars{}
	err := godotenv.Load()
	if err != nil {
		return nil, ErrLoadingEnvironmentVariables
	}
	e.LeagueID = os.Getenv("LEAGUE_ID")
	e.EspnS2 = os.Getenv("ESPN_S2")
	e.EspnSWID = os.Getenv("SWID")
	return &e, nil
}

// CreateHTTPClient creates an HTTP client to use for
// requests to the ESPN API.
func createHTTPClient() *http.Client {
	return &http.Client{}
}

// CreateHTTPRequest creates an HTTP request for the
// ESPN API. The API endpoint URL, the required cookies,
// and the needed headers are used to create the request.
func createHTTPRequest(params string, e *EnvVars) (*http.Request, error) {
	// Create a new request
	req, err := http.NewRequest("GET", BASE_SEASON_URL+"2024"+SEGMENT+e.LeagueID+params, nil)
	if err != nil {
		return nil, ErrCreatingHTTPRequest
	}

	// Create cookies for ESPN API access
	espnS2Cookie := http.Cookie{
		Name:  "espn_s2",
		Value: e.EspnS2,
	}
	swidCookie := http.Cookie{
		Name:  "SWID",
		Value: e.EspnSWID,
	}

	// Add cookies to request
	req.AddCookie(&espnS2Cookie)
	req.AddCookie(&swidCookie)

	// Set request header for JSON response
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("User-Agent", getRandomUserAgent())

	// Add special header if params is for player info
	if params == "?view=kona_player_info" {
		req.Header.Set("x-fantasy-filter", `{"players": {"limit": 4000,"sortDraftRanks": {"sortPriority": 100, "sortAsc": true, "value": "PPR"}}}`)
	}

	return req, nil
}

// MakeHTTPRequest executes the request and recieves the
// response from the ESPN API.
func makeHTTPRequest(client *http.Client, req http.Request) (*http.Response, error) {
	res, err := client.Do(&req)
	if err != nil {
		return nil, ErrMakingHTTPRequest
	}
	return res, nil
}
