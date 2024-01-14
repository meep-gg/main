package profile

import (
	"testing"

	. "meep.gg/log"
)

func Test_LolScylla_Profile(t *testing.T) {
	t.Log(Title("Lol Scylla Profile"))
	t.Run("GetProfile", TestGetProfile)
	t.Run("UpsertProfile", TestUpsertProfile)
	t.Run("DeleteProfile", TestDeleteProfile)
}
