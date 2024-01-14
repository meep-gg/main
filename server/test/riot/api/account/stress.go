package account

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	. "meep.gg/protos/riot/account/v1"
	. "meep.gg/test-server/riot/api"

	. "meep.gg/grpc"
	. "meep.gg/log"
)

func TestStress(t *testing.T) {
	conn := GRPCConnect(RiotAPIAddr)
	defer conn.Close()
	client := NewRiotAccountServiceClient(conn)

	t.Log(Info("It should get a list of RiotAccounts with Puuid"))
	var wg sync.WaitGroup
	errors := make(chan error, len(PuuidList))

	for _, puuid := range PuuidList {
		wg.Add(1)

		go func(p string) {
			defer wg.Done()
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			account, err := client.Puuid(ctx, &RiotAccountReq{
				Region: "americas",
				Puuid:  p,
			})
			if err != nil {
				errors <- fmt.Errorf("error for puuid %s: %s", p, err.Error())
				return
			}
			if account == nil {
				errors <- fmt.Errorf("account is nil for puuid %s", p)
				return
			}

			if account.Puuid != p {
				errors <- fmt.Errorf("Account puuid does not match for %s", p)
			}

			t.Log(Success(fmt.Sprintf("%s#%s", account.GameName, account.TagLine)))
		}(puuid)
	}

	wg.Wait()
	close(errors)

	for err := range errors {
		t.Error(err)
	}
}
