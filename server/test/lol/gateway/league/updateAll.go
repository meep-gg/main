package league

import (
	"context"
	"sync"
	"testing"
	"time"

	"meep.gg/grpc"
	"meep.gg/log"
	v1 "meep.gg/protos/gateway/lol/league/v1"
	"meep.gg/test-server/lol/gateway"
)

type Task struct {
	Tier     string
	Division string
	Page     int32
}

func worker(ctx context.Context, client v1.GatewayLeagueServiceClient, tasks <-chan Task, wg *sync.WaitGroup, t *testing.T) {
	for task := range tasks {
		resp, err := client.UpdateEntry(ctx, &v1.GatewayUpdateEntryReq{
			Region:    "na1",
			QueueType: "RANKED_SOLO_5x5",
			Tier:      task.Tier,
			Division:  task.Division,
			Page:      task.Page,
		})
		if err != nil {
			t.Errorf("Failed to update entry for tier %s, division %s, page %d: %v", task.Tier, task.Division, task.Page, err)
			return
		}
		if resp.Documents == 0 {
			wg.Done()
		}
	}
}

// func TestUpdateAll(t *testing.T) {
// 	conn := grpc.GRPCConnect(gateway.LolGatewayAddr)
// 	defer conn.Close()
// 	client := v1.NewGatewayLeagueServiceClient(conn)

// 	t.Log(log.Info("It should update all solo leagues"))
// 	tiers := []string{"IRON", "BRONZE", "SILVER", "GOLD", "EMERALD", "PLATINUM", "DIAMOND"}
// 	divisions := []string{"I", "II", "III", "IV"}

// 	ctx, cancel := context.WithTimeout(context.Background(), 100000*time.Second)
// 	defer cancel()

// 	// Number of workers
// 	numWorkers := 5
// 	tasks := make(chan Task, 50) // Buffer size of 50

// 	var wg sync.WaitGroup

// 	// Start workers
// 	for w := 0; w < numWorkers; w++ {
// 		go worker(ctx, client, tasks, &wg, t)
// 	}

// 	// Send tasks to workers
// 	for _, tier := range tiers {
// 		for _, division := range divisions {
// 			for page := 1; page <= 10; page++ { // Assuming 10 pages per division
// 				wg.Add(1)
// 				tasks <- Task{Tier: tier, Division: division, Page: uint32(page)}
// 				if len(tasks) == cap(tasks) {
// 					time.Sleep(10 * time.Second) // Sleep for 10 seconds after pushing 50 tasks
// 				}
// 			}
// 		}
// 	}

// 	close(tasks) // Close channel when all tasks are sent
// 	wg.Wait()    // Wait for all tasks to complete
// }

func TestUpdateAll(t *testing.T) {
	conn := grpc.GRPCConnect(gateway.LolGatewayAddr)
	defer conn.Close()
	client := v1.NewGatewayLeagueServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 100000000*time.Second)
	defer cancel()

	t.Log(log.Info("It should update all solo leagues"))
	tiers := []string{"IRON", "BRONZE", "SILVER", "GOLD", "PLATINUM", "EMERALD", "DIAMOND"}
	divisions := []string{"I", "II", "III", "IV"}

	for _, tier := range tiers {
		for _, division := range divisions {
			page := 1
			for {
				resp, err := client.UpdateEntry(ctx, &v1.GatewayUpdateEntryReq{
					Region:    "na1",
					QueueType: "RANKED_SOLO_5x5",
					Tier:      tier,
					Division:  division,
					Page:      int32(page),
				})
				if err != nil {
					t.Errorf("Failed to update entry for tier %s, division %s, page %d: %v", tier, division, page, err)
					return
				}

				if resp.Documents == 0 {
					break
				}
				page++
			}
		}
	}
}
