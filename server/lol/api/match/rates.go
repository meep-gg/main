package lol

import (
	"time"

	. "meep.gg/rates"
)

var rMatch string = "lol-match-"

var MatchPuuidRates = []Limit{
	{
		Name:      rMatch + "match",
		Reservoir: 2000,
		Interval:  10*time.Second + 1*time.Millisecond,
	},
}

var MatchIdRates = []Limit{
	{
		Name:      rMatch + "id",
		Reservoir: 2000,
		Interval:  10*time.Second + 1*time.Millisecond,
	},
}

var MatchTimelineRates = []Limit{
	{
		Name:      rMatch + "timeline",
		Reservoir: 2000,
		Interval:  10*time.Second + 1*time.Millisecond,
	},
}
