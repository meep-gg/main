package match

import (
	"testing"

	. "meep.gg/log"
)

func Test_LolScylla_Match(t *testing.T) {
	t.Log(Title("Lol Scylla Match"))
	t.Run("UpdatePlayerMatches", TestUpdatePlayerMatches)
}
