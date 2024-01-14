package participantdetail

import (
	"testing"

	. "meep.gg/log"
)

func Test_LolScylla_Participantdetail(t *testing.T) {
	t.Log(Title("Lol Scylla Participantdetail"))
	t.Run("GetParticipantdetail", TestGetParticipantdetail)
	t.Run("GetMatchParticipantsdetail", TestGetMatchParticipantsdetail)
	t.Run("UpsertParticipantdetail", TestUpsertParticipantdetail)
	t.Run("DeleteParticipantdetail", TestDeleteParticipantdetail)
}
