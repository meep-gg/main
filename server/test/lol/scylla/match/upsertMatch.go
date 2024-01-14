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

func TestUpsertMatch(t *testing.T) {
	conn := GRPCConnect(LolScyllaAddr)
	defer conn.Close()
	client := NewScyllaMatchServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Log(Info("It should upsert a match"))
	_, err := client.UpsertMatch(ctx, &ScyllaUpsertMatchReq{
		Region: "na1",
		Match:  TestMatch,
	})
	if err != nil {
		st := status.Convert(err)
		t.Error(Error(st.Message()))
		return
	}
	t.Log(Success("Upserted match: ") + Success(TestMatch.GameId))
}
