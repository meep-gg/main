package league

import (
	"context"
	"log"
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	. "meep.gg/log"
	. "meep.gg/protos/gateway/lol/league/v1"
	. "meep.gg/protos/riot/lol/league/v1"
	. "meep.gg/protos/scylla/lol/league/v1"
	. "meep.gg/types/lol"
)

type GatewayLeagueServer struct {
	GatewayLeagueServiceServer
}

func (s *GatewayLeagueServer) GetLeagues(ctx context.Context, in *GatewayLeagueReq) (*ScyllaLeagueEntries, error) {
	region := in.GetRegion()
	summonerId := in.GetSummonerId()

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	if summonerId == "" {
		return nil, status.Error(codes.InvalidArgument, "SummonerId is empty")
	}

	log.Printf(Info("Getting league from scylladb"))
	data, err := ScyllaLeagueClient.GetLeagues(ctx, &ScyllaLeagueReq{
		Region:     region,
		SummonerId: summonerId,
	})
	if err != nil {
		st := status.Convert(err)
		if st.Code() != codes.NotFound {
			return nil, status.Error(codes.Internal, st.Message())
		}
		log.Printf(Warn("League not found in scylladb"))
	}
	if data != nil {
		log.Printf(Success("Got league from scylladb: ") + Success(len(data.Entries)))
		return data, nil
	}

	log.Printf(Info("Getting league from riot"))
	RiotLeague, err := RiotLeagueClient.SummonerId(ctx, &RiotLeagueSummonerIdReq{
		Region:     region,
		SummonerId: summonerId,
	})

	if err != nil {
		st := status.Convert(err)
		if st.Code() == codes.NotFound {
			log.Printf(Warn("League not found in riot"))
			return nil, status.Error(codes.NotFound, st.Message())
		}
		return nil, status.Error(codes.Internal, st.Message())
	}

	log.Printf(Success("Got league from riot: ") + Success(len(RiotLeague.Leagues)))

	var leagues = SerializeEntries(RiotLeague.Leagues)
	for _, league := range leagues {
		log.Printf(Info("Upserting league to scylladb"))
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		_, err = ScyllaLeagueClient.UpsertLeague(ctx, &ScyllaUpsertLeagueReq{
			Region: region,
			League: league,
		})
		if err != nil {
			st := status.Convert(err)
			return nil, status.Error(codes.Internal, st.Message())
		}
		log.Printf(Success("Upserted league to scylladb: ") + Success(league.SummonerId))
	}

	return &ScyllaLeagueEntries{
		Entries: leagues,
	}, nil
}

func (s *GatewayLeagueServer) UpdateLeagues(ctx context.Context, in *GatewayLeagueReq) (*ScyllaLeagueEntries, error) {
	region := in.GetRegion()
	summonerId := in.GetSummonerId()

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	if summonerId == "" {
		return nil, status.Error(codes.InvalidArgument, "SummonerId is empty")
	}

	log.Printf(Info("Getting league from riot"))
	RiotLeague, err := RiotLeagueClient.SummonerId(ctx, &RiotLeagueSummonerIdReq{
		Region:     region,
		SummonerId: summonerId,
	})

	if err != nil {
		st := status.Convert(err)
		if st.Code() == codes.NotFound {
			log.Printf(Warn("League not found in riot"))
			return nil, status.Error(codes.NotFound, st.Message())
		}
		return nil, status.Error(codes.Internal, st.Message())
	}

	log.Printf(Success("Got league from riot: ") + Success(len(RiotLeague.Leagues)))
	var leagues = SerializeEntries(RiotLeague.Leagues)
	for _, league := range leagues {
		log.Printf(Info("Upserting league to scylladb"))
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		_, err = ScyllaLeagueClient.UpsertLeague(ctx, &ScyllaUpsertLeagueReq{
			Region: region,
			League: league,
		})
		if err != nil {
			st := status.Convert(err)
			return nil, status.Error(codes.Internal, st.Message())
		}
		log.Printf(Success("Upserted league to scylladb: ") + Success(league.SummonerId))
	}

	return &ScyllaLeagueEntries{
		Entries: leagues,
	}, nil
}

func (s *GatewayLeagueServer) UpdateChallenger(ctx context.Context, in *GatewayApexUpdateReq) (*emptypb.Empty, error) {
	region := in.GetRegion()
	queueType := in.GetQueueType()

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	if !IsQueue(queueType) {
		return nil, status.Error(codes.InvalidArgument, "QueueType is not of type Queue")
	}

	log.Printf(Info("Getting challenger from riot"))
	RiotChallenger, err := RiotLeagueClient.Challenger(ctx, &RiotLeagueApexReq{
		Region: region,
		Queue:  queueType,
	})
	if err != nil {
		st := status.Convert(err)
		return nil, status.Error(codes.Internal, st.Message())
	}

	log.Printf(Success("Got challenger from riot: ") + Success(len(RiotChallenger.Entries)))
	var wg sync.WaitGroup
	errChan := make(chan error, len(RiotChallenger.Entries))

	for _, league := range RiotChallenger.Entries {
		wg.Add(1)
		go func(league *LeagueItem) {
			defer wg.Done()

			select {
			case <-ctx.Done():
				errChan <- ctx.Err()
				return
			default:
				_, err := ScyllaLeagueClient.UpsertLeague(ctx, &ScyllaUpsertLeagueReq{
					Region: region,
					League: SerializeItem("CHALLENGER", "I", queueType, league),
				})
				if err != nil {
					errChan <- err
				} else {
					log.Printf(Success("Upserted league to scylladb: ") + Success(league.SummonerId))
				}
			}
		}(league)
	}

	wg.Wait()
	close(errChan)

	return &emptypb.Empty{}, nil
}

