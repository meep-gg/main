package participant

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	. "meep.gg/grpc"
	. "meep.gg/log"
	. "meep.gg/protos/scylla/lol/participant/v1"
	. "meep.gg/test-server/lol/scylla"
)

func TestGetParticipant(t *testing.T) {
	conn := GRPCConnect(LolScyllaAddr)
	defer conn.Close()
	client := NewScyllaParticipantServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Log("Upserting participant...")
	_, err := client.UpsertParticipant(ctx, &ScyllaUpsertParticipantReq{
		Region:      "na1",
		Participant: TestParticipant,
	})
	if err != nil {
		st := status.Convert(err)
		t.Error(Error(st.Message()))
		return
	}

	t.Log(Info("It should get a participant"))
	data, err := client.GetParticipant(ctx, &ScyllaParticipantReq{
		Region: "na1",
		GameId: TestParticipant.GameId,
		Puuid:  TestParticipant.Puuid,
	})
	if err != nil {
		st := status.Convert(err)
		t.Error(Error(st.Message()))
		return
	}
	t.Log(Success("Got participant: ") + Success(data.Puuid))

	t.Log(Info("It should return not found error if participant doesn't exist"))
	_, err = client.GetParticipant(ctx, &ScyllaParticipantReq{
		Region: "na1",
		GameId: TestParticipant.GameId,
		Puuid:  "123",
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
