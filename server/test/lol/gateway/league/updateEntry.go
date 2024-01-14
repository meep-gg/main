package league

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc/status"
	. "meep.gg/grpc"
	. "meep.gg/log"
	. "meep.gg/protos/gateway/lol/league/v1"
	. "meep.gg/test-server/lol/gateway"
)

func TestUpdateEntry(t *testing.T) {
	conn := GRPCConnect(LolGatewayAddr)
	defer conn.Close()
	client := NewGatewayLeagueServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Log(Info("It should update players leagues from DIAMOND I soloqueue page 1"))
	data, err := client.UpdateEntry(ctx, &GatewayUpdateEntryReq{
		Region:    "na1",
		QueueType: "RANKED_SOLO_5x5",
		Tier:      "DIAMOND",
		Division:  "I",
		Page:      1,
	})
	if err != nil {
		st := status.Convert(err)
		t.Error(Error(st.Message()))
		return
	}

	t.Log(Success("Updated players leagues from DIAMOND I soloqueue page 1, documents updated: ") + Success(data.Documents))
}
