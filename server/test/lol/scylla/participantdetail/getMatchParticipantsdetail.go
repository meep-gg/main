package participantdetail

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc/status"
	. "meep.gg/grpc"
	. "meep.gg/log"
	. "meep.gg/protos/scylla/lol/participantdetail/v1"
	. "meep.gg/test-server/lol/scylla"
)

func TestGetMatchParticipantsdetail(t *testing.T) {
	conn := GRPCConnect(LolScyllaAddr)
	defer conn.Close()
	client := NewScyllaParticipantdetailServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Log("Upserting participant detail...")
	_, err := client.UpsertParticipantdetail(ctx, &ScyllaUpsertParticipantdetailReq{
		Region: "na1",
		Detail: TestParticipantdetail,
	})
	if err != nil {
		st := status.Convert(err)
		t.Error(Error(st.Message()))
		return
	}

	t.Log(Info("It should get a every participant detail from a match from Scylla"))
	data, err := client.GetMatchParticipantsdetail(ctx, &ScyllaGetParticipantdetailsReq{
		Region: "na1",
		GameId: TestParticipantdetail.GameId,
		Count:  10,
	})

	if err != nil {
		st := status.Convert(err)
		t.Error(Error(st.Message()))
		return
	}
	for _, participantdetail := range data.Details {
		t.Log(Success("Got ParticipantId: ") + Success(participantdetail.ParticipantId))
	}

}
