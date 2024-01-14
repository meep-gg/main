package league

import (
	"testing"

	. "meep.gg/log"
)

func Test_LolGateway_League(t *testing.T) {
	t.Log(Title("Lol Gateway League"))
	t.Run("GetLeagues", TestGetLeagues)
	t.Run("UpdateLeague", TestUpdateLeagues)
	t.Run("UpdateChallenger", TestUpdateChallenger)
	t.Run("UpdateGrandmaster", TestUpdateGrandmaster)
	t.Run("UpdateMaster", TestUpdateMaster)
	t.Run("UpdateEntry", TestUpdateEntry)
	// t.Run("UpdateAll", TestUpdateAll)
}
