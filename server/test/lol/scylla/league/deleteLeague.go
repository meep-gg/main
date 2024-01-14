package league

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	. "meep.gg/grpc"
	. "meep.gg/log"
	. "meep.gg/protos/scylla/lol/league/v1"
	. "meep.gg/test-server/lol/scylla"
)

func TestDeleteLeague(t *testing.T) {
	conn := GRPCConnect(LolScyllaAddr)
	defer conn.Close()
	client := NewScyllaLeagueServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Log("Upserting league...")
	_, err := client.UpsertLeague(ctx, &ScyllaUpsertLeagueReq{
		Region: "na1",
		League: TestLeague,
	})
	if err != nil {
		st := status.Convert(err)
		t.Error(Error(st.Message()))
		return
	}

	t.Log(Info("It should delete league"))
	_, err = client.DeleteLeague(ctx, &ScyllaLeagueReq{
		Region:     "na1",
		QueueType:  "RANKED_SOLO_5x5",
		SummonerId: TestLeague.SummonerId,
	})
	if err != nil {
		st := status.Convert(err)
		t.Error(Error(st.Message()))
		return
	}
	t.Log(Success("Deleted league: ") + Success(TestLeague.SummonerId))

	t.Log(Info("It should return not found if the league does not exist"))
	_, err = client.DeleteLeague(ctx, &ScyllaLeagueReq{
		Region:     "na1",
		QueueType:  "RANKED_SOLO_5x5",
		SummonerId: "123456789",
	})
	if err != nil {
		st := status.Convert(err)
		if st.Code() != codes.NotFound {
			t.Error(Error(err))
		}
	}
	t.Log(Success("League does not exist: ") + Success(err))
}
