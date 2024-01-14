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

func TestGetMatches(t *testing.T) {
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

	t.Log(Info("It should get matches"))
	data, err := client.GetMatches(ctx, &ScyllaGetMatchesReq{
		Region: "na1",
		GameIds: []int64{
			TestMatch.GameId,
		},
	})
	if err != nil {
		st := status.Convert(err)
		t.Error(Error(st.Message()))
		return
	}
	t.Log(Success("Got matches: ") + Success(len(data.Matches)))
}
