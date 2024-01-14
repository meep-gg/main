package summoner

import (
	"fmt"
	"log"

	. "meep.gg/log"
	. "meep.gg/scylla"
)

func CreateSummonerTable(region string) {
	// Create the summoner table by puuid
	query := fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %s_lol.summoner (
						puuid text PRIMARY KEY,
						id text,
						name text,
						profileIconId int,
						summonerLevel int
				)`, region)
	err := Session.Query(query).Exec()
	if err != nil {
		log.Fatalf("error creating summoner table: %v", Error(err))
	}
}

func InitSummoner(region string) {
	CreateSummonerTable(region)
}
