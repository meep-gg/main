package participant

import (
	"testing"

	. "meep.gg/log"
)

func Test_LolScylla_Participant(t *testing.T) {
	t.Log(Title("Lol Scylla Participant"))
	t.Run("DeleteParticipant", TestDeleteParticipant)
	t.Run("GetParticipant", TestGetParticipant)
	t.Run("GetParticipants", TestGetParticipants)
	t.Run("UpsertParticipant", TestUpsertParticipant)
	t.Run("PaginateParticipant", TestPaginateParticipant)
	t.Run("GetMatchParticipant", TestGetMatchParticipant)
}