func (s *GatewayLeagueServer) UpdateGrandmaster(ctx context.Context, in *GatewayApexUpdateReq) (*emptypb.Empty, error) {
	region := in.GetRegion()
	queueType := in.GetQueueType()

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	if !IsQueue(queueType) {
		return nil, status.Error(codes.InvalidArgument, "QueueType is not of type Queue")
	}

	log.Printf(Info("Getting grandmaster from riot"))
	RiotGrandmaster, err := RiotLeagueClient.Grandmaster(ctx, &RiotLeagueApexReq{
		Region: region,
		Queue:  queueType,
	})
	if err != nil {
		st := status.Convert(err)
		return nil, status.Error(codes.Internal, st.Message())
	}

	log.Printf(Success("Got grandmaster from riot: ") + Success(len(RiotGrandmaster.Entries)))
	var wg sync.WaitGroup
	errChan := make(chan error, len(RiotGrandmaster.Entries))

	for _, league := range RiotGrandmaster.Entries {
		wg.Add(1)
		go func(league *LeagueItem) {
			defer wg.Done()

			select {
			case <-ctx.Done():
				errChan <- ctx.Err()
				return
			default:
				_, err := ScyllaLeagueClient.UpsertLeague(ctx, &ScyllaUpsertLeagueReq{
					Region: region,
					League: SerializeItem("GRANDMASTER", "I", queueType, league),
				})
				if err != nil {
					errChan <- err
				} else {
					log.Printf(Success("Upserted league to scylladb: ") + Success(league.SummonerId))
				}
			}
		}(league)
	}

	wg.Wait()
	close(errChan)

	return &emptypb.Empty{}, nil
}

func (s *GatewayLeagueServer) UpdateMaster(ctx context.Context, in *GatewayApexUpdateReq) (*emptypb.Empty, error) {
	region := in.GetRegion()
	queueType := in.GetQueueType()

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	if !IsQueue(queueType) {
		return nil, status.Error(codes.InvalidArgument, "QueueType is not of type Queue")
	}

	log.Printf(Info("Getting master from riot"))
	RiotMaster, err := RiotLeagueClient.Master(ctx, &RiotLeagueApexReq{
		Region: region,
		Queue:  queueType,
	})
	if err != nil {
		st := status.Convert(err)
		return nil, status.Error(codes.Internal, st.Message())
	}

	log.Printf(Success("Got master from riot: ") + Success(len(RiotMaster.Entries)))
	var wg sync.WaitGroup
	errChan := make(chan error, len(RiotMaster.Entries))

	for _, league := range RiotMaster.Entries {
		wg.Add(1)
		go func(league *LeagueItem) {
			defer wg.Done()

			select {
			case <-ctx.Done():
				errChan <- ctx.Err()
				return
			default:
				_, err := ScyllaLeagueClient.UpsertLeague(ctx, &ScyllaUpsertLeagueReq{
					Region: region,
					League: SerializeItem("MASTER", "I", queueType, league),
				})
				if err != nil {
					errChan <- err
				} else {
					log.Printf(Success("Upserted league to scylladb: ") + Success(league.SummonerId))
				}
			}
		}(league)
	}

	wg.Wait()
	close(errChan)

	return &emptypb.Empty{}, nil
}

func (s *GatewayLeagueServer) UpdateEntry(ctx context.Context, in *GatewayUpdateEntryReq) (*GatewayUpdateEntryRes, error) {
	region := in.GetRegion()
	queueType := in.GetQueueType()
	tier := in.GetTier()
	division := in.GetDivision()
	page := in.GetPage()

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	if !IsQueue(queueType) {
		return nil, status.Error(codes.InvalidArgument, "QueueType is not of type Queue")
	}

	if tier == "" || division == "" {
		return nil, status.Error(codes.InvalidArgument, "Tier or Division is empty")
	}

	if page == 0 {
		page = 1
	}

	log.Printf(Info("Getting entries from riot"))
	RiotEntries, err := RiotLeagueClient.Entry(ctx, &RiotLeagueEntryReq{
		Region:   region,
		Queue:    queueType,
		Tier:     tier,
		Division: division,
		Page:     page,
	})

	if err != nil {
		st := status.Convert(err)
		if st.Code() == codes.NotFound {
			log.Printf(Warn("Entries not found in riot"))
			return nil, status.Error(codes.NotFound, st.Message())
		}
		return nil, status.Error(codes.Internal, st.Message())
	}

	log.Printf(Success("Got entries from riot: ") + Success(len(RiotEntries.Leagues)))
	var wg sync.WaitGroup
	errChan := make(chan error, len(RiotEntries.Leagues))

	for _, league := range RiotEntries.Leagues {
		wg.Add(1)
		go func(league *RiotLeagueEntry) {
			defer wg.Done()

			select {
			case <-ctx.Done():
				errChan <- ctx.Err()
				return
			default:
				_, err := ScyllaLeagueClient.UpsertLeague(ctx, &ScyllaUpsertLeagueReq{
					Region: region,
					League: SerializeEntry(league),
				})
				if err != nil {
					log.Printf(Error("Upserting league to scylladb: ") + Error(err.Error()))
					errChan <- err
				} else {
					log.Printf(Success("Upserted league to scylladb: ") + Success(league.SummonerId))
				}
			}
		}(league)
	}

	wg.Wait()
	close(errChan)

	return &GatewayUpdateEntryRes{
		Documents: int32(len(RiotEntries.Leagues)),
	}, nil
}
