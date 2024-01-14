package league

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	. "meep.gg/grpc"
	. "meep.gg/log"
	. "meep.gg/protos/gateway/lol/league/v1"
	. "meep.gg/test-server/lol/gateway"
)

func TestUpdateLeagues(t *testing.T) {
	conn := GRPCConnect(LolGatewayAddr)
	defer conn.Close()
	client := NewGatewayLeagueServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Log(Info("It should update players leagues"))
	_, err := client.UpdateLeagues(ctx, &GatewayLeagueReq{
		Region:     "na1",
		SummonerId: SummonerId,
	})

	if err != nil {
		st := status.Convert(err)
		t.Error(Error(st.Message()))
		return
	}

	t.Log(Success("Updated players leagues"))

	t.Log(Info("It should return not found error if summoner doesn't exist"))
	_, err = client.UpdateLeagues(ctx, &GatewayLeagueReq{
		Region:     "na1",
		SummonerId: "123",
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
