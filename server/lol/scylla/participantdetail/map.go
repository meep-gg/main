package participantdetail

import (
	"fmt"
	"log"

	. "meep.gg/log"
	. "meep.gg/protos/scylla/lol/participantdetail/v1"
	. "meep.gg/scylla"
)

type DBPerkStyleSelection struct {
	Perk int32 `cql:"perk"`
	Var1 int32 `cql:"var1"`
	Var2 int32 `cql:"var2"`
	Var3 int32 `cql:"var3"`
}

type DBPerkStyle struct {
	Description string                  `cql:"description"`
	Selections  []*DBPerkStyleSelection `cql:"selections"`
	Style       int32                   `cql:"style"`
}

type DBStatPerks struct {
	Defense int32 `cql:"defense"`
	Flex    int32 `cql:"flex"`
	Offense int32 `cql:"offense"`
}

type DBPerks struct {
	StatPerks *DBStatPerks   `cql:"statperks"`
	Styles    []*DBPerkStyle `cql:"styles"`
}

type DBParticipantdetail struct {
	GameId                         int64    `cql:"gameid"`
	ParticipantId                  int32    `cql:"participantid"`
	DamageDealtToBuildings         int32    `cql:"damagedealttobuildings"`
	DamageDealtToObjectives        int32    `cql:"damagedealttoobjectives"`
	DamageDealtToTurrets           int32    `cql:"damagedealttoturrets"`
	DamageSelfMitigated            int32    `cql:"damageselfmitigated"`
	MagicDamageDealt               int32    `cql:"magicdamagedealt"`
	MagicDamageDealtToChampions    int32    `cql:"magicdamagedealttochampions"`
	MagicDamageTaken               int32    `cql:"magicdamagetaken"`
	NeutralMinionsKilled           int32    `cql:"neutralminionskilled"`
	Perks                          *DBPerks `cql:"perks"`
	PhysicalDamageDealt            int32    `cql:"physicaldamagedealt"`
	PhysicalDamageDealtToChampions int32    `cql:"physicaldamagedealttochampions"`
	PhysicalDamageTaken            int32    `cql:"physicaldamagetaken"`
	Spell1Casts                    int32    `cql:"spell1casts"`
	Spell2Casts                    int32    `cql:"spell2casts"`
	Spell3Casts                    int32    `cql:"spell3casts"`
	Spell4Casts                    int32    `cql:"spell4casts"`
	Summoner1Casts                 int32    `cql:"summoner1casts"`
	Summoner2Casts                 int32    `cql:"summoner2casts"`
	DamageTakenOnTeamPercentage    float32  `cql:"damagetakenonteampercentage"`
	MaxCsAdvantageOnLaneOpponent   float32  `cql:"maxcsadvantageonlaneopponent"`
	MaxLevelLeadLaneOpponent       float32  `cql:"maxlevelleadlaneopponent"`
}

func SerializePerks(perks *ScyllaPerks) *DBPerks {
	var styles []*DBPerkStyle
	if perks == nil {
		return nil
	}

	if perks.StatPerks == nil {
		perks.StatPerks = &ScyllaPerkStats{}
	}
	for _, style := range perks.Styles {
		var selections []*DBPerkStyleSelection
		for _, selection := range style.Selections {
			selections = append(selections, &DBPerkStyleSelection{
				Perk: selection.Perk,
				Var1: selection.Var1,
				Var2: selection.Var2,
				Var3: selection.Var3,
			})
		}
		styles = append(styles, &DBPerkStyle{
			Description: style.Description,
			Selections:  selections,
			Style:       style.Style,
		})
	}
	return &DBPerks{
		StatPerks: &DBStatPerks{
			Defense: perks.StatPerks.Defense,
			Flex:    perks.StatPerks.Flex,
			Offense: perks.StatPerks.Offense,
		},
		Styles: styles,
	}
}

func DeserializePerks(perks *DBPerks) *ScyllaPerks {
	var styles []*ScyllaPerkStyle
	if perks == nil {
		return nil
	}

	if perks.StatPerks == nil {
		perks.StatPerks = &DBStatPerks{}
	}
	for _, style := range perks.Styles {
		var selections []*ScyllaPerkStyleSelection
		for _, selection := range style.Selections {
			selections = append(selections, &ScyllaPerkStyleSelection{
				Perk: selection.Perk,
				Var1: selection.Var1,
				Var2: selection.Var2,
				Var3: selection.Var3,
			})
		}
		styles = append(styles, &ScyllaPerkStyle{
			Description: style.Description,
			Selections:  selections,
			Style:       style.Style,
		})
	}
	return &ScyllaPerks{
		StatPerks: &ScyllaPerkStats{
			Defense: perks.StatPerks.Defense,
			Flex:    perks.StatPerks.Flex,
			Offense: perks.StatPerks.Offense,
		},
		Styles: styles,
	}
}

