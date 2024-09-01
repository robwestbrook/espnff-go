package espnffgo

type MDraftDetail struct {
	DraftDetail struct {
		CompleteDate int64 `json:"completeDate"`
		Drafted      bool  `json:"drafted"`
		InProgress   bool  `json:"inProgress"`
		Picks        []struct {
			AutoDraftTypeID   int     `json:"autoDraftTypeId"`
			BidAmount         float64 `json:"bidAmount"`
			ID                int     `json:"id"`
			Keeper            bool    `json:"keeper"`
			LineupSlotID      int     `json:"lineupSlotId"`
			MemberID          string  `json:"memberId"`
			NominatingTeamID  int     `json:"nominatingTeamId"`
			OverallPickNumber int     `json:"overallPickNumber"`
			PlayerID          int64   `json:"playerId"`
			ReservedForKeeper bool    `json:"reservedForKeeper"`
			RoundID           int     `json:"roundId"`
			RoundPickNumber   int     `json:"roundPickNumber"`
			TeamID            int     `json:"teamId"`
			TradeLocked       bool    `json:"tradeLocked"`
		} `json:"picks"`
	} `json:"draftDetail"`
	GameID          int   `json:"gameId"`
	LeagueID        int64 `json:"leagueId"`
	ScoringPeriodID int   `json:"scoringPeriodId"`
	SeasonID        int64 `json:"seasonId"`
	SegmentID       int   `json:"segmentId"`
	Settings        struct {
		DraftSettings struct {
			AuctionBudget     float64 `json:"auctionBudget"`
			AvailableDate     int64   `json:"availableDate"`
			Date              int64   `json:"date"`
			IsTradingEnabled  bool    `json:"isTradingEnabled"`
			KeeperCount       int     `json:"keeperCount"`
			KeeperCountFuture int     `json:"keeperCountFuture"`
			KeeperOrderType   string  `json:"keeperOrderType"`
			LeagueSubType     string  `json:"leagueSubType"`
			OrderType         string  `json:"orderType"`
			PickOrder         []int   `json:"pickOrder"`
			TimePerSelection  int     `json:"timePerSelection"`
			Type              string  `json:"type"`
		} `json:"draftSettings"`
	} `json:"settings"`
	Status struct {
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
