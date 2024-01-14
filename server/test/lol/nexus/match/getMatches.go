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

func TestGetMatches(t *testing.T) {
	conn := GRPCConnect(LolNexusAddr)
	defer conn.Close()
	client := NewMatchServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Log(Info("It should get first 10 matches and participant from Scylla"))
	data, err := client.GetMatches(ctx, &GetMatchesReq{
		Region: "na1",
		Puuid:  "nSR_DUZ1B2PI9bjXmfnq7w0_0bhNkby1MeIhTnbBR1AoQVgAXXJzlhzVhwnuwenoc7M-O9od7UiK6g",
	})
	if err != nil {
		st := status.Convert(err)
		t.Error(Error(st.Message()))
		return
	}
	for i := 0; i < len(data.Matches); i++ {
		t.Log(Success("Got GameId: ") + Success(data.Matches[i].GameId))
		t.Log(Success("Got Participant ChampionId: ") + Success(data.Participants[i].ChampionId))
	}

	t.Log(Info("It should get next 10 matches and participant from Scylla"))
	data, err = client.GetMatches(ctx, &GetMatchesReq{
		Region: "na1",
		Puuid:  "nSR_DUZ1B2PI9bjXmfnq7w0_0bhNkby1MeIhTnbBR1AoQVgAXXJzlhzVhwnuwenoc7M-O9od7UiK6g",
		GameId: data.Matches[len(data.Matches)-1].GameId,
	})
	if err != nil {
		st := status.Convert(err)
		t.Error(Error(st.Message()))
		return
	}

	for i := 0; i < len(data.Matches); i++ {
		t.Log(Success("Got GameId: ") + Success(data.Matches[i].GameId))
		t.Log(Success("Got Participant ChampionId: ") + Success(data.Participants[i].ChampionId))
	}
}
