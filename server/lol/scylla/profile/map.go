package profile

import (
	"fmt"
	"log"

	. "meep.gg/log"
	. "meep.gg/protos/scylla/lol/profile/v1"
	. "meep.gg/scylla"
)

type DBProfile struct {
	GameName      string `cql:"gamename"`
	TagLine       string `cql:"tagline"`
	ProfileIconId int32  `cql:"profileiconid"`
	Puuid         string `cql:"puuid"`
	SummonerId    string `cql:"summonerid"`
	SummonerLevel int32  `cql:"summonerlevel"`
}

func DeserializeProfile(profile *DBProfile) *ScyllaProfile {
	return &ScyllaProfile{
		GameName:      profile.GameName,
		TagLine:       profile.TagLine,
		ProfileIconId: profile.ProfileIconId,
		Puuid:         profile.Puuid,
		SummonerId:    profile.SummonerId,
		SummonerLevel: profile.SummonerLevel,
	}
}

func CreateProfileTable(region string) {
	// Create the profile table by puuid
	query := fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %s_lol.profile (
						gameName text,
						tagLine text,
						profileIconId int,
						puuid text,
						summonerId text,
						summonerLevel int,
						PRIMARY KEY (gameName, tagLine)
				)`, region)
	err := Session.Query(query).Exec()
	if err != nil {
		log.Fatalf("error creating profile table: %v", Error(err))
	}
}

func InitProfile(region string) {
	CreateProfileTable(region)
}
