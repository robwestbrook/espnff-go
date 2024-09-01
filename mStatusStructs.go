package espnffgo

type MStatus struct {
	DraftDetail struct {
		Drafted    bool `json:"drafted"`
		InProgress bool `json:"inProgress"`
	} `json:"draftDetail"`
	GameID          int   `json:"gameId"`
	LeagueID        int64 `json:"id"`
	ScoringPeriodID int   `json:"scoringPeriodId"`
	SeasonID        int64 `json:"seasonId"`
	SegmentID       int   `json:"segmentId"`
	Status          struct {
		ActivatedDate       int64 `json:"activatedDate"`
		CreatedAsLeagueType int   `json:"createdAsLeagueType"`
		CreationInfo        struct {
			ClientAddress any    `json:"clientAddress"`
			Platform      string `json:"platform"`
			Source        string `json:"source"`
		}
		CurrentLeagueType      int  `json:"currentLeagueType"`
		CurrentMatchupPeriod   int  `json:"currentMatchupPeriod"`
		FinalScoringPeriod     int  `json:"finalScoringPeriod"`
		FirstScoringPeriod     int  `json:"firstScoringPeriod"`
		IsActive               bool `json:"isActive"`
		IsExpired              bool `json:"isExpired"`
		IsFull                 bool `json:"isFull"`
		IsPlayoffMatchupEdited bool `json:"isPlayoffMatchupEdited"`
		IsToBeDeleted          bool `json:"isToBeDeleted"`
		IsViewable             bool `json:"isViewable"`
		IsWaiverOrderEdited    bool `json:"isWaiverOrderEdited"`
		LastUpdateInfo         struct {
			ClientAddress any    `json:"clientAddress"`
			Platform      string `json:"platform"`
			Source        string `json:"source"`
		} `json:"lastUpdateInfo"`
		LatestScoringPeriod      int            `json:"latestScoringPeriod"`
		PreviousSeasons          []int          `json:"previousSeasons"`
		StandingsUpdateDate      int64          `json:"standingsUpdateDate"`
		TeamsJoined              int            `json:"teamsJoined"`
		TransactionScoringPeriod int            `json:"transactionScoringPeriod"`
		WaiverLastExecutionDate  int64          `json:"waiverLastExecutionDate"`
		WaiverProcessStatus      map[string]int `json:"waiverProcessStatus"`
	} `json:"status"`
}
