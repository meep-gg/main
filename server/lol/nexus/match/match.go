package match

import (
	"context"

	// . "meep.gg/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	. "meep.gg/protos/nexus/lol/match/v1"

	. "meep.gg/protos/scylla/lol/match/v1"
	. "meep.gg/protos/scylla/lol/participant/v1"
	. "meep.gg/protos/scylla/lol/participantdetail/v1"

	. "meep.gg/lol-nexus/connect"
	. "meep.gg/protos/gateway/lol/match/v1"
	. "meep.gg/types/lol"
)

type MatchServer struct {
	MatchServiceServer
}

func (s *MatchServer) GetMatchDetails(ctx context.Context, req *GetMatchReq) (*MatchDetailsRes, error) {
	region := req.GetRegion()
	gameId := req.GetGameId()

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	if gameId == 0 {
		return nil, status.Error(codes.InvalidArgument, "GameId is empty")
	}

	// check if match exists
	match, err := ScyllaMatchClient.GetMatch(ctx, &ScyllaMatchReq{
		Region: region,
		GameId: gameId,
	})
	if err != nil {
		st := status.Convert(err)
		if st.Code() != codes.NotFound {
			return nil, status.Error(codes.Internal, st.Message())
		}
		return nil, status.Error(codes.NotFound, "Match not found")
	}

	// Grab match participants
	data, err := ScyllaParticipantdetailClient.GetMatchParticipantsdetail(ctx, &ScyllaGetParticipantdetailsReq{
		Region: region,
		GameId: gameId,
		Count:  int32(len(match.Participants)),
	})
	if err != nil {
		st := status.Convert(err)
		if st.Code() != codes.NotFound {
			return nil, status.Error(codes.Internal, st.Message())
		}
		return nil, status.Error(codes.NotFound, "Match participants not found")
	}

	return &MatchDetailsRes{
		Details: data.Details,
	}, nil
}

func (s *MatchServer) GetMatchParticipants(ctx context.Context, req *GetMatchReq) (*MatchParticipantsRes, error) {
	region := req.GetRegion()
	gameId := req.GetGameId()

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	if gameId == 0 {
		return nil, status.Error(codes.InvalidArgument, "GameId is empty")
	}

	// check if match exists
	match, err := ScyllaMatchClient.GetMatch(ctx, &ScyllaMatchReq{
		Region: region,
		GameId: gameId,
	})
	if err != nil {
		st := status.Convert(err)
		if st.Code() != codes.NotFound {
			return nil, status.Error(codes.Internal, st.Message())
		}
		return nil, status.Error(codes.NotFound, "Match not found")
	}

	var puuids []string
	for _, participant := range match.Participants {
		puuids = append(puuids, participant.Puuid)
	}
	// Grab match participants
	data, err := ScyllaParticipantClient.GetMatchParticipants(ctx, &ScyllaMatchParticipantsReq{
		Region: region,
		GameId: gameId,
		Puuids: puuids,
	})
	if err != nil {
		st := status.Convert(err)
		if st.Code() != codes.NotFound {
			return nil, status.Error(codes.Internal, st.Message())
		}
		return nil, status.Error(codes.NotFound, "Match participants not found")
	}

	return &MatchParticipantsRes{
		Participants: data.Participants,
	}, nil
}

func (s *MatchServer) GetMatches(ctx context.Context, req *GetMatchesReq) (*MatchesRes, error) {
	region := req.GetRegion()
	puuid := req.GetPuuid()
	gameId := req.GetGameId()

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	if puuid == "" {
		return nil, status.Error(codes.InvalidArgument, "Puuid is empty")
	}

	var filter = &ScyllaPaginateParticipantsReq{
		Region: region,
		Puuid:  puuid,
		Count:  10,
	}
	if gameId != 0 {
		filter.GameId = gameId
	}
	// Grab players matches
	data, err := ScyllaParticipantClient.PaginateParticipant(ctx, filter)
	if err != nil {
		st := status.Convert(err)
		if st.Code() != codes.NotFound {
			return nil, status.Error(codes.Internal, st.Message())
		}
		return nil, status.Error(codes.NotFound, "Participant not found")
	}
	if len(data.Participants) == 0 {
		return nil, status.Error(codes.NotFound, "Participant not found")
	}

	var matches []*ScyllaMatch
	for _, participant := range data.Participants {
		// Grab matches
		match, err := ScyllaMatchClient.GetMatch(ctx, &ScyllaMatchReq{
			Region: region,
			GameId: participant.GameId,
		})
		if err != nil {
			st := status.Convert(err)
			if st.Code() != codes.NotFound {
				return nil, status.Error(codes.Internal, st.Message())
			}
			return nil, status.Error(codes.NotFound, "Match not found")
		}
		if match == nil {
			return nil, status.Error(codes.NotFound, "Match not found")
		}
		matches = append(matches, match)
	}

	return &MatchesRes{
		Participants: data.Participants,
		Matches:      matches,
	}, nil
}

func (s *MatchServer) UpdatePlayerMatches(ctx context.Context, req *UpdatePlayerMatchesReq) (*emptypb.Empty, error) {
	region := req.GetRegion()
	puuid := req.GetPuuid()

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	if puuid == "" {
		return nil, status.Error(codes.InvalidArgument, "Puuid is empty")
	}

	// Update player matches
	_, err := GatewayMatchClient.UpdatePlayerMatches(ctx, &GatewayMatchPlayerReq{
		Region: region,
		Puuid:  puuid,
	})
	if err != nil {
		st := status.Convert(err)
		if st.Code() != codes.NotFound {
			return nil, status.Error(codes.Internal, st.Message())
		}
		return nil, status.Error(codes.NotFound, st.Message())
	}

	return &emptypb.Empty{}, nil
}
