package profile

import (
	"context"
	"log"

	"github.com/gocql/gocql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	. "meep.gg/log"
	. "meep.gg/protos/scylla/lol/profile/v1"
	. "meep.gg/scylla"
	. "meep.gg/types/lol"
)

type ProfileServer struct {
	ScyllaProfileServiceServer
}

func (s *ProfileServer) GetProfile(ctx context.Context, req *ScyllaProfileReq) (*ScyllaProfile, error) {
	region := req.GetRegion()
	gameName := req.GetGameName()
	tagLine := req.GetTagLine()

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	if gameName == "" || tagLine == "" {
		return nil, status.Error(codes.InvalidArgument, "GameName or TagLine is empty")
	}

	log.Printf("Getting profile")
	var profile = DBProfile{}
	filter := `SELECT * FROM ` + region + `_lol.profile WHERE gameName = ? AND tagLine = ? LIMIT 1`
	err := Session.Query(filter, gameName, tagLine).Scan(
		&profile.GameName,
		&profile.TagLine,
		&profile.ProfileIconId,
		&profile.Puuid,
		&profile.SummonerId,
		&profile.SummonerLevel,
	)
	if err != nil {
		if err == gocql.ErrNotFound {
			log.Printf("Profile not found: %v", Warn(gameName))
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return DeserializeProfile(&profile), nil
}

func (s *ProfileServer) UpsertProfile(ctx context.Context, req *ScyllaUpsertProfileReq) (*ScyllaProfile, error) {
	region := req.GetRegion()
	profile := req.GetProfile()

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	if profile == nil {
		return nil, status.Error(codes.InvalidArgument, "Profile is nil")
	}

	log.Printf("Upserting profile")
	filter := `INSERT INTO ` + region + `_lol.profile (
		gameName, 
		tagLine, 
		puuid, 
		summonerId, 
		profileIconId, 
		summonerLevel
		) VALUES (?, ?, ?, ?, ?, ?)`
	err := Session.Query(filter,
		profile.GameName,
		profile.TagLine,
		profile.Puuid,
		profile.SummonerId,
		profile.ProfileIconId,
		profile.SummonerLevel,
	).Exec()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return profile, nil
}

func (s *ProfileServer) DeleteProfile(ctx context.Context, req *ScyllaProfileReq) (*emptypb.Empty, error) {
	region := req.GetRegion()
	gameName := req.GetGameName()
	tagLine := req.GetTagLine()

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	if gameName == "" || tagLine == "" {
		return nil, status.Error(codes.InvalidArgument, "GameName or TagLine is empty")
	}

	log.Printf("Deleting profile")
	filter := `DELETE FROM ` + region + `_lol.profile WHERE gameName = ? AND tagLine = ?`
	err := Session.Query(filter, gameName, tagLine).Exec()
	if err != nil {
		if err == gocql.ErrNotFound {
			log.Printf("Profile not found: %v", Warn(gameName))
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
