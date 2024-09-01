package espnffgo

type ScheduleStruct struct {
	Adjustment      float64 `json:"adjustment"`
	CumulativeScore struct {
		Losses     int     `json:"losses"`
		StatBySlot float64 `json:"statBySlot"`
		Ties       int     `json:"ties"`
		Wins       int     `json:"wins"`
	} `json:"cumulativeScore"`
	RosterForCurrentScoringPeriod struct {
		AppliedStatTotal float64 `json:"appliedStatTotal"`
		Entries          []struct {
			LineupSlotID    int `json:"lineupSlotId"`
			PlayerPoolEntry struct {
				Player struct {
					Stats []struct {
						AppliedStats    map[string]float64 `json:"appliedStats"`
						AppliedTotal    float64            `json:"appliedTotal"`
						ProTeamID       int                `json:"proTeamId"`
						ScoringPeriod   int                `json:"scoringPeriod"`
						SeasonID        int64              `json:"seasonId"`
						StatSourceID    int                `json:"statSourceId"`
						StatSplitTypeID int                `json:"statSplitTypeId"`
						Stats           map[string]float64 `json:"stats"`
						Variance        map[string]float64 `json:"variance"`
					} `json:"stats"`
				} `json:"player"`
			} `json:"playerPoolEntry"`
		} `json:"entries"`
	} `json:"rosterForCurrentScoringPeriod"`
	RosterForMatchupPeriodDelayed struct {
		Entries []any `json:"entries"`
	} `json:"rosterForMatchupPeriodDelayed"`
	TeamID          int     `json:"teamId"`
	TieBreak        int     `json:"tiebreak"`
	TotalPoints     float64 `json:"totalPoints"`
	TotalPointsLive float64 `json:"totalPointsLive"`
}

type MMAtchupScore struct {
	DraftDetail struct {
		Drafted    bool `json:"drafted"`
		InProgress bool `json:"inProgress"`
	} `json:"draftDetail"`
	GameID   int   `json:"gameId"`
	LeagueID int64 `json:"leagueId"`
	Schedule []struct {
		Away struct {
			ScheduleStruct ScheduleStruct
		} `json:"away"`
		Home struct {
			ScheduleStruct ScheduleStruct
		} `json:"home"`
		Id              int    `json:"id"`
		MatchupPeriodID int    `json:"matchupPeriodId"`
		PlayoffTierType string `json:"playoffTierType"`
		Winner          string `json:"winner"`
	} `json:"schedule"`
	ScoringPeriodID int   `json:"scoringPeriod"`
	SeasonID        int64 `json:"seasonId"`
	SegmentID       int   `json:"segmentId"`
	Status          struct {
		ActivatedDate            int64          `json:"activatedDate"`
		CreatedLeagueType        int            `json:"createdLeagueType"`
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
}
