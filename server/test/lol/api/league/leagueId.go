package league

import (
	"context"
	"testing"
	"time"

	. "meep.gg/grpc"
	. "meep.gg/log"
	. "meep.gg/protos/riot/lol/league/v1"
	. "meep.gg/test-server/lol/api"
)

func TestLeagueId(t *testing.T) {
	conn := GRPCConnect(LolAPIAddr)
	defer conn.Close()
	client := NewRiotLeagueServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Log(Info("It should return the list of league entries by league id"))
	league, err := client.LeagueId(ctx, &RiotLeagueIdReq{
		Region:   "na1",
		LeagueId: "a68eddc7-80ee-4032-b5bc-f80ac8d8a21f",
	})

	if err != nil {
		t.Error(Error(err))
	}
	t.Log(Success("Retrevied list of league entries by league id: ") + Success(len(league.Entries)))
}
