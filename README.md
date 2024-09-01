# ESPNFF-GO

A go library for retrieving data from the unofficial ESPN Fantasy Football API.

### ESPN Fantasy Football Endpoints

ESPN has mumerous endpoints. Some provide useful data, while others don't seem to offer any additional info from those used in this library. 

The following is a list of endpoints that are available from the ESPNFF-GO library:

1.   mTeam
2.   mMatchup
3.   mRoster
4.   mSettings
5.   mBoxscore
6.   mMatchupScore
7.   kona_player_info
8.   mDraftDetail
9.   mTransactions2
10.  mStatus
11.  mPositionalRatingsStats

This is a list of endpoints still to be implemented into this library.

1.  mLiveScoring (todo)
2.  kona_league_communication (todo)
3.  mPositionalRatings (todo)

This is a list of endpoints that do not seem to offer any unique data:

1.  player_wl
2.  mSchedule
3.  mScoreboard
4.  mPendingTransactions
5.  modular
6.  mNav
7.  kona_game_state
8.  proTeamSchedules_wl

### Create a .env file

ESPNFF-GO requires a .env containing the following parameters:

1. YEAR - the current year
2. LEAGUE_ID - the fantasy football league ID.
3. ESPN_S2 - the cookie value espn_s2
4. SWID - the cookie value SWID