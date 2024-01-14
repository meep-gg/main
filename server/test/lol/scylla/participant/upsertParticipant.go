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

func TestUpsertParticipant(t *testing.T) {
	conn := GRPCConnect(LolScyllaAddr)
	defer conn.Close()
	client := NewScyllaParticipantServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Log(Info("It should upsert participant"))
	_, err := client.UpsertParticipant(ctx, &ScyllaUpsertParticipantReq{
		Region:      "na1",
		Participant: TestParticipant,
	})
	if err != nil {
		st := status.Convert(err)
		t.Error(Error(st.Message()))
		return
	}
	t.Log(Success("Upserted participant: ") + Success(TestParticipant.GameId))
}
