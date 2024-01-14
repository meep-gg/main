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

func TestMaster(t *testing.T) {
	conn := GRPCConnect(LolAPIAddr)
	defer conn.Close()
	client := NewRiotLeagueServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Log(Info("It should return the list of solo queue master entries"))
	master, err := client.Master(ctx, &RiotLeagueApexReq{
		Region: "na1",
		Queue:  "RANKED_SOLO_5x5",
	})
	if err != nil {
		t.Error(Error(err))
	}
	t.Log(Success("Retrevied list of master entries: ") + Success(len(master.Entries)))

	t.Log(Info("It should return the list of flex queue master entries"))
	master, err = client.Master(ctx, &RiotLeagueApexReq{
		Region: "na1",
		Queue:  "RANKED_FLEX_SR",
	})
	if err != nil {
		t.Error(Error(err))
	}
	t.Log(Success("Retrevied list of master entries: ") + Success(len(master.Entries)))
}
