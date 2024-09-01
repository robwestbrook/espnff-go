package espnffgo

type RatingsByOpponent struct{}

type PositionalRating struct {
	Average           float64           `json:"average"`
	RatingsByOpponent RatingsByOpponent `json:"ratingsByOpponent"`
}

type PositionalRatings struct {
	PositionalRatings map[string]PositionalRating `json:"positionalRatings"`
}

type KonaPlayerInfo struct {
	Players []struct {
		DraftAuctionValue float64 `json:"draftAuctionValue"`
		ID                int64   `json:"id"`
		KeeperValue       float64 `json:"keeperValue"`
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
			ID            int64  `json:"id"`
			Injured       bool   `json:"injured"`
			InjuryStatus  string `json:"injuryStatus"`
			Jersey        string `json:"jersey"`
			LastName      string `json:"lastName"`
			LastNewsDate  int64  `json:"lastNewsDate"`
			Ownership     struct {
				ActivityLevel                     int64   `json:"activityLevel"`
				AuctionValueAverage               float64 `json:"auctionValueAverage"`
				AuctionValueAverageChange         float64 `json:"auctionValueAverageChange"`
				AverageDraftPosition              float64 `json:"averageDraftPosition"`
				AverageDraftPositionPercentChange float64 `json:"averageDraftPositionPercentChange"`
				Date                              int64   `json:"date"`
				LeagueType                        int     `json:"leagueType"`
				PercentChange                     float64 `json:"percentChange"`
				PercentOwned                      float64 `json:"percentOwned"`
				PercentStarted                    float64 `json:"percentStarted"`
			} `json:"ownership"`
			ProTeamID     int                      `json:"proTeamId"`
			Rankings      map[string][]interface{} `json:"rankings"`
			SeasonOutlook string                   `json:"seasonOutlook"`
			Stats         []struct {
				AppliedAverage  float64            `json:"appliedAverage"`
				AppliedTotal    float64            `json:"appliedTotal"`
				ExternalID      string             `json:"eternalId"`
				ID              string             `json:"id"`
				ProTeamID       int                `json:"proTeamId"`
				ScoringPeriodID int                `json:"scoringPeriodId"`
				SeasonID        int                `json:"seasonId"`
				StatSourceID    int                `json:"statSourceId"`
				StatSplitTypeID int                `json:"statSplitTypeId"`
				Stats           map[string]float64 `json:"stats"`
			} `json:"stats"`
			Ratings      map[string]RankingsPlayer `json:"ratings"`
			RosterLocked bool                      `json:"rosterLocked"`
			Status       string                    `json:"status"`
			TradeLocked  bool                      `json:"tradeLocked"`
		} `json:"player"`
	} `json:"players"`
	PositionAgainstOpponent struct {
		PositionAgainstOpponent PositionalRatings `json:"positionAgainstOpponent"`
	}
}
