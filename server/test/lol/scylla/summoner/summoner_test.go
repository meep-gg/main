package summoner

import (
	"testing"

	. "meep.gg/log"
)

func Test_LolScylla_Summoner(t *testing.T) {
	t.Log(Title("Lol Scylla Summoner"))
	t.Run("GetSummoner", TestGetSummoner)
	t.Run("UpsertSummoner", TestUpsertSummoner)
	t.Run("DeleteSummoner", TestDeleteSummoner)
}
