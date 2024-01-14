package match

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc/status"
	. "meep.gg/grpc"
	. "meep.gg/log"
	. "meep.gg/protos/scylla/lol/match/v1"
	. "meep.gg/test-server/lol/scylla"
)

func TestDeleteMatch(t *testing.T) {
	conn := GRPCConnect(LolScyllaAddr)
	defer conn.Close()
	client := NewScyllaMatchServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Log("Upserting match...")
	_, err := client.UpsertMatch(ctx, &ScyllaUpsertMatchReq{
		Region: "na1",
		Match:  TestMatch,
	})
	if err != nil {
		st := status.Convert(err)
		t.Error(Error(st.Message()))
		return
	}

	t.Log(Info("It should delete a match"))
	_, err = client.DeleteMatch(ctx, &ScyllaMatchReq{
		Region: "na1",
		GameId: TestMatch.GameId,
	})
	if err != nil {
		st := status.Convert(err)
		t.Error(Error(st.Message()))
		return
	}
	t.Log(Success("Deleted match: ") + Success(TestMatch.GameId))
}
