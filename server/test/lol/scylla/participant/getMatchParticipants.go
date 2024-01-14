package participant

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc/status"
	. "meep.gg/grpc"
	. "meep.gg/log"
	. "meep.gg/protos/scylla/lol/participant/v1"
	. "meep.gg/test-server/lol/scylla"
)

func TestGetMatchParticipant(t *testing.T) {
	conn := GRPCConnect(LolScyllaAddr)
	defer conn.Close()
	client := NewScyllaParticipantServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Log("Grabbing puuids from match...")
	var puuids []string
	for _, participant := range TestMatch.Participants {
		puuids = append(puuids, participant.Puuid)
	}

	t.Log(Info("It should get participants"))
	data, err := client.GetMatchParticipants(ctx, &ScyllaMatchParticipantsReq{
		Region: "na1",
		GameId: TestMatch.GameId,
		Puuids: puuids,
	})
	if err != nil {
		st := status.Convert(err)
		t.Error(Error(st.Message()))
		return
	}

	for _, participant := range data.Participants {
		t.Log(Success("Got Participant Champion: ") + Success(participant.ChampionId))
	}
}
