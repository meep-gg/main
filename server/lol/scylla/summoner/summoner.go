package summoner

import (
	"context"
	"log"

	"github.com/gocql/gocql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	. "meep.gg/log"
	. "meep.gg/protos/scylla/lol/summoner/v1"
	. "meep.gg/scylla"
	. "meep.gg/types/lol"
)

type SummonerServer struct {
	ScyllaSummonerServiceServer
}

func (s *SummonerServer) GetSummoner(ctx context.Context, in *ScyllaSummonerReq) (*ScyllaSummoner, error) {
	region := in.GetRegion()
	puuid := in.GetPuuid()

	if Session == nil {
		return nil, status.Error(codes.Internal, "Scylla Session is nil")
	}

	if puuid == "" {
		return nil, status.Error(codes.InvalidArgument, "Puuid is empty")
	}

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	log.Printf(Info("Getting summoner"))
	var summoner = ScyllaSummoner{}
	filter := `SELECT * FROM ` + region + `_lol.summoner WHERE puuid = ? LIMIT 1`
	err := Session.Query(filter, puuid).Scan(
		&summoner.Puuid,
		&summoner.Id,
		&summoner.Name,
		&summoner.ProfileIconId,
		&summoner.SummonerLevel,
	)
	if err != nil {
		if err == gocql.ErrNotFound {
			log.Printf("Summoner not found: %v", Success(puuid))
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Printf("Retrieved Summoner: %v", Success(summoner.Name))
	return &summoner, nil
}

func (s *SummonerServer) UpsertSummoner(ctx context.Context, in *ScyllaUpsertSummonerReq) (*emptypb.Empty, error) {
	region := in.GetRegion()
	summoner := in.GetSummoner()

	if Session == nil {
		return nil, status.Error(codes.Internal, "Scylla Session is nil")
	}

	if summoner == nil {
		return nil, status.Error(codes.InvalidArgument, "Summoner is nil")
	}

	if summoner.Puuid == "" && summoner.Id == "" && summoner.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "Summoner is empty")
	}

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	log.Printf(Info("Upserting summoner"))
	query := `INSERT INTO ` + region + `_lol.summoner (
		puuid, 
		id, 
		name, 
		profileIconId, 
		summonerLevel
		) VALUES (?, ?, ?, ?, ?)`
	if err := Session.Query(query,
		summoner.Puuid,
		summoner.Id,
		summoner.Name,
		summoner.ProfileIconId,
		summoner.SummonerLevel).Exec(); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Printf("Updated Summoner: %v", Success(summoner.Name))
	return &emptypb.Empty{}, nil
}

func (s *SummonerServer) DeleteSummoner(ctx context.Context, in *ScyllaSummonerReq) (*emptypb.Empty, error) {
	region := in.GetRegion()
	puuid := in.GetPuuid()

	if Session == nil {
		return nil, status.Error(codes.Internal, "Scylla Session is nil")
	}

	if puuid == "" {
		return nil, status.Error(codes.InvalidArgument, "Puuid is empty")
	}

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	log.Printf(Info("Deleting summoner"))
	filter := `DELETE FROM ` + region + `_lol.summoner WHERE puuid = ?`
	err := Session.Query(filter, puuid).Exec()
	if err != nil {
		if err == gocql.ErrNotFound {
			log.Printf("Summoner not found: %v", Success(puuid))
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Printf(Success("Deleted Summoner"))
	return &emptypb.Empty{}, nil
}
