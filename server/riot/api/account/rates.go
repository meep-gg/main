package account

import (
	"time"

	rates "meep.gg/rates"
)

var rAccount string = "riot-account-"

// GET /riot/account/v1/accounts/by-puuid/{puuid}
var PuuidRates = []rates.Limit{
	{
		Name:      rAccount + "puuid",
		Reservoir: 1000,
		Interval:  time.Minute,
	},
}

// GET /riot/account/v1/accounts/by-riot-id/{gameName}/{tagLine}
var RiotIdRates = []rates.Limit{
	{
		Name:      rAccount + "riotid",
		Reservoir: 1000,
		Interval:  time.Minute,
	},
}
