package espnffgo

type MatchupPeriodStruct struct {
	MatchupPeriods map[string][]int `json:"matchupPeriods"`
}

type MSettings struct {
	DraftDetail struct {
		Drafted    bool `json:"draftDetail"`
		InProgress bool `json:"inProgress"`
	} `json:"draftDetail"`
	GameID        int `json:"gameId"`
	ScoringPeriod int `json:"scoringPeriod"`
	SeasonID      int `json:"seasonId"`
	SegmentID     int `json:"segmentId"`
	Settings      struct {
		AcquisitionSettings struct {
			AcquisitionBudget             float64  `json:"acquisitionBudget"`
			AcquisitionLimit              float64  `json:"acquisitionLimit"`
			AcquisitionType               string   `json:"acquisitionType"`
			FinalPlaceTransactionEligible int      `json:"finalPlaceTransactionEligible"`
			IsUsingAcquisitionBudget      bool     `json:"isUsingAcquisitionBudget"`
			MatchupAcquisitionLimit       float64  `json:"matchupAcquisitionLimit"`
			MatchupLimitPerScoringPeriod  bool     `json:"matchupLimitPerScoringPeriod"`
			MinimumBid                    float64  `json:"minimumBid"`
			TransactionLockingEnabled     bool     `json:"transactionLockingEnabled"`
			WaiverHours                   int      `json:"waiverHours"`
			WaiverOrderReset              bool     `json:"waiverOrderReset"`
			WaiverProcessDays             []string `json:"waiverProcessDays"`
			WaiverProcessHour             int      `json:"waiverProcessHour"`
		} `json:"acquisitionSettings"`
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
			TimePerSelection  int     `json:"timePerSelectiom"`
			Type              string  `json:"type"`
		} `json:"draftSettings"`
		FinancialSettings struct {
			EntryFee           float64 `json:"entryFee"`
			MiscFee            float64 `json:"miscFee"`
			PerLoss            float64 `json:"perLoss"`
			PerTrade           float64 `json:"perTrade"`
			PlayerAcquisition  float64 `json:"playerAcquisition"`
			PlayerDrop         float64 `json:"playerDrop"`
			PlayerMoveToActive float64 `json:"playerMoveToActive"`
			PlayerMoveToIR     float64 `json:"playerMoveToIR"`
		} `json:"financialSettings"`
		IsCustomizable  bool   `json:"isCustomizable"`
		IsPublic        bool   `json:"isPublic"`
		Name            string `json:"name"`
		RestrictionType string `json:"restrictionType"`
		RosterSettings  struct {
			IsBenchUnlimited       bool           `json:"isBenchUnlimited"`
			IsUsingUndroppableList bool           `json:"isUsingUndroppableList"`
			LineupLocktimeType     string         `json:"lineupLocktimeType"`
			LineupSlotCounts       map[string]int `json:"lineupSlotCounts"`
			LineupSlotStatLimits   map[string]int `json:"lineupSlotStatLimits"`
			MoveLimit              float64        `json:"moveLimit"`
			PositionLimits         map[string]int `json:"positionLimits"`
			RosterLocktimeType     string         `json:"rosterLocktimeType"`
			UniversalIDs           []int          `json:"universalIds"`
		} `json:"rosterSettings"`
		ScheduleSettings struct {
			Divisions []struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				Size int    `json:"size"`
			} `json:"divisions"`
			MatchupPeriodCount                 int `json:"matchupPeriodCount"`
			MatchupPeriodLength                int `json:"matchupPeriodLength"`
			MatchupPeriods                     MatchupPeriodStruct
			PeriodTypeID                       int    `json:"periodTypeId"`
			PlayoffMatchupPeriodLength         int    `json:"playoffMatchupPeriodLength"`
			PlayoffReseed                      bool   `json:"playoffReseed"`
			PlayoffSeedingRule                 string `json:"playoffSeedingRule"`
			PlayoffSeedingRuleBy               int    `json:"playoffSeedingRuleBy"`
			PlayoffTeamCount                   int    `json:"playoffTeamCount"`
			VariablePlayoffMatchupPeriodLength bool   `json:"variablePlayoffMatchupPeriodLength"`
		} `json:"scheduleSettings"`
		ScoringSettings struct {
			AllowOutOfPositionScoring bool    `json:"allowOutOfPositionScoring"`
			HomeTeamBonus             float64 `json:"homeTeamBonus"`
			MatchupTieRule            string  `json:"matchupTieRule"`
			MatchupTieRuleBy          int     `json:"matchupTieRuleBy"`
			PlayerRankType            string  `json:"playerRankType"`
			PlayoffHomeTeamBonus      float64 `json:"playoffHomeTeamBonus"`
			PlayoffMatchupTieRule     string  `json:"plyoffMatchupTieRule"`
			PlayoffMatchupTieRuleBy   float64 `json:"playoffMatchupTieRuleBy"`
			ScoringItems              []struct {
				IsReverseItem   bool           `json:"isReverseItem"`
				LeagueRanking   float64        `json:"leagueRanking"`
				LeagueTotal     float64        `json:"leagueTotal"`
				Points          float64        `json:"points"`
				PointsOverrides map[string]int `json:"pointsOverride"`
				StatID          int64          `json:"statId"`
				ScoringType     string         `json:"scoringType"`
			} `json:"scoringItems"`
		} `json:"scoringSettings"`
		Size          int `json:"size"`
		TradeSettings struct {
			AllowOutOfUniverse bool    `json:"allowOutOfUniverse"`
			DeadlineDate       int64   `json:"deadlineDate"`
			Max                float64 `json:"max"`
			RevisionHours      int     `json:"revisionHours"`
			VetoVotesRequired  int     `json:"vetoVotesRequired"`
		} `json:"tradeSettings"`
	} `json:"settings"`
	Status struct {
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
}
