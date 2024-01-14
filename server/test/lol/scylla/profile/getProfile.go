package profile

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	. "meep.gg/grpc"
	. "meep.gg/log"
	. "meep.gg/protos/scylla/lol/profile/v1"
	. "meep.gg/test-server/lol/scylla"
)

func TestGetProfile(t *testing.T) {
	conn := GRPCConnect(LolScyllaAddr)
	defer conn.Close()
	client := NewScyllaProfileServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Log("Upserting profile...")
	_, err := client.UpsertProfile(ctx, &ScyllaUpsertProfileReq{
		Region:  "na1",
		Profile: TestProfile,
	})
	if err != nil {
		st := status.Convert(err)
		t.Error(Error(st.Message()))
		return
	}

	t.Log(Info("It should get a profile"))
	profile, err := client.GetProfile(ctx, &ScyllaProfileReq{
		Region:   "na1",
		GameName: TestProfile.GameName,
		TagLine:  TestProfile.TagLine,
	})

	if err != nil {
		st := status.Convert(err)
		t.Error(Error(st.Message()))
		return
	}
	t.Log(Success("Got profile: ") + Success(profile.GameName))

	t.Log(Info("It should return not found error if profile doesn't exist"))
	_, err = client.GetProfile(ctx, &ScyllaProfileReq{
		Region:   "na1",
		GameName: "122re3",
		TagLine:  "12w3",
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
