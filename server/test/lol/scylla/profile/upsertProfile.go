package profile

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc/status"
	. "meep.gg/grpc"
	. "meep.gg/log"
	. "meep.gg/protos/scylla/lol/profile/v1"
	. "meep.gg/test-server/lol/scylla"
)

func TestUpsertProfile(t *testing.T) {
	conn := GRPCConnect(LolScyllaAddr)
	defer conn.Close()
	client := NewScyllaProfileServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Log(Info("It should upsert a profile"))
	_, err := client.UpsertProfile(ctx, &ScyllaUpsertProfileReq{
		Region:  "na1",
		Profile: TestProfile,
	})
	if err != nil {
		st := status.Convert(err)
		t.Error(Error(st.Message()))
		return
	}
	t.Log(Success("Upserted profile: ") + Success(TestProfile.Puuid))
}
