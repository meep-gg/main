package league

import (
	. "meep.gg/protos/riot/lol/league/v1"
	. "meep.gg/protos/scylla/lol/league/v1"
	. "meep.gg/types/lol"
)

func SerializeEntries(Leagues []*RiotLeagueEntry) []*ScyllaLeagueEntry {
	var entries []*ScyllaLeagueEntry
	for _, league := range Leagues {
		entries = append(entries, SerializeEntry(league))
	}
	return entries
}

func SerializeEntry(league *RiotLeagueEntry) *ScyllaLeagueEntry {
	current, err := SerializeRank(league.Tier, league.Rank, league.LeaguePoints)
	if err != nil {
		return nil
	}
	return &ScyllaLeagueEntry{
		SummonerId:   league.SummonerId,
		QueueType:    league.QueueType,
		Tier:         league.Tier,
		Rank:         league.Rank,
		LeaguePoints: league.LeaguePoints,
		Wins:         league.Wins,
		Losses:       league.Losses,
		Current:      current,
	}
}

func SerializeItem(tier, rank, queueType string, item *LeagueItem) *ScyllaLeagueEntry {
	current, err := SerializeRank(tier, rank, item.LeaguePoints)
	if err != nil {
		return nil
	}
	if !IsQueue(queueType) {
		return nil
	}
	return &ScyllaLeagueEntry{
		SummonerId:   item.SummonerId,
		QueueType:    queueType,
		Tier:         tier,
		Rank:         rank,
		LeaguePoints: item.LeaguePoints,
		Wins:         item.Wins,
		Losses:       item.Losses,
		Current:      current,
	}
}
