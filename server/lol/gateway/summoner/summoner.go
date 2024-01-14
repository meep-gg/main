package summoner

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	. "meep.gg/log"
	. "meep.gg/protos/gateway/lol/summoner/v1"
	. "meep.gg/protos/riot/lol/summoner/v1"
	. "meep.gg/protos/scylla/lol/summoner/v1"
	. "meep.gg/types/lol"
)

type GatewaySummonerServer struct {
	GatewaySummonerServiceServer
}

func (s *GatewaySummonerServer) GetSummoner(ctx context.Context, in *GatewaySummonerReq) (*ScyllaSummoner, error) {
	region := in.GetRegion()
	puuid := in.GetPuuid()

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	if puuid == "" {
		return nil, status.Error(codes.InvalidArgument, "Puuid is empty")
	}

	log.Printf(Info("Getting summoner from scylladb"))
	summoner, err := ScyllaSummonerClient.GetSummoner(ctx, &ScyllaSummonerReq{
		Region: region,
		Puuid:  puuid,
	})
	if err != nil {
		st := status.Convert(err)
		if st.Code() != codes.NotFound {
			return nil, status.Error(codes.Internal, st.Message())
		}
		log.Printf(Warn("Summoner not found in scylladb"))
	}
	if summoner != nil {
		log.Printf(Success("Got summoner from scylladb: ") + Success(summoner.Puuid))
		return summoner, nil
	}

	log.Printf(Info("Getting summoner from riot"))
	RiotSummoner, err := RiotSummonerClient.Puuid(ctx, &RiotSummonerPuuidReq{
		Region: region,
		Puuid:  puuid,
	})
	if err != nil {
		st := status.Convert(err)
		if st.Code() == codes.NotFound {
			log.Printf(Warn("Summoner not found in riot"))
			return nil, status.Error(codes.NotFound, st.Message())
		}
		return nil, status.Error(codes.Internal, st.Message())
	}

	log.Printf(Info("Upserting summoner to scylladb"))
	summoner = SerializeRiotSummoner(RiotSummoner)
	_, err = ScyllaSummonerClient.UpsertSummoner(ctx, &ScyllaUpsertSummonerReq{
		Region:   region,
		Summoner: summoner,
	})
	if err != nil {
		st := status.Convert(err)
		return nil, status.Error(codes.Internal, st.Message())
	}

	log.Printf(Success("Got summoner from riot: ") + Success(summoner.Puuid))
	return summoner, nil
}

func (s *GatewaySummonerServer) UpdateSummoner(ctx context.Context, in *GatewaySummonerReq) (*ScyllaSummoner, error) {
	region := in.GetRegion()
	summonerId := in.GetSummonerId()
	puuid := in.GetPuuid()

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	if summonerId == "" && puuid == "" {
		return nil, status.Error(codes.InvalidArgument, "SummonerId, or Puuid is empty")
	}

	log.Printf(Info("Getting summoner from riot"))
	var RiotSummoner *RiotSummoner
	var err error
	if summonerId != "" {
		RiotSummoner, err = RiotSummonerClient.Id(ctx, &RiotSummonerIdReq{
			Region: region,
			Id:     summonerId,
		})
	} else {
		RiotSummoner, err = RiotSummonerClient.Puuid(ctx, &RiotSummonerPuuidReq{
			Region: region,
			Puuid:  puuid,
		})
	}
	if err != nil {
		st := status.Convert(err)
		if st.Code() == codes.NotFound {
			log.Printf(Warn("Summoner not found in riot"))
			return nil, status.Error(codes.NotFound, st.Message())
		}
		return nil, status.Error(codes.Internal, st.Message())
	}

	log.Printf(Info("Upserting summoner to scylladb"))
	summoner := SerializeRiotSummoner(RiotSummoner)
	_, err = ScyllaSummonerClient.UpsertSummoner(ctx, &ScyllaUpsertSummonerReq{
		Region:   region,
		Summoner: summoner,
	})
	if err != nil {
		st := status.Convert(err)
		return nil, status.Error(codes.Internal, st.Message())
	}

	log.Printf(Success("Got summoner from riot: ") + Success(summoner.Puuid))
	return summoner, nil
}
