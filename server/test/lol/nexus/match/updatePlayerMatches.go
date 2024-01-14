package match

import (
	"context"
	"testing"
	"time"

	. "meep.gg/grpc"
	. "meep.gg/log"
	. "meep.gg/protos/nexus/lol/match/v1"
	. "meep.gg/test-server/lol/nexus"
)

func TestUpdatePlayerMatches(t *testing.T) {
	conn := GRPCConnect(LolNexusAddr)
	defer conn.Close()
	client := NewMatchServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Log(Info("It should update a player's matches"))
	_, err := client.UpdatePlayerMatches(ctx, &UpdatePlayerMatchesReq{
		Region: "na1",
		Puuid:  TestAccount.Puuid,
	})
	if err != nil {
		t.Error(Error(err.Error()))
		return
	}
	t.Log(Success("Updated player matches: ") + Success(TestAccount.Puuid))
}
