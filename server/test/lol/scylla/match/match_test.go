package match

import (
	"testing"

	. "meep.gg/log"
)

func Test_LolScylla_Match(t *testing.T) {
	t.Log(Title("Lol Scylla Match"))
	t.Run("CheckMatch", TestCheckMatch)
	t.Run("GetMatch", TestGetMatch)
	t.Run("GetMatches", TestGetMatches)
	t.Run("UpsertMatch", TestUpsertMatch)
	t.Run("DeleteMatch", TestDeleteMatch)
}
