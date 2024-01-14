package lol

import (
	"context"
	"encoding/json"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	. "meep.gg/http/riot"
	. "meep.gg/log"
	. "meep.gg/protos/riot/lol/match/v1"
	. "meep.gg/rates"
	. "meep.gg/types/lol"
)

type MatchServer struct {
	RiotMatchServiceServer
}

func (s *MatchServer) Puuid(ctx context.Context, in *RiotMatchPuuidReq) (*RiotMatchIds, error) {
	region := in.GetRegion()
	puuid := in.GetPuuid()

	startTime := in.GetStartTime()
	endTime := in.GetEndTime()
	queue := in.GetQueue()
	type_ := in.GetType()
	start := in.GetStart()
	count := in.GetCount()

	if region == "" || puuid == "" {
		return nil, status.Error(codes.InvalidArgument, "Region or Puuid is empty")
	}

	if !IsRegional(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Regional")
	}

	log.Printf("Checking API Rate Limits")
	err := CheckRates(ctx, region, MatchPuuidRates)
	if err != nil {
		return nil, status.Error(codes.ResourceExhausted, err.Error())
	}
	log.Printf("Lol API Server Processing task Match: %v", Info(in))

	if start < 0 || count < 0 {
		return nil, status.Error(codes.InvalidArgument, "Start or Count is less than 0")
	}

	filter := make(map[string]interface{})
	if startTime != 0 && endTime != 0 {
		if startTime > endTime {
			return nil, status.Error(codes.InvalidArgument, "Start time is greater than end time")
		}
	}
	if startTime != 0 {
		filter["startTime"] = startTime
	}

	if endTime != 0 {
		filter["endTime"] = endTime
	}

	if queue != 0 {
		filter["queue"] = queue
	}

	if type_ != "" {
		filter["type"] = type_
	}

	if start != 0 {
		filter["start"] = start
	}

	if count != 0 {
		if count > 100 {
			count = 100
		}
		filter["count"] = count
	}
	body, err := Get(ctx, region, "/lol/match/v5/matches/by-puuid/"+puuid+"/ids", filter)
	if err != nil {
		st := status.Convert(err)
		return nil, st.Err()
	}

	var matchIds []string
	err = json.Unmarshal(body, &matchIds)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &RiotMatchIds{Matches: matchIds}, nil
}

func (s *MatchServer) Id(ctx context.Context, in *RiotMatchIdReq) (*RiotMatch, error) {
	region := in.GetRegion()
	matchId := in.GetMatchId()

	if region == "" || matchId == "" {
		return nil, status.Error(codes.InvalidArgument, "Region or MatchId is empty")
	}

	if !IsRegional(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Regional")
	}

	log.Printf("Checking API Rate Limits")
	err := CheckRates(ctx, region, MatchIdRates)
	if err != nil {
		return nil, status.Error(codes.ResourceExhausted, err.Error())
	}

	log.Printf("Lol API Server Processing task Match: %v", Info(in))
	body, err := Get(ctx, region, "/lol/match/v5/matches/"+matchId)
	if err != nil {
		st := status.Convert(err)
		return nil, st.Err()
	}

	var match RiotMatch
	err = json.Unmarshal(body, &match)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &match, nil
}
