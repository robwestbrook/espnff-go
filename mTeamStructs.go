package espnffgo

// NotificationSettings struct for members notification
// settings retrieved from the ESPN "?view=mTeam" endpoint.
type NotificationSettings struct {
	Enabled bool   `json:"enabled"`
	ID      string `json:"id"`
	Type    string `json:"type"`
}

// Record struct for team records retrieved from the
// ESPM "?view=mTeam" endpoint.
type Record struct {
	GamesBack     float64 `json:"gamesBack"`
	Losses        int     `json:"losses"`
	Percentage    float64 `json:"percentage"`
	PointsAgainst float64 `json:"pointsAgainst"`
	PointsFor     float64 `json:"pointsFor"`
	StreakLength  int     `json:"streakLength"`
	StreakType    string  `json:"streakType"`
	Ties          int     `json:"ties"`
	Wins          int     `json:"wins"`
}

// MTeam defines a struct to store the ESPN API response
// to the params "?view=mTeam"
type MTeam struct {
	DraftDetail struct {
		Drafted    bool `json:"drafted"`
		InProgress bool `json:"inProgress"`
	} `json:"draftDetail"`
	GameID          int   `json:"gameId"`
	LeagueID        int64 `json:"id"`
	ScoringPeriodID int   `json:"scoringPeriodId"`
	SeasonID        int64 `json:"seasonId"`
	SegmentID       int   `int:"segmentId"`
	Status          struct {
		ActivatedDate            int64          `json:"activatedDate"`
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
		PreviousSeasons          []int          `json:"previousSeasons"`
		StandingsUpdateDate      int64          `json:"standingsUpdateDate"`
		TeamsJoined              int            `json:"teamsJoined"`
		TransactionScoringPeriod int            `json:"transactionScoringPeriod"`
		WaiverLastExecutionDate  int64          `json:"waiverLastExecutionDate"`
		WaiverProcessStatus      map[string]int `json:"waiverProcessStatus"`
	} `json:"status"`
	Members []struct {
		DisplayName          string `json:"displayName"`
		FirstName            string `json:"firstName"`
		ID                   string `json:"id"`
		LastName             string `json:"lastName"`
		NotificationSettings []struct {
			NotificationSettings
		} `json:"notificationSettings"`
	} `json:"members"`
	Teams []struct {
		Abbrev                string `json:"abbrev"`
		CurrentProjectedRank  int    `json:"currentProjectedRank"`
		DivisionID            int    `json:"divisionId"`
		DraftDayProjectedRank int    `json:"draftDayProjectedRank"`
		DraftStrategy         struct {
			KeeperPlayerIDs []struct{ int } `json:"keeperPlayerIds"`
		} `json:"draftStrategy"`
		MemberID       int      `json:"memberId"`
		IsActive       bool     `json:"isActive"`
		Logo           string   `json:"logo"`
		LogoType       string   `json:"logoType"`
		Name           string   `json:"name"`
		Owners         []string `json:"owners"`
		PlayoffSeed    int      `json:"playoffSeed"`
		Points         float64  `json:"points"`
		PointsAdjusted float64  `json:"pointsAdjusted"`
		PointsDelta    float64  `json:"pointsDelta"`
		PrimaryOwner   string   `json:"primaryOwner"`
		RankFinal      int      `json:"rankFinal"`
		Record         struct {
			Away struct {
				Record
			} `json:"away"`
			Division struct {
				Record
			} `json:"division"`
			Home struct {
				Record
			} `json:"home"`
			Overall struct {
				Record
			} `json:"overall"`
		} `json:"record"`
		TradeBlock struct {
			Players map[string]string `json:"players"`
		} `json:"tradeBlock"`
		TransactionCounter struct {
			AcquisitionBudgetSpent   float64 `json:"acquisitionBudgetSpent"`
			Acquisitions             int     `json:"acquisitions"`
			Drops                    int     `json:"drops"`
			MatchupAcquisitionTotals struct {
				MatchupAcquisitionTotals map[string]int `json:"matchupAcquisitionTotals"`
			}
			Misc         int     `json:"misc"`
			MoveToActive int     `json:"moveToActive"`
			MoveToIR     int     `json:"moveToIR"`
			Paid         float64 `json:"paid"`
			TeamChanges  int     `json:"teamChanges"`
			Trades       int     `json:"trades"`
		} `json:"transactionCounter"`
		ValuesByStat struct {
			ValuesByStat map[string]int `json:"valuesByStat"`
		}
		WaiverRank int `json:"waiverRank"`
	} `json:"teams"`
}
