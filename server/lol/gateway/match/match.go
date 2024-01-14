package match

import (
	"context"
	"log"
	"strconv"
	"strings"
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	. "meep.gg/protos/gateway/lol/match/v1"
	. "meep.gg/protos/riot/lol/match/v1"
	. "meep.gg/protos/scylla/lol/match/v1"

	. "meep.gg/log"
	. "meep.gg/types/lol"
)

type GatewayMatchServer struct {
	GatewayMatchServiceServer
}

func (s *GatewayMatchServer) UpdatePlayerMatches(ctx context.Context, req *GatewayMatchPlayerReq) (*emptypb.Empty, error) {
	region := req.GetRegion()
	puuid := req.GetPuuid()

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	if puuid == "" {
		return nil, status.Error(codes.InvalidArgument, "Puuid is empty")
	}

	regional, err := PlatformToRegional(region)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	var matchCounter int32 = 0

	for {
		data, err := RiotMatchClient.Puuid(ctx, &RiotMatchPuuidReq{
			Region:    regional,
			Puuid:     puuid,
			StartTime: SeasonStart14_1,
			Start:     matchCounter,
			Count:     100,
		})
		if err != nil {
			st := status.Convert(err)
			if st.Code() == codes.NotFound {
				return nil, status.Error(codes.NotFound, st.Message())
			}
			return nil, status.Error(codes.Internal, st.Message())
		}

		var wg sync.WaitGroup
		errChan := make(chan error, len(data.Matches))
		var matchesFound int32 = 0

		for _, match := range data.Matches {
			wg.Add(1)
			go func(match string) {
				defer wg.Done()
				split := strings.Split(match, "_")
				region := strings.ToLower(split[0])
				gameId64, err := strconv.ParseUint(split[1], 10, 64)
				if err != nil {
					errChan <- err
					return
				}
				gameId := int64(gameId64)
				_, err = ScyllaMatchClient.CheckMatch(ctx, &ScyllaMatchReq{
					Region: region,
					GameId: gameId,
				})
				if err != nil {
					st := status.Convert(err)
					if st.Code() != codes.NotFound {
						errChan <- err
						return
					}
					matchesFound = 0
					log.Printf(Info("Getting match from riot"))
					RiotMatch, err := RiotMatchClient.Id(ctx, &RiotMatchIdReq{
						Region:  regional,
						MatchId: match,
					})
					if err != nil {
						st := status.Convert(err)
						if st.Code() != codes.NotFound {
							errChan <- err
							return
						}
						log.Printf(Warn("Match not found in riot"))
					}
					if RiotMatch.Info == nil {
						errChan <- err
						return
					}
					log.Printf(Info("Upserting match to scylladb"))
					err = UpsertMatch(ctx, region, gameId, RiotMatch)
					if err != nil {
						errChan <- err
						return
					}
					log.Printf(Success("Got match from riot: ") + Success(match))
				} else {
					log.Printf(Success("Got match from scylladb: ") + Success(match))
					matchesFound++
				}
			}(match)
		}

		wg.Wait()
		close(errChan)

		for err := range errChan {
			st := status.Convert(err)
			return nil, status.Error(st.Code(), st.Message())
		}

		if matchesFound > 75 {
			break
		}
		// Check if the number of documents returned is less than 100
		if len(data.Matches) < 100 {
			break
		}
		matchCounter += 100
	}
	return &emptypb.Empty{}, nil
}
