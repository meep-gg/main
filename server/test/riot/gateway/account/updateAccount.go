package account

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc/status"
	. "meep.gg/grpc"
	. "meep.gg/log"
	. "meep.gg/protos/gateway/riot/account/v1"
	. "meep.gg/test-server/riot/gateway"
)

func TestUpdateAccount(t *testing.T) {
	conn := GRPCConnect(GatewayAddr)
	defer conn.Close()
	client := NewGatewayAccountServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	t.Log(Info("It should update account by puuid"))
	account, err := client.UpdateAccount(ctx, &GatewayAccountReq{
		Region: "na1",
		Puuid:  TestAccount.Puuid,
	})
	if err != nil {
		st := status.Convert(err)
		t.Error(Error(st.Message()))
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

	t.Log(Info("It should update account by RiotId"))
	account, err = client.UpdateAccount(ctx, &GatewayAccountReq{
		Region:   "na1",
		GameName: TestAccount.GameName,
		TagLine:  TestAccount.TagLine,
	})

	if err != nil {
		st := status.Convert(err)
		t.Error(Error(st.Message()))
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
}
