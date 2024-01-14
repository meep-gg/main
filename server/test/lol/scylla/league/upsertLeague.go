package league

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc/status"
	. "meep.gg/grpc"
	. "meep.gg/log"
	. "meep.gg/protos/scylla/lol/league/v1"
	. "meep.gg/test-server/lol/scylla"
)

func TestUpsertLeague(t *testing.T) {
	conn := GRPCConnect(LolScyllaAddr)
	defer conn.Close()
	client := NewScyllaLeagueServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Log(Info("It should upsert solo league"))
	_, err := client.UpsertLeague(ctx, &ScyllaUpsertLeagueReq{
		Region: "na1",
		League: TestLeague,
	})
	if err != nil {
		st := status.Convert(err)
		t.Error(Error(st.Message()))
		return
	}
	t.Log(Success("Upserted league: ") + Success(TestLeague.SummonerId))

	t.Log(Info("It should upsert flex league"))
	_, err = client.UpsertLeague(ctx, &ScyllaUpsertLeagueReq{
		Region: "na1",
		League: TestLeague,
	})

	if err != nil {
		st := status.Convert(err)
		t.Error(Error(st.Message()))
		return
	}
	t.Log(Success("Upserted league: ") + Success(TestLeague.SummonerId))
}
