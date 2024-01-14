package match

import (
	"context"
	"sync"
	"testing"
	"time"

	. "meep.gg/grpc"
	. "meep.gg/log"
	. "meep.gg/protos/riot/lol/match/v1"
	. "meep.gg/test-server/lol/api"
)

func TestStress(t *testing.T) {
	conn := GRPCConnect(LolAPIAddr)
	defer conn.Close()
	client := NewRiotMatchServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Log(Info("It should get Matches by using MatchId from a list"))
	var wg sync.WaitGroup
	type result struct {
		matchId string
		match   *RiotMatch
	}

	errors := make(chan error, len(Matches))

	for _, matchId := range Matches {
		wg.Add(1)
		go func(matchId string) {
			defer wg.Done()
			match, err := client.Id(ctx, &RiotMatchIdReq{
				Region:  "americas",
				MatchId: matchId,
			})
			if err != nil {
				errors <- err
				return
			}
			if match == nil {
				errors <- err
				return
			}
			if match.Metadata.MatchId != matchId {
				errors <- err
				return
			}
			t.Log(Success("Retrevied match by match id: ") + Success(match.Metadata.MatchId))
		}(matchId)
	}
	wg.Wait()
	close(errors)
	for err := range errors {
		t.Error(Error(err))
	}
	t.Log(Success("Retrevied all matches by match id"))
}
