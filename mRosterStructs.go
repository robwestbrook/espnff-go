package espnffgo

// Ranking struct for team rosters retrieved from the
// ESPM "?view=mRoster" endpoint.
type RankingsRoster struct {
	AuctionValue int     `json:"auctionValue"`
	Published    bool    `json:"published"`
	Rank         float64 `json:"rank"`
	RankSourceID int     `json:"rankSourceId"`
	RankType     string  `json:"rankType"`
	SlotID       int     `json:"slotId"`
}

type RatingsPlayer struct {
	PositionalRanking int     `json:"positionalRanking"`
	TotalRanking      int     `json:"totalRanking"`
	TotalRating       float64 `json:"totalRating"`
}

type MRoster struct {
	DraftDetail struct {
		Drafted    bool `json:"drafted"`
		InProgress bool `json:"inProgress"`
	} `json:"draftDetail"`
	GameID        int   `json:"gameId"`
	LeagueID      int64 `json:"id"`
	ScoringPeriod int   `json:"scoringPeriod"`
	SeasonID      int   `json:"seasonId"`
	SegmentID     int   `json:"segmentId"`
	Status        struct {
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
		PreviousSeasons          []int          `json:"previousSeasons"`
		StandingsUpdateDate      int64          `json:"standingsUpdateDate"`
		TeamsJoined              int            `json:"teamsJoined"`
		TransactionScoringPeriod int            `json:"transactionScoringPeriod"`
		WaiverLastExecutionDate  int64          `json:"waiverLastExecutionDate"`
		WaiverProcessStatus      map[string]int `json:"waiverProcessStatus"`
	} `json:"status"`
	Teams []struct {
		TeamID int `json:"id"`
		Roster struct {
			AppliedStatTotal float64 `json:"appliedStatTotal"`
			Entries          []struct {
				AcquisitionDate       int64    `json:"acquisitionDate"`
				AcquisitionType       string   `json:"acquisitionType"`
				InjuryStatus          string   `json:"injuryStatus"`
				LineupSlotID          int      `json:"lineupSlotId"`
				PendingTransactionIDs []string `json:"pendingTransactionIds"`
				PlayerID              int64    `json:"playerId"`
				PlayerPoolEntry       struct {
					AppliedStatTotal  float64 `json:"appliedStatTotal"`
					PlayerID          int64   `json:"id"`
					KeeperValue       float64 `json:"keeprValue"`
					KeeperValueFuture float64 `json:"keeperValueFuture"`
					LineupLocked      bool    `json:"lineupLocked"`
					OnTeamID          int     `json:"onTeamId"`
					Player            struct {
						Active               bool `json:"active"`
						DefaultPositionID    int  `json:"defaultPositionId"`
						DraftRanksByRankType struct {
							Standard struct {
								AuctionValue float64 `json:"auctionValue"`
								Published    bool    `json:"published"`
								Rank         int     `json:"rank"`
								RankSourceID int     `json:"rankSourceId"`
								RankType     string  `json:"rankType"`
								SlotID       int     `json:"slotId"`
							} `json:"STANDARD"`
							PPR struct {
								AuctionValue float64 `json:"auctionValue"`
								Published    bool    `json:"published"`
								Rank         int     `json:"rank"`
								RankSourceID int     `json:"rankSourceId"`
								RankType     string  `json:"rankType"`
								SlotID       int     `json:"slotId"`
							} `json:"PPR"`
						} `json:"draftRanksByRankType"`
						Droppable     bool   `json:"droppable"`
						EligibleSlots []int  `json:"eligibleSlots"`
						FirstName     string `json:"firstName"`
						FullName      string `json:"fullName"`
						PlayerID      int64  `json:"id"`
						Injured       bool   `json:"injured"`
						InjuryStatus  string `json:"injuryStatus"`
						LastName      string `json:"lastName"`
						LastNewsDate  int64  `json:"lastNewsDate"`
						LastVideoDate int64  `json:"lastVideoDate"`
						Ownership     struct {
							AuctionValueAverage  float64 `json:"auctionValueAverage"`
							AverageDraftPosition float64 `json:"averageDraftPosition"`
							PercentChange        float64 `json:"percentChange"`
							PercentOwned         float64 `json:"percentOwned"`
							PercentStarted       float64 `json:"percentStarted"`
						} `json:"ownership"`
						ProTeamID     int                         `json:"proTeamId"`
						Rankings      map[string][]RankingsRoster `json:"rankings"`
						SeasonOutlook string                      `json:"seasonOutlook"`
						Stats         []struct {
							AppliedAverage  float64            `json:"appliedAverage"`
							AppliedStats    map[string]float64 `json:"appliedStats"`
							AppliedTotal    float64            `json:"appliedTotal"`
							ExternalID      string             `json:"externalId"`
							ID              string             `json:"id"`
							ProTeamID       int                `json:"proTeamId"`
							ScoringPeriodID int                `json:"scoringPeriodId"`
							SeasonID        int                `json:"seasonId"`
							StatSourceID    int                `json:"statSourceId"`
							StatSplitTypeID int                `json:"statSplitTypeId"`
							Stats           map[string]float64 `json:"stats"`
						} `json:"stats"`
						UniverseID int `json:"universeId"`
					} `json:"player"`
					Ratings      map[string]RatingsPlayer `json:"ratings"`
					RosterLocked bool                     `json:"rosterLocked"`
					Status       string                   `json:"status"`
					TradeLocked  bool                     `json:"tradeLocked"`
				} `json:"playerPoolEntry"`
				Status string `json:"status"`
			} `json:"entries"`
		} `json:"roster"`
	} `json:"teams"`
}
