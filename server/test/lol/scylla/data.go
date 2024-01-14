package scylla

import (
	. "meep.gg/protos/scylla/lol/league/v1"
	. "meep.gg/protos/scylla/lol/match/v1"
	. "meep.gg/protos/scylla/lol/participant/v1"
	. "meep.gg/protos/scylla/lol/participantdetail/v1"
	. "meep.gg/protos/scylla/lol/profile/v1"
	. "meep.gg/protos/scylla/lol/summoner/v1"
)

var LolScyllaAddr = ":50062"

var TestSummoner = &ScyllaSummoner{
	Puuid:         "9WX5pU7pGi8EQvSFfC5hsiDJRpqdvmp1Lf57qbq1zSpSIiz38oakVBcFR_dMC2Rd6nbYf4n_pLI8sw",
	Id:            "C-dE3eK53OefJnIPW58zvr2Qjo2fiKxFG_aA5zip9XjT7MtP",
	Name:          "Leong",
	ProfileIconId: 5453,
	SummonerLevel: 341,
}

var TestLeague = &ScyllaLeagueEntry{
	SummonerId:   "C-dE3eK53OefJnIPW58zvr2Qjo2fiKxFG_aA5zip9XjT7MtP",
	QueueType:    "RANKED_SOLO_5x5",
	LeaguePoints: 75,
	Losses:       259,
	Rank:         "IV",
	Tier:         "EMERALD",
	Wins:         254,
	Current:      2075,
}

var TestMatch = &ScyllaMatch{
	GameId:           4817867075,
	GameEndTimestamp: 1698884045813,
	Season:           13,
	Patch:            21,
	QueueId:          420,
	GameDuration:     1773,
	Rank:             2000,
	Bans: []int32{
		200, 86, 143, 105, 53, 51, 518, 11, 58, 233,
	},
	Participants: []*ScyllaMatchParticipant{
		{
			Puuid:      "oH6WSYolMK8CEo9XcTahbceAeCXfRBdVZPKoyr83ugstZmPLkiq1bUOmzSn1_c6ga9r-CbOxmlp0Kg",
			GameName:   "Roxol",
			ChampionId: 78,
		},
		{
			Puuid:      "9WX5pU7pGi8EQvSFfC5hsiDJRpqdvmp1Lf57qbq1zSpSIiz38oakVBcFR_dMC2Rd6nbYf4n_pLI8sw",
			GameName:   "Leong",
			ChampionId: 28,
		},
		{
			Puuid:      "O0YKqmCl-GFHbs8fOPq7juVFUBf7xGLM-ZvtkQnS1eGdPplDZa_ifoht2QAwDAPrOabWXqgKfR3jnw",
			GameName:   "mTGordon",
			ChampionId: 45,
		},
		{
			Puuid:      "hCOg79FnFLZlN5QHL1FRdTZ35B4h16vfIdC68Q64S09my-cX721dryiCuRSIb4tm-2DGRr8s3RRFoA",
			GameName:   "Nqrwhql",
			ChampionId: 21,
		},
		{
			Puuid:      "8fFX65cXqDvkHSCbB6neUAli-uvAjxi3Hy49mLc7MyuGCo-NMc5MGIVSf5H1pt9AGMdvYyre4Kz88w",
			GameName:   "Spensir",
			ChampionId: 235,
		},
		{
			Puuid:      "2-sjpE33aGB7UnRrX6FRs4u-JsbNaiajvdBT2ZA5DocYypk6ZCdOtcg0OpUPPBJn7C832Xzn9do3Jw",
			GameName:   "Vuramix",
			ChampionId: 122,
		},
		{
			Puuid:      "7XXTxSdi-ecbGjbtmWiXheRJk-O6la9wbCQrcKvIaKN2ZvYJolSSdI9lwnOqszeFnlfV6UYc9Y405Q",
			GameName:   "BlakeBMD",
			ChampionId: 876,
		},
		{
			Puuid:      "QahIDL9OgC38kDQQepHlxI9d-yqLLNftXKDXTqq_gfKUiEmxvUkP1nigcEebFCzdp1iWHNKsngNHHw",
			GameName:   "Goulairith",
			ChampionId: 91,
		},
		{
			Puuid:      "Rdn9DNU-6lQFngINlFRV10-KNhhasApVKXw2CB6nOOMAT8BfEpTP81ZUPiWgVtYkWmZ4Sklo3uIzlg",
			GameName:   "French Horn",
			ChampionId: 115,
		},
		{
			Puuid:      "LuPDZzoTIj1O6Sqh-aWhawl-FRTg3uUvIUa28MljgV2mSXegIiDqdprYhhwX7F_p3zV7XEgdnDWaJg",
			GameName:   "Yeetothehaw",
			ChampionId: 43,
		},
	},
}

