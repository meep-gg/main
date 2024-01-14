package account

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	. "meep.gg/grpc"
	. "meep.gg/log"
	. "meep.gg/protos/gateway/riot/account/v1"
	. "meep.gg/test-server/riot/gateway"
)

func TestGetAccount(t *testing.T) {
	conn := GRPCConnect(GatewayAddr)
	defer conn.Close()
	client := NewGatewayAccountServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	t.Log(Info("It should get account by puuid"))
	account, err := client.GetAccount(ctx, &GatewayAccountReq{
		Region: "na1",
		Puuid:  TestAccount.Puuid,
	})

	if err != nil {
		t.Error(Error(err))
		return
	}
	if account == nil {
		t.Error(Error("Account is nil"))
		return
	}

	if TestAccount.Puuid != account.Puuid {
		t.Error(Error("Account puuid is not equal"))
	}
	t.Log(Success("Account puuid is equal to TestAccount puuid"))

	t.Log(Info("It should return empty account if not found"))
	account, err = client.GetAccount(ctx, &GatewayAccountReq{
		Region: "na1",
		Puuid:  "9WX5pU7pGi8EQvSFf6nbYf4n_pLI8s",
	})
	if err != nil {
		st := status.Convert(err)
		if st.Code() != codes.NotFound {
			t.Error(Error(st.Message()))
			return
		}
	}
	if account != nil {
		t.Error(Error("Account is not nil"))
	}
	t.Log(Success("Account is nil"))
}
