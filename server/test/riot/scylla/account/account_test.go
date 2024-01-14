package account

import (
	"testing"

	. "meep.gg/log"
)

func Test_RiotScylla_Account(t *testing.T) {
	t.Log(Title("Riot Scylla Account"))
	t.Run("UpsertAccount", TestUpsertAccount)
	t.Run("DeleteAccount", TestDeleteAccount)
	t.Run("GetAccount", TestGetAccount)
}
