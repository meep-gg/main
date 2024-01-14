package participant

import (
	"fmt"
	"log"

	. "meep.gg/log"
	. "meep.gg/protos/scylla/lol/participant/v1"
	. "meep.gg/scylla"
)

type DBParticipant struct {
	Puuid                       string  `cql:"puuid"`
	GameId                      int64   `cql:"gameid"`
	ChampionId                  int32   `cql:"championid"`
	ChampLevel                  int32   `cql:"champlevel"`
	Kda                         float32 `cql:"kda"`
	Kills                       int32   `cql:"kills"`
	NeutralMinionsKilled        int32   `cql:"neutralminionskilled"`
	Deaths                      int32   `cql:"deaths"`
	Assists                     int32   `cql:"assists"`
	GoldEarned                  int32   `cql:"goldearned"`
	GoldPerMinute               float32 `cql:"goldperminute"`
	TotalMinionsKilled          int32   `cql:"totalminionskilled"`
	TotalDamageDealtToChampions int32   `cql:"totaldamagedealttochampions"`
	DamagePerMinute             float32 `cql:"damageperminute"`
	TotalDamageTaken            int32   `cql:"totaldamagetaken"`
	TimeCCingOthers             int32   `cql:"timeccingothers"`
	IndividualPosition          string  `cql:"individualposition"`
	Win                         bool    `cql:"win"`
	Rank                        int32   `cql:"rank"`
	ParticipantId               int32   `cql:"participantid"`
	Item0                       int32   `cql:"item0"`
	Item1                       int32   `cql:"item1"`
	Item2                       int32   `cql:"item2"`
	Item3                       int32   `cql:"item3"`
	Item4                       int32   `cql:"item4"`
	Item5                       int32   `cql:"item5"`
	Item6                       int32   `cql:"item6"`
	Summoner1Id                 int32   `cql:"summoner1id"`
	Summoner2Id                 int32   `cql:"summoner2id"`
	PrimaryStyle                int32   `cql:"primarystyle"`
	SubStyle                    int32   `cql:"substyle"`
	TeamId                      int32   `cql:"teamid"`
	Placement                   int32   `cql:"placement"`
	PlayerAugment1              int32   `cql:"playeraugment1"`
	PlayerAugment2              int32   `cql:"playeraugment2"`
	PlayerAugment3              int32   `cql:"playeraugment3"`
	PlayerAugment4              int32   `cql:"playeraugment4"`
	PlayerSubteamId             int32   `cql:"playersubteamid"`
	TeamDamagePercentage        float32 `cql:"teamdamagepercentage"`
	VisionScorePerMinute        float32 `cql:"visionscoreperminute"`
	KillParticipation           float32 `cql:"killparticipation"`
}

func DeserializeParticipant(participant *DBParticipant) *ScyllaParticipant {
	return &ScyllaParticipant{
		Puuid:                       participant.Puuid,
		GameId:                      participant.GameId,
		ChampionId:                  participant.ChampionId,
		ChampLevel:                  participant.ChampLevel,
		Kda:                         participant.Kda,
		Kills:                       participant.Kills,
		Deaths:                      participant.Deaths,
		Assists:                     participant.Assists,
		GoldEarned:                  participant.GoldEarned,
		GoldPerMinute:               participant.GoldPerMinute,
		TotalMinionsKilled:          participant.TotalMinionsKilled,
		TotalDamageDealtToChampions: participant.TotalDamageDealtToChampions,
		DamagePerMinute:             participant.DamagePerMinute,
		TotalDamageTaken:            participant.TotalDamageTaken,
		TimeCCingOthers:             participant.TimeCCingOthers,
		IndividualPosition:          participant.IndividualPosition,
		Win:                         participant.Win,
		Rank:                        participant.Rank,
		ParticipantId:               participant.ParticipantId,
		Item0:                       participant.Item0,
		Item1:                       participant.Item1,
		Item2:                       participant.Item2,
		Item3:                       participant.Item3,
		Item4:                       participant.Item4,
		Item5:                       participant.Item5,
		Item6:                       participant.Item6,
		Summoner1Id:                 participant.Summoner1Id,
		Summoner2Id:                 participant.Summoner2Id,
		PrimaryStyle:                participant.PrimaryStyle,
		SubStyle:                    participant.SubStyle,
		TeamId:                      participant.TeamId,
		Placement:                   participant.Placement,
		PlayerAugment1:              participant.PlayerAugment1,
		PlayerAugment2:              participant.PlayerAugment2,
		PlayerAugment3:              participant.PlayerAugment3,
		PlayerAugment4:              participant.PlayerAugment4,
		PlayerSubteamId:             participant.PlayerSubteamId,
		TeamDamagePercentage:        participant.TeamDamagePercentage,
		VisionScorePerMinute:        participant.VisionScorePerMinute,
		KillParticipation:           participant.KillParticipation,
		NeutralMinionsKilled:        participant.NeutralMinionsKilled,
	}
}

func CreateParticipantTable(region string) {
	query := fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %v_lol.participant (
						puuid text,
						gameId bigint,	
						assists int,
						championId int,
						champLevel int,
						damagePerMinute float,
						deaths int,
						goldEarned int,
						goldPerMinute float,
						individualPosition text,
						item0 int,
						item1 int,
						item2 int,
						item3 int,
						item4 int,
						item5 int,
						item6 int,
						kda float,
						killParticipation float,
						kills int,
						neutralMinionsKilled int,
						participantId int,
						placement int,
						playerAugment1 int,
						playerAugment2 int,
						playerAugment3 int,
						playerAugment4 int,
						playerSubteamId int,
						primaryStyle int,
						rank int,
						subStyle int,
						summoner1Id int,
						summoner2Id int,
						teamDamagePercentage float,
						teamId int,
						timeCCingOthers int,
						totalDamageDealtToChampions int,
						totalDamageTaken int,
						totalMinionsKilled int,
						visionScorePerMinute float,
						win boolean,
						PRIMARY KEY (puuid, gameId)
				) WITH CLUSTERING ORDER BY (gameId DESC)`, region)
	err := Session.Query(query).Exec()
	if err != nil {
		log.Fatalf("error creating participant table: %v", Error(err))
	}
}

func InitParticipant(region string) {
	CreateParticipantTable(region)
}
