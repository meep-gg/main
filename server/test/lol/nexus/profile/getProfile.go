package profile

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc/status"
	. "meep.gg/grpc"
	. "meep.gg/log"
	. "meep.gg/protos/nexus/lol/profile/v1"
	. "meep.gg/test-server/lol/nexus"
)

func TestGetProfile(t *testing.T) {
	conn := GRPCConnect(LolNexusAddr)
	defer conn.Close()
	client := NewProfileServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Log("Upserting profile...")
	_, err := client.UpdateProfile(ctx, &ProfileReq{
		Region:   "na1",
		GameName: TestAccount.GameName,
		TagLine:  TestAccount.TagLine,
	})
	if err != nil {
		st := status.Convert(err)
		t.Error(Error(st.Message()))
		return
	}

	t.Log(Info("It should get a profile from Scylla"))
	profile, err := client.GetProfile(ctx, &ProfileReq{
		Region:   "na1",
		GameName: TestAccount.GameName,
		TagLine:  TestAccount.TagLine,
	})

	if err != nil {
		st := status.Convert(err)
		t.Error(Error(st.Message()))
		return
	}
	t.Log(Success("Got profile: ") + Success(profile.Profile.GameName))
}
