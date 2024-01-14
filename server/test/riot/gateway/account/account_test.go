package account

import (
	"testing"

	. "meep.gg/log"
)

func Test_RiotGateway_Account(t *testing.T) {
	t.Log(Title("Riot Gateway Account"))

	t.Run("GetAccount", TestGetAccount)
	t.Run("UpdateAccount", TestUpdateAccount)
}
