package summoner

import (
	"testing"

	. "meep.gg/log"
)

func Test_LolAPI_Summoner(t *testing.T) {
	t.Log(Title("Lol API Summoner"))
	t.Run("Account", TestAccount)
	t.Run("Name", TestName)
	t.Run("Puuid", TestPuuid)
	t.Run("SummonerId", TestId)
}
