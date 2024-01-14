package summoner

import (
	. "meep.gg/protos/riot/lol/summoner/v1"
	. "meep.gg/protos/scylla/lol/summoner/v1"
)

func SerializeRiotSummoner(summoner *RiotSummoner) *ScyllaSummoner {
	return &ScyllaSummoner{
		Puuid:         summoner.Puuid,
		Id:            summoner.Id,
		Name:          summoner.Name,
		ProfileIconId: summoner.ProfileIconId,
		SummonerLevel: summoner.SummonerLevel,
	}
}
