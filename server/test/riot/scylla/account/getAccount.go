package account

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	. "meep.gg/grpc"
	. "meep.gg/log"
	. "meep.gg/protos/scylla/riot/account/v1"
	. "meep.gg/test-server/riot/scylla"
)

func TestGetAccount(t *testing.T) {
	conn := GRPCConnect(ScyllaAddr)
	defer conn.Close()
	client := NewScyllaAccountServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Log("Importing Test Account...")
	_, err := client.UpsertAccount(ctx, &ScyllaUpsertAccountReq{
		Region:  "na1",
		Account: TestAccount,
	})
	if err != nil {
		st := status.Convert(err)
		t.Log(Error(st.Message()))
		return
	}

	t.Log(Info("It should get Account"))
	account, err := client.GetAccount(ctx, &ScyllaAccountReq{
		Region: "na1",
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
	t.Log(Success("Account is not nil"))

	t.Log(Info("It should return notFound err if not found"))
	account, err = client.GetAccount(ctx, &ScyllaAccountReq{
		Region: "na1",
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
	t.Log(Success("notFound err returned"))
}
