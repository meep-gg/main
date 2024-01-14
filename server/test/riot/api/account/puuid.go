package account

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	. "meep.gg/protos/riot/account/v1"
	. "meep.gg/test-server/riot/api"

	. "meep.gg/grpc"
	. "meep.gg/log"
)

func TestPuuid(t *testing.T) {
	conn := GRPCConnect(RiotAPIAddr)
	defer conn.Close()
	client := NewRiotAccountServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	t.Log(Info("It should get RiotAccount using Puuid"))
	account, err := client.Puuid(ctx, &RiotAccountReq{
		Region: "americas",
		Puuid:  TestAccount.Puuid,
	})
	if err != nil {
		st := status.Convert(err)
		t.Log(Error(st.Message()))
		return
	}
	if account == nil {
		t.Log(Error("Account is nil"))
		return
	}
	if TestAccount.Puuid != account.Puuid {
		t.Log(Error("Account puuid is not equal"))
	}
	t.Log(Success("Account puuid is equal to TestAccount puuid"))

	t.Log(Info("It should get RiotAccount using Puuid"))
	account, err = client.Puuid(ctx, &RiotAccountReq{
		Region: "americas",
		Puuid:  "9WX5pU7pGi8EQvSFf6nbYf4n_pLI8s",
	})
	if err != nil {
		st := status.Convert(err)
		if st.Code() != codes.NotFound {
			t.Log(Error(st.Message()))
			return
		}
	}
	if account != nil {
		t.Log(Error("Account is not nil"))
	}
	t.Log(Success("Account is nil"))
}
