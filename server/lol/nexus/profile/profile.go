package profile

import (
	"context"
	"log"
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	. "meep.gg/log"
	. "meep.gg/protos/nexus/lol/profile/v1"

	. "meep.gg/protos/scylla/lol/league/v1"
	. "meep.gg/protos/scylla/lol/profile/v1"
	. "meep.gg/protos/scylla/riot/account/v1"

	. "meep.gg/protos/gateway/lol/league/v1"
	. "meep.gg/protos/gateway/lol/summoner/v1"

	. "meep.gg/protos/riot/account/v1"
	. "meep.gg/types/lol"

	. "meep.gg/lol-nexus/connect"
)

type ProfileServer struct {
	ProfileServiceServer
}

func (s *ProfileServer) GetProfile(ctx context.Context, req *ProfileReq) (*GetProfileRes, error) {
	region := req.GetRegion()
	gameName := req.GetGameName()
	tagLine := req.GetTagLine()

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	if gameName == "" || tagLine == "" {
		return nil, status.Error(codes.InvalidArgument, "GameName or TagLine is empty")
	}

	var profileResponse = &GetProfileRes{}
	log.Printf(Info("Getting profile"))
	// Grab profile data
	profile, err := ScyllaProfileClient.GetProfile(ctx, &ScyllaProfileReq{
		Region:   region,
		GameName: gameName,
		TagLine:  tagLine,
	})
	if err != nil {
		st := status.Convert(err)
		if st.Code() != codes.NotFound {
			return nil, status.Error(codes.Internal, st.Message())
		}
		log.Printf(Warn("Profile not found in scylladb"))
		return nil, status.Error(codes.NotFound, "Profile not found")
	}

	if profile == nil {
		return nil, status.Error(codes.NotFound, "Profile not found")
	}
	profileResponse.Profile = profile

	// Grab league data
	leagues, err := ScyllaLeagueClient.GetLeagues(ctx, &ScyllaLeagueReq{
		Region:     region,
		SummonerId: profile.SummonerId,
	})
	if err != nil {
		st := status.Convert(err)
		if st.Code() != codes.NotFound {
			return nil, status.Error(codes.Internal, st.Message())
		}
		log.Printf(Warn("Leagues not found in scylladb"))
	}
	if leagues.Entries == nil {
		return nil, status.Error(codes.NotFound, "Leagues not found")
	}
	if len(leagues.Entries) > 0 {
		for _, league := range leagues.Entries {
			if league.QueueType == "RANKED_SOLO_5x5" {
				profileResponse.RankedSolo = league
			} else if league.QueueType == "RANKED_FLEX_SR" {
				profileResponse.RankedFlex = league
			}
		}
	}

	log.Printf(Success("Profile found"))
	return profileResponse, nil
}

