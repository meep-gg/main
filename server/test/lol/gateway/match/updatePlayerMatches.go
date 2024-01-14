package match

import (
	"context"
	"testing"
	"time"

	. "meep.gg/grpc"
	. "meep.gg/log"
	. "meep.gg/protos/gateway/lol/match/v1"
	. "meep.gg/test-server/lol/gateway"
)

func TestUpdatePlayerMatches(t *testing.T) {
	conn := GRPCConnect(LolGatewayAddr)
	defer conn.Close()
	client := NewGatewayMatchServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	t.Log(Info("It should update profie: ") + Info(Puuid))
	_, err := client.UpdatePlayerMatches(ctx, &GatewayMatchPlayerReq{
		Region: "na1",
		Puuid:  Puuid,
	})
	if err != nil {
		t.Error(Error(err))
		return
	}

	for _, puuid := range PuuidList {
		t.Log(Info("It should update profie: ") + Info(puuid))
		_, err = client.UpdatePlayerMatches(ctx, &GatewayMatchPlayerReq{
			Region: "na1",
			Puuid:  puuid,
		})
		if err != nil {
			t.Error(Error(err))
			return
		}
	}
	t.Log(Success("Init player matches"))
}
