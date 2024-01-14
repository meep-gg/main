package lol

import (
	"time"

	. "meep.gg/rates"
)

var rSummoner string = "lol-summoner-"

var SummonerAccountRates = []Limit{
	{
		Name:      rSummoner + "account",
		Reservoir: 1600,
		Interval:  time.Minute,
	},
}

var SummonerNameRates = []Limit{
	{
		Name:      rSummoner + "name",
		Reservoir: 1600,
		Interval:  time.Minute,
	},
}

var SummonerPuuidRates = []Limit{
	{
		Name:      rSummoner + "puuid",
		Reservoir: 1600,
		Interval:  time.Minute,
	},
}

var SummonerIdRates = []Limit{
	{
		Name:      rSummoner + "id",
		Reservoir: 1600,
		Interval:  time.Minute,
	},
}
