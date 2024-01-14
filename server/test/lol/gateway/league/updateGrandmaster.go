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

func TestUpdateGrandmaster(t *testing.T) {
	conn := GRPCConnect(LolGatewayAddr)
	defer conn.Close()
	client := NewGatewayLeagueServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Log(Info("It should update solo queue grandmaster league"))
	_, err := client.UpdateGrandmaster(ctx, &GatewayApexUpdateReq{
		Region:    "na1",
		QueueType: "RANKED_SOLO_5x5",
	})

	if err != nil {
		st := status.Convert(err)
		t.Error(Error(st.Message()))
		return
	}
	t.Log(Success("Updated solo queue grandmaster league"))

	t.Log(Info("It should update flex queue grandmaster league"))
	_, err = client.UpdateGrandmaster(ctx, &GatewayApexUpdateReq{
		Region:    "na1",
		QueueType: "RANKED_FLEX_SR",
	})

	if err != nil {
		st := status.Convert(err)
		t.Error(Error(st.Message()))
		return
	}
	t.Log(Success("Updated flex queue grandmaster league"))
}
