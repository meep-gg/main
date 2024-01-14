package summoner

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	. "meep.gg/grpc"
	. "meep.gg/log"
	. "meep.gg/protos/gateway/lol/summoner/v1"
	. "meep.gg/test-server/lol/gateway"
)

func TestUpdateSummoner(t *testing.T) {
	conn := GRPCConnect(LolGatewayAddr)
	defer conn.Close()
	client := NewGatewaySummonerServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Log(Info("It should update a summoner from Riot"))
	summoner, err := client.UpdateSummoner(ctx, &GatewaySummonerReq{
		Region: "na1",
		Puuid:  Puuid,
	})
	if err != nil {
		st := status.Convert(err)
		t.Error(Error(st.Message()))
		return
	}
	t.Log(Success("Updated summoner: ") + Success(summoner.Name))

	t.Log(Info("It should return not found error if summoner doesn't exist"))
	summoner, err = client.UpdateSummoner(ctx, &GatewaySummonerReq{
		Region: "na1",
		Puuid:  "123",
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
