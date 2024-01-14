package account

import (
	"testing"

	. "meep.gg/log"
)

func Test_RiotAPI_Account(t *testing.T) {
	t.Log(Title("Riot API Account"))
	t.Run("Puuid", TestPuuid)
	t.Run("RiotId", TestRiotId)
	t.Run("Stress", TestStress)
}
