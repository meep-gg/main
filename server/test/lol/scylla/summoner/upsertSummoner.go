package summoner

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc/status"
	. "meep.gg/grpc"
	. "meep.gg/log"
	. "meep.gg/protos/scylla/lol/summoner/v1"
	. "meep.gg/test-server/lol/scylla"
)

func TestUpsertSummoner(t *testing.T) {
	conn := GRPCConnect(LolScyllaAddr)
	defer conn.Close()
	client := NewScyllaSummonerServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Log(Info("It should upsert summoner"))
	_, err := client.UpsertSummoner(ctx, &ScyllaUpsertSummonerReq{
		Region:   "na1",
		Summoner: TestSummoner,
	})
	if err != nil {
		st := status.Convert(err)
		t.Error(Error(st.Message()))
		return
	}
	t.Log(Success("Upserted summoner: ") + Success(TestSummoner.Name))
}
