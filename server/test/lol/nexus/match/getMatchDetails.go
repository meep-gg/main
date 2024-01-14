package match

import (
	"context"
	"testing"
	"time"

	. "meep.gg/grpc"
	. "meep.gg/log"
	. "meep.gg/protos/nexus/lol/match/v1"
	. "meep.gg/test-server/lol/nexus"
)

func TestGetMatchDetails(t *testing.T) {
	conn := GRPCConnect(LolNexusAddr)
	defer conn.Close()
	client := NewMatchServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Log(Info("It should get a detailed match"))
	data, err := client.GetMatchDetails(ctx, &GetMatchReq{
		Region: "na1",
		GameId: 4889768422,
	})
	if err != nil {
		t.Error(Error(err.Error()))
		return
	}
	t.Log(Success("Got participants: ") + Success(len(data.Details)))
	for _, p := range data.Details {
		t.Log(Success("Got participantId: ") + Success(p.ParticipantId))
	}
}
