package match

import (
	"fmt"
	"log"

	. "meep.gg/log"
	. "meep.gg/protos/scylla/lol/match/v1"
	. "meep.gg/scylla"
)

type DBMatch struct {
	GameId           int64                `cql:"gameid"`
	Bans             []int32              `cql:"bans"`
	GameDuration     int32                `cql:"gameduration"`
	GameEndTimestamp int64                `cql:"gameendtimestamp"`
	Participants     []DBMatchParticipant `cql:"participants"`
	Patch            int32                `cql:"patch"`
	QueueId          int32                `cql:"queueid"`
	Rank             int32                `cql:"rank"`
	Season           int32                `cql:"season"`
}

type DBMatchParticipant struct {
	Puuid      string `cql:"puuid"`
	GameName   string `cql:"gamename"`
	TagLine    string `cql:"tagline"`
	ChampionId int32  `cql:"championid"`
}

func DeserializeMatch(match *DBMatch) *ScyllaMatch {
	participants := make([]*ScyllaMatchParticipant, len(match.Participants))
	for i, participant := range match.Participants {
		participants[i] = DeserializeMatchParticipant(&participant)
	}

	return &ScyllaMatch{
		GameId:           match.GameId,
		Bans:             match.Bans,
		GameDuration:     match.GameDuration,
		GameEndTimestamp: match.GameEndTimestamp,
		Participants:     participants,
		Patch:            match.Patch,
		QueueId:          match.QueueId,
		Rank:             match.Rank,
		Season:           match.Season,
	}
}

func DeserializeMatchParticipant(participant *DBMatchParticipant) *ScyllaMatchParticipant {
	return &ScyllaMatchParticipant{
		Puuid:      participant.Puuid,
		GameName:   participant.GameName,
		TagLine:    participant.TagLine,
		ChampionId: participant.ChampionId,
	}
}

func SerializeMParticipants(participants []*ScyllaMatchParticipant) []DBMatchParticipant {
	mParticipants := make([]DBMatchParticipant, len(participants))
	for i, participant := range participants {
		mParticipants[i] = DBMatchParticipant{
			Puuid:      participant.Puuid,
			GameName:   participant.GameName,
			TagLine:    participant.TagLine,
			ChampionId: participant.ChampionId,
		}
	}
	return mParticipants
}

func CreateMatchTypes(region string) {
	query := fmt.Sprintf(
		`CREATE TYPE IF NOT EXISTS %v_lol.match_participant (
						puuid text,
						gameName text,
						tagLine text,
						championId int
		)`, region)
	err := Session.Query(query).Exec()
	if err != nil {
		log.Fatalf("error creating match_participant type: %v", Error(err))
	}
}

func CreateMatchTable(region string) {
	query := fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %v_lol.match (
						gameId bigint,
						bans list<int>,
						gameDuration int,
						gameEndTimestamp bigint,
						participants list<frozen<match_participant>>,
						patch int,
						queueId int,
						rank int,
						season int,
						PRIMARY KEY (gameId)
						)`, region)
	err := Session.Query(query).Exec()
	if err != nil {
		log.Printf("error creating match table: %v", Error(err))
	}
}

func InitMatch(region string) {
	CreateMatchTypes(region)
	CreateMatchTable(region)
}