func (s *ProfileServer) UpdateProfile(ctx context.Context, req *ProfileReq) (*GetProfileRes, error) {
	region := req.GetRegion()
	gameName := req.GetGameName()
	tagLine := req.GetTagLine()

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	if gameName == "" || tagLine == "" {
		return nil, status.Error(codes.InvalidArgument, "GameName or TagLine is empty")
	}

	// Check Riot Games if Account exists
	regional, err := PlatformToRegional(region)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	riotAccount, err := RiotAccountClient.RiotId(ctx, &RiotAccountReq{
		Region:   regional,
		GameName: gameName,
		TagLine:  tagLine,
	})
	if err != nil {
		st := status.Convert(err)
		return nil, status.Error(codes.Internal, st.Message())
	}
	if riotAccount == nil {
		return nil, status.Error(codes.NotFound, "RiotAccount not found")
	}

	// Check if Account exists in Scylla
	scyllaAccount, err := ScyllaAccountClient.GetAccount(ctx, &ScyllaAccountReq{
		Region: region,
		Puuid:  riotAccount.Puuid,
	})
	if err != nil {
		st := status.Convert(err)
		if st.Code() != codes.NotFound {
			return nil, status.Error(codes.Internal, st.Message())
		}
		log.Printf(Warn("Account not found in scylladb"))
	}
	if scyllaAccount != nil {
		// Check if GameName or TagLine has changed
		if scyllaAccount.GameName != riotAccount.GameName || scyllaAccount.TagLine != riotAccount.TagLine {
			// Delete old profile
			_, err := ScyllaProfileClient.DeleteProfile(ctx, &ScyllaProfileReq{
				Region:   region,
				GameName: scyllaAccount.GameName,
				TagLine:  scyllaAccount.TagLine,
			})
			if err != nil {
				st := status.Convert(err)
				return nil, status.Error(codes.Internal, st.Message())
			}
			// Update Account
			_, err = ScyllaAccountClient.UpsertAccount(ctx, &ScyllaUpsertAccountReq{
				Region: region,
				Account: &ScyllaAccount{
					Puuid:    riotAccount.Puuid,
					GameName: riotAccount.GameName,
					TagLine:  riotAccount.TagLine,
				},
			})
			if err != nil {
				st := status.Convert(err)
				return nil, status.Error(codes.Internal, st.Message())
			}
			// TODO Update Names in Matches*****************
		}
	} else {
		// Update Account
		_, err = ScyllaAccountClient.UpsertAccount(ctx, &ScyllaUpsertAccountReq{
			Region: region,
			Account: &ScyllaAccount{
				Puuid:    riotAccount.Puuid,
				GameName: riotAccount.GameName,
				TagLine:  riotAccount.TagLine,
			},
		})
		if err != nil {
			st := status.Convert(err)
			return nil, status.Error(codes.Internal, st.Message())
		}
	}

	// Update Summoner from gateway
	summoner, err := GatewaySummonerClient.UpdateSummoner(ctx, &GatewaySummonerReq{
		Region: region,
		Puuid:  riotAccount.Puuid,
	})
	if err != nil {
		st := status.Convert(err)
		return nil, status.Error(codes.Internal, st.Message())
	}
	if summoner == nil {
		return nil, status.Error(codes.NotFound, "Summoner not found")
	}

	log.Printf(Info("Creating Profile"))
	profileResponse := &GetProfileRes{
		Profile: &ScyllaProfile{
			GameName:      riotAccount.GameName,
			TagLine:       riotAccount.TagLine,
			Puuid:         riotAccount.Puuid,
			SummonerId:    summoner.Id,
			ProfileIconId: summoner.ProfileIconId,
			SummonerLevel: summoner.SummonerLevel,
		},
	}

	// Initialize the wait group and error channels
	var wg sync.WaitGroup
	errChan := make(chan error, 2)

	wg.Add(1)
	go func() {
		defer wg.Done()
		_, err := ScyllaProfileClient.UpsertProfile(ctx, &ScyllaUpsertProfileReq{
			Region:  region,
			Profile: profileResponse.Profile,
		})
		if err != nil {
			errChan <- err
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		leagues, err := GatewayLeagueClient.UpdateLeagues(ctx, &GatewayLeagueReq{
			Region:     region,
			SummonerId: summoner.Id,
		})
		if err != nil {
			errChan <- err
			return
		}

		if leagues.Entries != nil {
			for _, league := range leagues.Entries {
				if league.QueueType == "RANKED_SOLO_5x5" {
					profileResponse.RankedSolo = league
				} else if league.QueueType == "RANKED_FLEX_SR" {
					profileResponse.RankedFlex = league
				}
			}
		}
	}()

	wg.Wait()
	close(errChan)

	for err := range errChan {
		st := status.Convert(err)
		if st.Code() != codes.NotFound {
			return nil, status.Error(codes.Internal, st.Message())
		}
		log.Printf(Warn("An operation failed in gateway"))
	}

	log.Printf(Success("Profile created"))
	return profileResponse, nil
}
