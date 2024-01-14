package match

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc/status"
	. "meep.gg/grpc"
	. "meep.gg/log"
	. "meep.gg/protos/nexus/lol/match/v1"
	. "meep.gg/test-server/lol/nexus"
)

func TestGetMatchParticipants(t *testing.T) {
	conn := GRPCConnect(LolNexusAddr)
	defer conn.Close()
	client := NewMatchServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Log(Info("It should get a participants from a match"))
	data, err := client.GetMatchParticipants(ctx, &GetMatchReq{
		Region: "na1",
		GameId: 4889768422,
	})
	if err != nil {
		st := status.Convert(err)
		t.Error(Error(st.Message()))
		return
	}
	t.Log(Success("Got participants: ") + Success(len(data.Participants)))
	for _, p := range data.Participants {
		t.Log(Success("Got participantId: ") + Success(p.ParticipantId))
	}
}