func DeserializeParticipantdetail(participantdetail *DBParticipantdetail) *ScyllaParticipantdetail {
	return &ScyllaParticipantdetail{
		GameId:                         participantdetail.GameId,
		ParticipantId:                  participantdetail.ParticipantId,
		DamageDealtToBuildings:         participantdetail.DamageDealtToBuildings,
		DamageDealtToObjectives:        participantdetail.DamageDealtToObjectives,
		DamageDealtToTurrets:           participantdetail.DamageDealtToTurrets,
		DamageSelfMitigated:            participantdetail.DamageSelfMitigated,
		MagicDamageDealt:               participantdetail.MagicDamageDealt,
		MagicDamageDealtToChampions:    participantdetail.MagicDamageDealtToChampions,
		MagicDamageTaken:               participantdetail.MagicDamageTaken,
		NeutralMinionsKilled:           participantdetail.NeutralMinionsKilled,
		Perks:                          DeserializePerks(participantdetail.Perks),
		PhysicalDamageDealt:            participantdetail.PhysicalDamageDealt,
		PhysicalDamageDealtToChampions: participantdetail.PhysicalDamageDealtToChampions,
		PhysicalDamageTaken:            participantdetail.PhysicalDamageTaken,
		Spell1Casts:                    participantdetail.Spell1Casts,
		Spell2Casts:                    participantdetail.Spell2Casts,
		Spell3Casts:                    participantdetail.Spell3Casts,
		Spell4Casts:                    participantdetail.Spell4Casts,
		Summoner1Casts:                 participantdetail.Summoner1Casts,
		Summoner2Casts:                 participantdetail.Summoner2Casts,
		DamageTakenOnTeamPercentage:    participantdetail.DamageTakenOnTeamPercentage,
		MaxCsAdvantageOnLaneOpponent:   participantdetail.MaxCsAdvantageOnLaneOpponent,
		MaxLevelLeadLaneOpponent:       participantdetail.MaxLevelLeadLaneOpponent,
	}
}

func SerializeParticipantdetail(participantdetail *ScyllaParticipantdetail) *DBParticipantdetail {
	return &DBParticipantdetail{
		GameId:                         participantdetail.GameId,
		ParticipantId:                  participantdetail.ParticipantId,
		DamageDealtToBuildings:         participantdetail.DamageDealtToBuildings,
		DamageDealtToObjectives:        participantdetail.DamageDealtToObjectives,
		DamageDealtToTurrets:           participantdetail.DamageDealtToTurrets,
		DamageSelfMitigated:            participantdetail.DamageSelfMitigated,
		MagicDamageDealt:               participantdetail.MagicDamageDealt,
		MagicDamageDealtToChampions:    participantdetail.MagicDamageDealtToChampions,
		MagicDamageTaken:               participantdetail.MagicDamageTaken,
		NeutralMinionsKilled:           participantdetail.NeutralMinionsKilled,
		Perks:                          SerializePerks(participantdetail.Perks),
		PhysicalDamageDealt:            participantdetail.PhysicalDamageDealt,
		PhysicalDamageDealtToChampions: participantdetail.PhysicalDamageDealtToChampions,
		PhysicalDamageTaken:            participantdetail.PhysicalDamageTaken,
		Spell1Casts:                    participantdetail.Spell1Casts,
		Spell2Casts:                    participantdetail.Spell2Casts,
		Spell3Casts:                    participantdetail.Spell3Casts,
		Spell4Casts:                    participantdetail.Spell4Casts,
		Summoner1Casts:                 participantdetail.Summoner1Casts,
		Summoner2Casts:                 participantdetail.Summoner2Casts,
		DamageTakenOnTeamPercentage:    participantdetail.DamageTakenOnTeamPercentage,
		MaxCsAdvantageOnLaneOpponent:   participantdetail.MaxCsAdvantageOnLaneOpponent,
		MaxLevelLeadLaneOpponent:       participantdetail.MaxLevelLeadLaneOpponent,
	}
}

func CreatePDTypes(region string) {
	query := fmt.Sprintf(
		`CREATE TYPE IF NOT EXISTS %v_lol.perkStyle_selection (
						perk int,
						var1 int,
						var2 int,
						var3 int
		)`, region)
	err := Session.Query(query).Exec()
	if err != nil {
		log.Fatalf("error creating perkStyleSelection type: %v", Error(err))
	}

	query = fmt.Sprintf(
		`CREATE TYPE IF NOT EXISTS %v_lol.perk_style (
						description text,
						selections frozen<list<perkStyle_selection>>,
						style int
		)`, region)
	err = Session.Query(query).Exec()
	if err != nil {
		log.Fatalf("error creating perkStyle type: %v", Error(err))
	}

	query = fmt.Sprintf(
		`CREATE TYPE IF NOT EXISTS %v_lol.stat_perks (
						defense int,
						flex int,
						offense int
		)`, region)
	err = Session.Query(query).Exec()
	if err != nil {
		log.Printf("error creating statPerks type: %v", Error(err))
	}

	query = fmt.Sprintf(
		`CREATE TYPE IF NOT EXISTS %v_lol.perks (
						statPerks frozen<stat_perks>,
						styles frozen<list<perk_style>>
		)`, region)
	err = Session.Query(query).Exec()
	if err != nil {
		log.Printf("error creating perks type: %v", Error(err))
	}
}

func CreateParticipantdetailTable(region string) {
	query := fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %v_lol.participant_detail (
						gameId bigint,
						participantId int,
						damageDealtToBuildings int,
						damageDealtToObjectives int,
						damageDealtToTurrets int,
						damageSelfMitigated int,
						damageTakenOnTeamPercentage float,
						magicDamageDealt int,
						magicDamageDealtToChampions int,
						magicDamageTaken int,
						maxCsAdvantageOnLaneOpponent float,
						maxLevelLeadLaneOpponent float,
						neutralMinionsKilled int,
						perks frozen<perks>,
						physicalDamageDealt int,
						physicalDamageDealtToChampions int,
						physicalDamageTaken int,
						spell1Casts int,
						spell2Casts int,
						spell3Casts int,
						spell4Casts int,
						summoner1Casts int,
						summoner2Casts int,
						PRIMARY KEY (gameId, participantId)
		)`, region)

	err := Session.Query(query).Exec()
	if err != nil {
		log.Fatalf("error creating participant_detail table: %v", Error(err))
	}
}

func InitParticipantdetail(region string) {
	CreatePDTypes(region)
	CreateParticipantdetailTable(region)
}