var TestParticipant = &ScyllaParticipant{
	// primary
	Puuid:                       "oH6WSYolMK8CEo9XcTahbceAeCXfRBdVZPKoyr83ugstZmPLkiq1bUOmzSn1_c6ga9r-CbOxmlp0Kg",
	GameId:                      4817867075,
	ChampionId:                  78,
	ChampLevel:                  15,
	Kda:                         0.64,
	Kills:                       3,
	Deaths:                      11,
	Assists:                     4,
	GoldEarned:                  9776,
	GoldPerMinute:               330.69,
	TotalMinionsKilled:          164,
	TotalDamageDealtToChampions: 20281,
	DamagePerMinute:             686.02,
	TotalDamageTaken:            32988,
	TimeCCingOthers:             39,
	IndividualPosition:          "TOP",
	Rank:                        4,
	Win:                         false,
	ParticipantId:               1,
	Item0:                       2033,
	Item1:                       3047,
	Item2:                       3066,
	Item3:                       2055,
	Item4:                       3075,
	Item5:                       6632,
	Item6:                       3340,
	// Placement:                   0,
	// PlayerAugment1:              0,
	// PlayerAugment2:              0,
	// PlayerAugment3:              0,
	// PlayerAugment4:              0,
	// PlayerSubteamId:             0,
	// Position:                    0,
	TeamId:               100,
	Summoner1Id:          12,
	Summoner2Id:          4,
	PrimaryStyle:         8437,
	SubStyle:             8009,
	KillParticipation:    0.24,
	VisionScorePerMinute: 0.47,
	TeamDamagePercentage: 0.22,
	NeutralMinionsKilled: 0,
}

var TestParticipantdetail = &ScyllaParticipantdetail{
	GameId:                         4817867075,
	ParticipantId:                  1,
	DamageDealtToBuildings:         2154,
	DamageDealtToObjectives:        4646,
	DamageDealtToTurrets:           2154,
	DamageSelfMitigated:            37905,
	MagicDamageDealt:               10429,
	MagicDamageDealtToChampions:    3719,
	MagicDamageTaken:               9883,
	NeutralMinionsKilled:           0,
	Perks:                          TestPerks,
	PhysicalDamageDealt:            97533,
	PhysicalDamageDealtToChampions: 16562,
	PhysicalDamageTaken:            19711,
	Spell1Casts:                    113,
	Spell2Casts:                    25,
	Spell3Casts:                    35,
	Spell4Casts:                    7,
	Summoner1Casts:                 2,
	Summoner2Casts:                 4,
	DamageTakenOnTeamPercentage:    0.23,
	MaxCsAdvantageOnLaneOpponent:   5,
	MaxLevelLeadLaneOpponent:       2,
}

var TestPerks = &ScyllaPerks{
	StatPerks: &ScyllaPerkStats{
		Defense: 5001,
		Flex:    5002,
		Offense: 5008,
	},
	Styles: []*ScyllaPerkStyle{
		{
			Description: "primaryStyle",
			Selections: []*ScyllaPerkStyleSelection{
				{
					Perk: 8437,
					Var1: 1117,
					Var2: 694,
					Var3: 0,
				},
				{
					Perk: 8401,
					Var1: 182,
					Var2: 0,
					Var3: 0,
				},
				{
					Perk: 8473,
					Var1: 949,
					Var2: 0,
					Var3: 0,
				},
				{
					Perk: 8451,
					Var1: 204,
					Var2: 0,
					Var3: 0,
				},
			},
			Style: 8400,
		},
		{
			Description: "subStyle",
			Selections: []*ScyllaPerkStyleSelection{
				{
					Perk: 8009,
					Var1: 1629,
					Var2: 0,
					Var3: 0,
				},
				{
					Perk: 8299,
					Var1: 746,
					Var2: 0,
					Var3: 0,
				},
			},
			Style: 8000,
		},
	},
}

var TestProfile = &ScyllaProfile{
	Puuid:         "9WX5pU7pGi8EQvSFfC5hsiDJRpqdvmp1Lf57qbq1zSpSIiz38oakVBcFR_dMC2Rd6nbYf4n_pLI8sw",
	SummonerId:    "C-dE3eK53OefJnIPW58zvr2Qjo2fiKxFG_aA5zip9XjT7MtP",
	GameName:      "Leong",
	TagLine:       "NA1",
	ProfileIconId: 5453,
	SummonerLevel: 341,
}
