package summoner

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	. "meep.gg/grpc"
	. "meep.gg/log"
	. "meep.gg/protos/riot/lol/summoner/v1"
	. "meep.gg/test-server/lol/api"
)

func TestPuuid(t *testing.T) {
	conn := GRPCConnect(LolAPIAddr)
	defer conn.Close()
	client := NewRiotSummonerServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Log(Info("It should return the summoner by puuid"))
	summoner, err := client.Puuid(ctx, &RiotSummonerPuuidReq{
		Region: "na1",
		Puuid:  Summoners[3].Puuid,
	})

	if err != nil {
		t.Error(Error(err))
	}
	t.Log(Success("Retrevied summoner by puuid: ") + Success(summoner.Name))

	t.Log(Info("It should return not found if the summoner does not exist"))
	summoner, err = client.Puuid(ctx, &RiotSummonerPuuidReq{
		Region: "na1",
		Puuid:  "123456789",
	})
	if err != nil {
		st, _ := status.FromError(err)
		if codes.NotFound != st.Code() {
			t.Error(Error(err))
		}
	}
	t.Log(Success("Summoner does not exist: ") + Success(err))
}
