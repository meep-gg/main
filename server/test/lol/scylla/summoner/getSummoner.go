package summoner

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	. "meep.gg/grpc"
	. "meep.gg/log"
	. "meep.gg/protos/scylla/lol/summoner/v1"
	. "meep.gg/test-server/lol/scylla"
)

func TestGetSummoner(t *testing.T) {
	conn := GRPCConnect(LolScyllaAddr)
	defer conn.Close()
	client := NewScyllaSummonerServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Log("Upserting summoner...")
	_, err := client.UpsertSummoner(ctx, &ScyllaUpsertSummonerReq{
		Region:   "na1",
		Summoner: TestSummoner,
	})
	if err != nil {
		st := status.Convert(err)
		t.Error(Error(st.Message()))
		return
	}

	t.Log(Info("It should return summoner by Puuid"))
	summoner, err := client.GetSummoner(ctx, &ScyllaSummonerReq{
		Region: "na1",
		Puuid:  TestSummoner.Puuid,
	})

	if err != nil {
		st := status.Convert(err)
		t.Error(Error(st.Message()))
		return
	}
	t.Log(Success("Retrieved summoner by Puuid: ") + Success(summoner.Name))

	t.Log(Info("It should return not found if the summoner does not exist"))
	summoner, err = client.GetSummoner(ctx, &ScyllaSummonerReq{
		Region: "na1",
		Puuid:  "123456789",
	})
	if err != nil {
		st := status.Convert(err)
		if st.Code() != codes.NotFound {
			t.Error(Error(err))
		}
	}
	t.Log(Success("Summoner does not exist: ") + Success(err))
}
