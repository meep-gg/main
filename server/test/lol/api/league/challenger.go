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

func TestChallenger(t *testing.T) {
	conn := GRPCConnect(LolAPIAddr)
	defer conn.Close()
	client := NewRiotLeagueServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Log(Info("It should return the list of solo queue challenger entries"))
	challenger, err := client.Challenger(ctx, &RiotLeagueApexReq{
		Region: "na1",
		Queue:  "RANKED_SOLO_5x5",
	})
	if err != nil {
		t.Error(Error(err))
	}
	t.Log(Success("Retrevied list of challenger entries: ") + Success(len(challenger.Entries)))

	t.Log(Info("It should return the list of flex queue challenger entries"))
	challenger, err = client.Challenger(ctx, &RiotLeagueApexReq{
		Region: "na1",
		Queue:  "RANKED_FLEX_SR",
	})
	if err != nil {
		t.Error(Error(err))
	}
	t.Log(Success("Retrevied list of challenger entries: ") + Success(len(challenger.Entries)))
}
