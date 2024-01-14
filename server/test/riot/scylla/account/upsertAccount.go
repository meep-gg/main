package account

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc/status"
	. "meep.gg/grpc"
	. "meep.gg/log"
	. "meep.gg/protos/scylla/riot/account/v1"
	. "meep.gg/test-server/riot/scylla"
)

func TestUpsertAccount(t *testing.T) {
	conn := GRPCConnect(ScyllaAddr)
	defer conn.Close()
	client := NewScyllaAccountServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Log(Info("It should upsert Account"))
	_, err := client.UpsertAccount(ctx, &ScyllaUpsertAccountReq{
		Region:  "na1",
		Account: TestAccount,
	})
	if err != nil {
		st := status.Convert(err)
		t.Log(Error(st.Message()))
		return
	}

	t.Log(Success("Account upserted"))
}
