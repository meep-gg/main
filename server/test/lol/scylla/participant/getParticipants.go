package participant

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc/status"
	. "meep.gg/grpc"
	. "meep.gg/log"
	. "meep.gg/protos/scylla/lol/participant/v1"
	. "meep.gg/test-server/lol/scylla"
)

func TestGetParticipants(t *testing.T) {
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

	t.Log(Info("It should get participants"))
	data, err := client.GetParticipants(ctx, &ScyllaGetParticipantsReq{
		Region: "na1",
		Puuid:  TestParticipant.Puuid,
	})
	if err != nil {
		st := status.Convert(err)
		t.Error(Error(st.Message()))
		return
	}
	t.Log(Success("Got participants: ") + Success(len(data.Participants)))
}
