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

func TestPaginateParticipant(t *testing.T) {
	conn := GRPCConnect(LolScyllaAddr)
	defer conn.Close()
	client := NewScyllaParticipantServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Log(Info("It should get the first 10 rows for a participant"))
	data, err := client.PaginateParticipant(ctx, &ScyllaPaginateParticipantsReq{
		Region: "na1",
		Puuid:  "nSR_DUZ1B2PI9bjXmfnq7w0_0bhNkby1MeIhTnbBR1AoQVgAXXJzlhzVhwnuwenoc7M-O9od7UiK6g",
		Count:  10,
	})
	if err != nil {
		st := status.Convert(err)
		t.Error(Error(st.Message()))
		return
	}
	for _, participant := range data.Participants {
		t.Log(Success("Got GameId: ") + Success(participant.GameId))
	}

	t.Log(Info("It should get the next 10 rows for a participant"))
	data, err = client.PaginateParticipant(ctx, &ScyllaPaginateParticipantsReq{
		Region: "na1",
		Puuid:  "nSR_DUZ1B2PI9bjXmfnq7w0_0bhNkby1MeIhTnbBR1AoQVgAXXJzlhzVhwnuwenoc7M-O9od7UiK6g",
		Count:  10,
		GameId: data.Participants[2].GameId,
	})

	if err != nil {
		st := status.Convert(err)
		t.Error(Error(st.Message()))
		return
	}
	for _, participant := range data.Participants {
		t.Log(Success("Got GameId: ") + Success(participant.GameId))
	}

}
