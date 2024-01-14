package league

import (
	"testing"

	. "meep.gg/log"
)

func Test_LolScylla_League(t *testing.T) {
	t.Log(Title("Lol Scylla League"))
	t.Run("GetLeague", TestGetLeague)
	t.Run("UpsertLeague", TestUpsertLeague)
	t.Run("DeleteLeague", TestDeleteLeague)
}
