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

func TestUpsertParticipantdetail(t *testing.T) {
	conn := GRPCConnect(LolScyllaAddr)
	defer conn.Close()
	client := NewScyllaParticipantdetailServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Log(Info("It should upsert a participant detail"))
	_, err := client.UpsertParticipantdetail(ctx, &ScyllaUpsertParticipantdetailReq{
		Region: "na1",
		Detail: TestParticipantdetail,
	})

	if err != nil {
		st := status.Convert(err)
		t.Error(Error(st.Message()))
		return
	}
	t.Log(Success("Upserted participant detail: ") + Success(TestParticipantdetail.GameId))

}
