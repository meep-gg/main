package profile

import (
	"testing"

	. "meep.gg/log"
)

func Test_LolNexus_Profile(t *testing.T) {
	t.Log(Title("Lol Nexus Profile"))
	// t.Run("GetProfile", TestGetProfile)
	t.Run("UpdateProfile", TestUpdateProfile)
}
