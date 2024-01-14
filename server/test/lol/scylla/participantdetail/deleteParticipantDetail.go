package participantdetail

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	. "meep.gg/grpc"
	. "meep.gg/log"
	. "meep.gg/protos/scylla/lol/participantdetail/v1"
	. "meep.gg/test-server/lol/scylla"
)

func TestDeleteParticipantdetail(t *testing.T) {
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

	t.Log("Deleting participant detail...")
	_, err = client.DeleteParticipantdetail(ctx, &ScyllaParticipantdetailReq{
		Region:        "na1",
		GameId:        TestParticipantdetail.GameId,
		ParticipantId: TestParticipantdetail.ParticipantId,
	})
	if err != nil {
		st := status.Convert(err)
		t.Error(Error(st.Message()))
		return
	}
	t.Log(Success("Deleted participant detail: ") + Success(TestParticipantdetail.GameId))

	t.Log(Info("It should return not found error if participant detail doesn't exist"))
	_, err = client.DeleteParticipantdetail(ctx, &ScyllaParticipantdetailReq{
		Region:        "na1",
		GameId:        41232,
		ParticipantId: TestParticipantdetail.ParticipantId,
	})
	if err != nil {
		st := status.Convert(err)
		if st.Code() != codes.NotFound {
			t.Error(Error(st.Message()))
			return
		}
	}
	t.Log(Success("Got not found error"))
}
