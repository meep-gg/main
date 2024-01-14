package account

import (
	"context"
	"encoding/json"
	"log"

	http "meep.gg/http/riot"
	"meep.gg/rates"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	. "meep.gg/log"
	. "meep.gg/protos/riot/account/v1"
	. "meep.gg/types/lol"
)

type AccountServer struct {
	RiotAccountServiceServer
}

func (s *AccountServer) Puuid(ctx context.Context, in *RiotAccountReq) (*RiotAccount, error) {
	region := in.GetRegion()
	puuid := in.GetPuuid()

	if region == "" || puuid == "" {
		return nil, status.Error(codes.InvalidArgument, "Region or Puuid is empty")
	}

	if !IsRegional(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Regional")
	}

	err := rates.CheckRates(ctx, region, PuuidRates)
	if err != nil {
		return nil, status.Error(codes.ResourceExhausted, err.Error())
	}

	log.Printf("Riot API Server Processing task Puuid: %v", Info(puuid))
	data, err := http.Get(ctx, region, "/riot/account/v1/accounts/by-puuid/"+puuid)
	if err != nil {
		st := status.Convert(err)
		return nil, st.Err()
	}

	var account RiotAccount
	if err := json.Unmarshal(data, &account); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Printf("Retrieved Account: %v", Success(account.GameName+"#"+account.TagLine))
	return &account, nil
}

func (s *AccountServer) RiotId(ctx context.Context, in *RiotAccountReq) (*RiotAccount, error) {
	region := in.GetRegion()
	gameName := in.GetGameName()
	tagLine := in.GetTagLine()

	if region == "" || gameName == "" || tagLine == "" {
		return nil, status.Error(codes.InvalidArgument, "Region, GameName, or TagLine is empty")
	}

	if !IsRegional(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Regional")
	}

	err := rates.CheckRates(ctx, region, RiotIdRates)
	if err != nil {
		return nil, status.Error(codes.ResourceExhausted, err.Error())
	}

	log.Printf("Riot API Server Processing task RiotId: %v", Info(gameName+"#"+tagLine))
	data, err := http.Get(ctx, region, "/riot/account/v1/accounts/by-riot-id/"+gameName+"/"+tagLine)
	if err != nil {
		st := status.Convert(err)
		return nil, st.Err()
	}

	var account RiotAccount
	if err := json.Unmarshal(data, &account); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Printf("Retrieved Account: %v", Success(account.GameName+"#"+account.TagLine))
	return &account, nil
}
