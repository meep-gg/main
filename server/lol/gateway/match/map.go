package match

import (
	. "meep.gg/protos/riot/lol/match/v1"
	. "meep.gg/protos/scylla/lol/participantdetail/v1"
)

func SerialPerks(perks *Perks) *ScyllaPerks {
	var styles []*ScyllaPerkStyle
	if perks == nil {
		return nil
	}

	if perks.StatPerks == nil {
		perks.StatPerks = &PerkStats{}
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
