package summoner

import (
	"testing"

	. "meep.gg/log"
)

func Test_LolGateway_Summoner(t *testing.T) {
	t.Log(Title("Lol Gateway Summoner"))
	t.Run("GetSummoner", TestGetSummoner)
	t.Run("UpdateSummoner", TestUpdateSummoner)
}
