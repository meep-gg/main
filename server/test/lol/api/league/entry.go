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

func TestEntry(t *testing.T) {
	conn := GRPCConnect(LolAPIAddr)
	defer conn.Close()
	client := NewRiotLeagueServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Log(Info("It should get get  list of Entries without inputing a page"))
	data, err := client.Entry(ctx, &RiotLeagueEntryReq{
		Region:   "na1",
		Queue:    "RANKED_SOLO_5x5",
		Tier:     "DIAMOND",
		Division: "I",
	})
	if err != nil {
		st, _ := status.FromError(err)
		t.Error(Error(st.Message()))
		return
	}

	if data == nil {
		t.Error(Error("Data is nil"))
		return
	}
	t.Log(Success("Retrieved list of Entries : ") + Success(len(data.Leagues)))

	t.Log(Info("It should get get  list of Entries with inputing a page"))
	data, err = client.Entry(ctx, &RiotLeagueEntryReq{
		Region:   "na1",
		Queue:    "RANKED_SOLO_5x5",
		Tier:     "DIAMOND",
		Division: "I",
		Page:     2,
	})

	if err != nil {
		st, _ := status.FromError(err)
		t.Error(Error(st.Message()))
		return
	}

	if data == nil {
		t.Error(Error("Data is nil"))
		return
	}

	t.Log(Success("Retrieved list of Entries : ") + Success(len(data.Leagues)))
}
