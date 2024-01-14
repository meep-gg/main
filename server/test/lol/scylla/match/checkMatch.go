package match

import (
	"context"
	"log"
	"testing"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	. "meep.gg/grpc"
	. "meep.gg/log"
	. "meep.gg/protos/scylla/lol/match/v1"
	. "meep.gg/test-server/lol/scylla"
)

func TestCheckMatch(t *testing.T) {
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

	t.Log(Info("It should check a match"))
	_, err = client.CheckMatch(ctx, &ScyllaMatchReq{
		Region: "na1",
		GameId: TestMatch.GameId,
	})
	if err != nil {
		t.Error(Error(err.Error()))
		return
	}
	t.Log(Success("Checked match: ") + Success(TestMatch.GameId))

	t.Log(Info("It should return not found error if match doesn't exist"))
	_, err = client.CheckMatch(ctx, &ScyllaMatchReq{
		Region: "na1",
		GameId: 123,
	})
	if err != nil {
		st := status.Convert(err)
		log.Printf("%+v", st.Code())
		if st.Code() != codes.NotFound {
			t.Error(Error(st.Message()))
			return
		}
	}
	t.Log(Success("Got not found error"))
}
