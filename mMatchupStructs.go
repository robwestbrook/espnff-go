package espnffgo

// ScoreByStatData struct for team schedules retrieved from the
// ESPM "?view=mMatchup" endpoint.
type ScoreByStatData struct {
	Ineligible bool        `json:"ineligible"`
	Rank       float64     `json:"rank"`
	Result     interface{} `json:"result"` // Use interface{} for handling nulls or various data types
	Score      float64     `json:"score"`
}

// RankingsPlayer struct for team schedules retrieved from the
// ESPM "?view=mMatchup" endpoint.
type RankingsPlayer struct {
	AuctionValue int     `json:"auctionValue"`
	Rank         float64 `json:"rank"`
	RankSourceID int     `json:"rankSourceId"`
	RankType     string  `json:"rankType"`
	SlotID       int     `json:"slotId"`
}

// Player struct for team schedules retrieved from the
// ESPM "?view=mMatchup" endpoint.
type PlayerData struct {
	Active            bool   `json:"active"`
	DefaultPositionID int    `json:"defaultPositionId"`
	Droppable         bool   `json:"droppable"`
	EligibleSlots     []int  `json:"eligibleSlots"`
	FirstName         string `json:"firstName"`
	FullName          string `json:"fullName"`
	PlayerID          int64  `json:"id"`
	Injured           bool   `json:"injured"`
	InjuryStatus      string `json:"injuryStatus"`
	LastName          string `json:"lastName"`
	LastNewsDate      int64  `json:"lastNewsDate"`
	LastVideoDate     int64  `json:"lastVideoDate"`
	ProTeamID         int    `json:"proTeamId"`
}

// MMatchup defines a struct to store the ESPN API response
// to the params "?view=mMatchup"
type MMatchup struct {
	DraftDetail struct {
		Drafted    bool `json:"drafted"`
		InProgress bool `json:"inProgress"`
	} `json:"draftDetail"`
	GameID   int   `json:"gameId"`
	LeagueID int64 `json:"id"`
	Schedule []struct {
		Away struct {
			CumulativeScore struct {
				Losses      int                     `json:"losses"`
				StatBySlot  interface{}             `json:"statBySlot"`
				Ties        int                     `json:"ties"`
				Wins        int                     `json:"wins"`
				ScoreByStat map[int]ScoreByStatData `json:"scoreByStat"`
			} `json:"cumulativeScore"`
			GamesPlayed                   int `json:"gamesPlayed"`
			RosterForCurrentScoringPeriod struct {
				AppliedStatTotal float64 `json:"appliedStatTotal"`
				Entries          []struct {
					AcquisitionDate       int64  `json:"acquisitionDate"`
					AcquisitionType       string `json:"acquisitionType"`
					InjuryStatus          string `json:"injuryStatus"`
					LineupSlotID          int    `json:"lineupSlotId"`
					PendingTransactionIDs []int  `json:"pendingTransactionIds"`
					PlayerID              int64  `json:"playerId"`
					PlayerPoolEntry       struct {
						AppliedStatTotal float64                     `json:"appliedStatTotal"`
						PlayerID         int64                       `json:"id"`
						KeeperValue      int                         `json:"keeperValue"`
						LineupLocked     bool                        `json:"lineupLocked"`
						OnTeamID         int                         `json:"onTeamId"`
						Rankings         map[string][]RankingsPlayer `json:"rankings"`
						UniverseID       int                         `json:"universeId"`
						RosterLocked     bool                        `json:"rosterLocked"`
						Status           string                      `json:"status"`
						TradeLocked      bool                        `json:"readeLocked"`
					} `json:"playerPoolEntry"`
					Status string `json:"status"`
				} `json:"entries"`
			} `json:"rosterForCurrentScoringPeriod"`
		} `json:"away"`
	} `json:"schedule"`
}
