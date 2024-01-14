package league

import (
	"fmt"
	"log"

	. "meep.gg/log"
	. "meep.gg/protos/scylla/lol/league/v1"
	. "meep.gg/scylla"
)

type DBLeague struct {
	SummonerId   string `cql:"summonerid"`
	QueueType    string `cql:"queuetype"`
	LeaguePoints int    `cql:"leaguepoints"`
	Losses       int    `cql:"losses"`
	Rank         string `cql:"rank"`
	Tier         string `cql:"tier"`
	Wins         int    `cql:"wins"`
	Current      int    `cql:"current"`
}

func DeserializeLeague(league *DBLeague) *ScyllaLeagueEntry {
	return &ScyllaLeagueEntry{
		SummonerId:   league.SummonerId,
		QueueType:    league.QueueType,
		LeaguePoints: int32(league.LeaguePoints),
		Losses:       int32(league.Losses),
		Rank:         league.Rank,
		Tier:         league.Tier,
		Wins:         int32(league.Wins),
		Current:      int32(league.Current),
	}
}

func CreateLeagueTable(region string) {
	query := fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %s_lol.league (
						summonerId text,
						queueType text,
						leaguePoints int,
						losses int,
						rank text,
						tier text,
						wins int,
						current int,
						PRIMARY KEY (summonerId, queueType)
				)`, region)
	err := Session.Query(query).Exec()
	if err != nil {
		log.Fatalf("error creating solo league table: %v", Error(err))
	}
}

func InitLeague(region string) {
	CreateLeagueTable(region)
}
