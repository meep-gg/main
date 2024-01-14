package match

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	. "meep.gg/grpc"
	. "meep.gg/log"
	. "meep.gg/protos/scylla/lol/match/v1"
	. "meep.gg/test-server/lol/scylla"
)

func TestGetMatch(t *testing.T) {
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

	t.Log(Info("It should get a match"))
	match, err := client.GetMatch(ctx, &ScyllaMatchReq{
		Region: "na1",
		GameId: TestMatch.GameId,
	})
	if err != nil {
		st := status.Convert(err)
		t.Error(Error(st.Message()))
		return
	}
	t.Log(Success("Got match: ") + Success(match.GameId))

	t.Log(Info("It should return not found error if match doesn't exist"))
	_, err = client.GetMatch(ctx, &ScyllaMatchReq{
		Region: "na1",
		GameId: 123,
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
