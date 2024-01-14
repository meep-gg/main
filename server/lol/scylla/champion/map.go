package champion

import (
	"fmt"
	"log"

	. "meep.gg/log"
	. "meep.gg/scylla"
)

func CreateChampionTable(region string) {
	// Create the champion table by puuid
	query := fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %s_lol.champion (
						puuid text,
						championId int,
						season int,
						patch int,
						queueId int,
						winRate int,
						wins int,
						losses int,
						kda float,
						kills int,
						deaths int,
						assists int,
						maxKills int,
						maxDeaths int,
						minionCs int,
						jungleCs int,
						damageDone int,
						damageTaken int,
						gold int,
						PRIMARY KEY (puuid, championId, season, patch, queueId)
				)`, region)
	err := Session.Query(query).Exec()
	if err != nil {
		log.Fatalf("error creating champion table: %v", Error(err))
	}
}

func InitChampion(region string) {
	CreateChampionTable(region)
}
