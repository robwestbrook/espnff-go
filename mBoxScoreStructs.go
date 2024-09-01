package espnffgo

type CumulativeStat struct {
	Ineligible bool        `json:"ineligible"`
	Rank       int         `json:"rank"`
	Result     interface{} `json:"result"`
	Score      float64     `json:"score"`
}

type Games struct {
	Away struct {
		TeamID      int     `json:"teamId"`
		TieBreak    int     `json:"tiebreak"`
		TotalPoints float64 `json:"totalPoints"`
	} `json:"away"`
	Home struct {
		TeamID      int     `json:"teamId"`
		TieBreak    int     `json:"tiebreak"`
		TotalPoints float64 `json:"totalPoints"`
	} `json:"home"`
	ID              int `json:"id"`
	MatchupPeriodID int `json:"matchupPeriodId"`
}

type MBoxscore struct {
	DraftDetail struct {
		Drafted    bool `json:"drafted"`
		InProgress bool `json:"inProgress"`
	} `json:"draftDetail"`
	GameID   int   `json:"gameId"`
	LeagueID int64 `json:"id"`
	Schedule []struct {
		Away struct {
			CumulativeScore struct {
				Losses      int                        `json:"losses"`
				ScoreByStat map[string]ScoreByStatData `json:"scoreByStat"`
				StatBySlot  float64                    `json:"statBySlot"`
				Ties        int                        `json:"ties"`
				Wins        int                        `json:"wins"`
			} `json:"cumulativeScore"`
			RosterForCurrentScoringPeriod struct {
				AppliedStatTotal float64 `json:"appliedStatTotal"`
				Entries          []struct {
					LineupSlotID    int   `json:"lineupSlotId"`
					PlayerID        int64 `json:"playerId"`
					PlayerPoolEntry struct {
						AppliedStatTotal float64 `json:"appliedStatTotal"`
						ID               int64   `json:"id"`
						Player           struct {
							DefaultPositionID int    `json:"defaultPositionId"`
							EligibleSlots     []int  `json:"eligibleSlots"`
							FirstName         string `json:"firstName"`
							FullName          string `json:"fullName"`
							ID                int64  `json:"id"`
							LastName          string `json:"lastName"`
							ProteamID         int    `json:"proTeamId"`
							Stats             []struct {
								AppliedStats    map[string]float64 `json:"appliedStats"`
								AppliedTotal    float64            `json:"appliedTotal"`
								ExternalID      string             `json:"externalId"`
								ID              string             `json:"id"`
								ProteamID       int                `json:"proTeamId"`
								ScoringPeriodID int                `json:"scoringPeriodId"`
								SeasonID        int64              `json:"seasonId"`
								StatSourceID    int                `json:"statSourceId"`
								StatSplitTypeID int                `json:"statSplitTypeId"`
								Stats           map[string]float64 `json:"stats"`
							} `json:"stats"`
						} `json:"player"`
					} `json:"playerPoolEntry"`
					UniverseID int `json:"universeId"`
				} `json:"entries"`
			} `json:"rosterForCurrentScoringPeriod"`
			RosterForMatchupPeriod struct {
				Entries []any `json:"entries"`
			} `json:"rosterForMatchupPeriod"`
			RosterForMatchupPeriodDelayed struct {
				Entries []any `json:"entries"`
			} `json:"rosterForMatchupPeriodDelayed"`
			TeamID      int     `json:"teamId"`
			TieBreak    float64 `json:"tieBreak"`
			TotalPoints float64 `json:"totalPoints"`
		} `json:"away"`
		Home struct {
			CumulativeScore struct {
				Losses      int                        `json:"losses"`
				ScoreByStat map[string]ScoreByStatData `json:"scoreByStat"`
				StatBySlot  float64                    `json:"statBySlot"`
				Ties        int                        `json:"ties"`
				Wins        int                        `json:"wins"`
			} `json:"cumulativeScore"`
			RosterForCurrentScoringPeriod struct {
				AppliedStatTotal float64 `json:"appliedStatTotal"`
				Entries          []struct {
					LineupSlotID    int   `json:"lineupSlotId"`
					PlayerID        int64 `json:"playerId"`
					PlayerPoolEntry struct {
						AppliedStatTotal float64 `json:"appliedStatTotal"`
						ID               int64   `json:"id"`
						Player           struct {
							DefaultPositionID int    `json:"defaultPositionId"`
							EligibleSlots     []int  `json:"eligibleSlots"`
							FirstName         string `json:"firstName"`
							FullName          string `json:"fullName"`
							ID                int64  `json:"id"`
							LastName          string `json:"lastName"`
							ProteamID         int    `json:"proTeamId"`
							Stats             []struct {
								AppliedStats    map[string]float64 `json:"appliedStats"`
								AppliedTotal    float64            `json:"appliedTotal"`
								ExternalID      string             `json:"externalId"`
								ID              string             `json:"id"`
								ProteamID       int                `json:"proTeamId"`
								ScoringPeriodID int                `json:"scoringPeriodId"`
								SeasonID        int64              `json:"seasonId"`
								StatSourceID    int                `json:"statSourceId"`
								StatSplitTypeID int                `json:"statSplitTypeId"`
								Stats           map[string]float64 `json:"stats"`
							} `json:"stats"`
						} `json:"player"`
					} `json:"playerPoolEntry"`
					UniverseID int `json:"universeId"`
				} `json:"entries"`
			} `json:"rosterForCurrentScoringPeriod"`
			RosterForMatchupPeriod struct {
				Entries []any `json:"entries"`
			} `json:"rosterForMatchupPeriod"`
			RosterForMatchupPeriodDelayed struct {
				Entries []any `json:"entries"`
			} `json:"rosterForMatchupPeriodDelayed"`
			TeamID      int     `json:"teamId"`
			TieBreak    float64 `json:"tieBreak"`
			TotalPoints float64 `json:"totalPoints"`
		} `json:"home"`
		ID              int `json:"id"`
		MatchupPeriodID int `json:"matchupPeriodId"`
		Games           Games
	} `json:"schedule"`
	ScoringPeriod int   `json:"scoringPeriodId"`
	SeasonID      int64 `json:"seasonId"`
	SegmentID     int   `json:"segmentId"`
	Settings      struct {
		IsCustomizable   bool `json:"isCustomizable"`
		IsPublic         bool `json:"isPublic"`
		ScheduleSettings struct {
			Divisions []struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"divisions"`
			MatchupPeriods map[string][]int `json:"matchupPeriods"`
			PeriodTypeID   int              `json:"periodTypeId"`
		} `json:"scheduleSettings"`
		ScoringSettings struct {
			ScoringItems []struct {
				IsReverseItem bool    `json:"isReverseItem"`
				Points        float64 `json:"points"`
				StatID        int     `json:"statId"`
			} `json:"scoringItems"`
		} `json:"scoringSettings"`
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
	Teams []struct {
		Abbrev              string `json:"abbrev"`
		DivisionID          int    `json:"divisionId"`
		TeamID              int    `json:"id"`
		Logo                string `json:"logo"`
		Name                string `json:"name"`
		RankCalculatedFinal int    `json:"rankCalculatedFinal"`
		Record              struct {
			Overall struct {
				Losses int `json:"losses"`
				Ties   int `json:"ties"`
				Wins   int `json:"wins"`
			} `json:"overall"`
		} `json:"record"`
	} `json:"teams"`
}
