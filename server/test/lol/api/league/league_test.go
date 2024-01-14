package league

import (
	"testing"

	. "meep.gg/log"
)

func Test_LolAPI_League(t *testing.T) {
	t.Log(Title("Lol API League"))
	t.Run("Challenger", TestChallenger)
	t.Run("Grandmaster", TestGrandmaster)
	t.Run("Master", TestMaster)
	t.Run("LeagueId", TestLeagueId)
	t.Run("Entries", TestEntry)
	t.Run("SummonerId", TestSummonerId)
}
