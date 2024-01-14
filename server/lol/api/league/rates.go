package lol

import (
	"time"

	. "meep.gg/rates"
)

var rLeague string = "lol-league-"
var LeagueChallengerRates = []Limit{
	{
		Name:      rLeague + "challenger-sb",
		Reservoir: 30,
		Interval:  10*time.Second + 1*time.Millisecond,
	},
	{
		Name:      rLeague + "challenger-lb",
		Reservoir: 500,
		Interval:  10 * time.Minute,
	},
}

var LeagueGrandmasterRates = []Limit{
	{
		Name:      rLeague + "grandmaster-sb",
		Reservoir: 30,
		Interval:  10*time.Second + 1*time.Millisecond,
	},
	{
		Name:      rLeague + "grandmaster-lb",
		Reservoir: 500,
		Interval:  10 * time.Minute,
	},
}

var LeagueMasterRates = []Limit{
	{
		Name:      rLeague + "master-sb",
		Reservoir: 30,
		Interval:  10*time.Second + 1*time.Millisecond,
	},
	{
		Name:      rLeague + "master-lb",
		Reservoir: 500,
		Interval:  10 * time.Minute,
	},
}

var LeagueSummonerIdRates = []Limit{
	{
		Name:      rLeague + "summoner-id",
		Reservoir: 100,
		Interval:  time.Minute,
	},
}

var LeagueEntryRates = []Limit{
	{
		Name:      rLeague + "entry",
		Reservoir: 50,
		Interval:  10 * time.Second,
	},
}

var LeagueIdRates = []Limit{
	{
		Name:      rLeague + "id",
		Reservoir: 500,
		Interval:  10 * time.Second,
	},
}
