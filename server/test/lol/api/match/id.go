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

func TestId(t *testing.T) {
	conn := GRPCConnect(LolAPIAddr)
	defer conn.Close()
	client := NewRiotMatchServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Log(Info("It should return the match by Match Id"))
	match, err := client.Id(ctx, &RiotMatchIdReq{
		Region:  "americas",
		MatchId: Matches[0],
	})
	if err != nil {
		st, _ := status.FromError(err)
		t.Error(Error(st.Message()))
		return
	}
	t.Log(Success("Retrevied match by match id: ") + Success(match.Metadata.MatchId))

	t.Log(Info("It should return not found if the match does not exist"))
	match, err = client.Id(ctx, &RiotMatchIdReq{
		Region:  "americas",
		MatchId: "123456789",
	})
	if err != nil {
		st, _ := status.FromError(err)
		if st.Code() != codes.NotFound {
			t.Error(Error(err))
		}
	}
	t.Log(Success("Match does not exist: ") + Success(err))
}
