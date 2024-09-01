package espnffgo

type MTransactions2 struct {
	DraftDetail struct {
		Drafted    bool `json:"draftDetail"`
		InProgress bool `json:"inProgress"`
	} `json:"draftDetail"`
	GameID        int `json:"gameId"`
	ScoringPeriod int `json:"scoringPeriod"`
	SeasonID      int `json:"seasonId"`
	SegmentID     int `json:"segmentId"`
	Status        struct {
		ActivationDate           int64          `json:"activationDate"`
		CreatedAsLeagueType      int            `json:"createdAsLeagueType"`
		CurrentLeagueType        int            `json:"currentLeagueType"`
		CurrentMatchupPeriod     int            `json:"currentMatchupPeriod"`
		FinalScoringPeriod       int            `json:"finalScoringPeriod"`
		FirstScoringPeriod       int            `json:"firstScoringPeriod"`
		IsActive                 bool           `json:"isActive"`
		IsExpired                bool           `json:"isExpired"`
		IsFull                   bool           `json:"isFull"`
		IsPlayoffMatchupEdited   bool           `json:"isPlayoffMatchupEdited"`
		IsToBeDeleted            bool           `json:"isToBeDeleted"`
		IsViewable               bool           `json:"isViewable"`
		IsWaiverOrderEdited      bool           `json:"isWaiverOrderEdited"`
		LatestScoringPeriod      int            `json:"latestScoringPeriod"`
		PreviousSeasons          []int64        `json:"previousSeasons"`
		StandingsUpdateDate      int64          `json:"standingsUpdateDate"`
		TeamsJoined              int            `json:"teamsJoined"`
		TransactionScoringPeriod int            `json:"transactionScoringPeriod"`
		WaiverLastExecutionDate  int64          `json:"waiverLastExecutionDate"`
		WaiverProcessStatus      map[string]int `json:"waiverProcessStatus"`
	} `json:"status"`
	Transactions []struct {
		BidAmount         float64 `json:"bidAmount"`
		ExecutionType     string  `json:"executionType"`
		ID                string  `jsonL"id"`
		IsActiveTeamOwner bool    `json:"isActiveTeamOwner"`
		IsLeagueManager   bool    `json:"isLeagueManager"`
		IsPending         bool    `json:"isPending"`
		Items             []struct {
			FromLineupSlotID  int    `json:"fromLineupSlotId"`
			FromTeamID        int    `json:"fromTeamId"`
			IsKeeper          bool   `json:"isKeeper"`
			OverallPickNumber int64  `json:"overallPickNumber"`
			PlayerID          int64  `json:"playerId"`
			ToLineupSlotID    int    `json:"toLineupSlotId"`
			ToTeamID          int    `json:"toTeamId"`
			Type              string `json:"type"`
		} `json:"items"`
		ProposedDate    int64   `json:"proposedDate"`
		Rating          float64 `json:"rating"`
		ScoringPeriodID int     `json:"scoringPeriodId"`
		Status          string  `json:"status"`
		TeamID          int     `json:"teamId"`
		Type            string  `json:"type"`
	} `json:"transactions"`
}
