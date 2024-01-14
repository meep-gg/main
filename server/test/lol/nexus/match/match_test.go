package match

import (
	"testing"

	. "meep.gg/log"
)

func Test_LolNexus_Match(t *testing.T) {
	t.Log(Title("Lol Nexus Match"))
	t.Run("GetMatches", TestGetMatches)
	t.Run("GetMatchParticipants", TestGetMatchParticipants)
	t.Run("UpdatePlayerMatches", TestUpdatePlayerMatches)
	t.Run("GetMatchDetails", TestGetMatchDetails)

}
