package league

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc/status"
	. "meep.gg/grpc"
	. "meep.gg/log"
	. "meep.gg/protos/riot/lol/league/v1"
	. "meep.gg/test-server/lol/api"
)

func TestSummonerId(t *testing.T) {
	conn := GRPCConnect(LolAPIAddr)
	defer conn.Close()
	client := NewRiotLeagueServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Log(Info("It should return the list of league entries by summoner id"))
	leagues, err := client.SummonerId(ctx, &RiotLeagueSummonerIdReq{
		Region:     "na1",
		SummonerId: "C-dE3eK53OefJnIPW58zvr2Qjo2fiKxFG_aA5zip9XjT7MtP",
	})

	if err != nil {
		st, _ := status.FromError(err)
		t.Error(Error(st.Message()))
		return
	}

	if leagues == nil {
		t.Error(Error("Data is nil"))
		return
	}
	t.Log(Success("Retrevied list of league entries by summoner id : ") + Success(len(leagues.Leagues)))

}
