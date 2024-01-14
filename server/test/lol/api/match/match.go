package match

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	. "meep.gg/grpc"
	. "meep.gg/log"
	. "meep.gg/protos/riot/lol/match/v1"
	. "meep.gg/test-server/lol/api"
)

func TestMatch(t *testing.T) {
	conn := GRPCConnect(LolAPIAddr)
	defer conn.Close()
	client := NewRiotMatchServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Log(Info("It should return the match by puuid"))
	match, err := client.Puuid(ctx, &RiotMatchPuuidReq{
		Region: "americas",
		Puuid:  Summoners[0].Puuid,
	})

	if err != nil {
		st, _ := status.FromError(err)
		t.Error(Error(st.Message()))
		return
	}
	t.Log(Success("Retrevied matches by puuid: ") + Success(len(match.Matches)))

	t.Log(Info("It should return not found if the puuid does not exist"))
	match, err = client.Puuid(ctx, &RiotMatchPuuidReq{
		Region: "americas",
		Puuid:  "123456789",
	})
	if err != nil {
		st, _ := status.FromError(err)
		if st.Code() != codes.NotFound {
			t.Error(Error(err))
		}
	}
	t.Log(Success("Puuid does not exist: ") + Success(err))
}
