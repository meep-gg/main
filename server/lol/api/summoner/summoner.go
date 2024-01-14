package lol

import (
	"context"
	"encoding/json"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	. "meep.gg/http/riot"
	. "meep.gg/log"
	. "meep.gg/protos/riot/lol/summoner/v1"
	. "meep.gg/rates"
	. "meep.gg/types/lol"
)

type SummonerServer struct {
	RiotSummonerServiceServer
}

func (s *SummonerServer) Account(ctx context.Context, in *RiotSummonerAccountReq) (*RiotSummoner, error) {
	region := in.GetRegion()
	accountId := in.GetAccountId()

	if region == "" || accountId == "" {
		return nil, status.Error(codes.InvalidArgument, "Region or AccountId is empty")
	}

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}
	log.Print("Checking API Rate Limits")
	err := CheckRates(ctx, region, SummonerAccountRates)
	if err != nil {
		return nil, status.Error(codes.ResourceExhausted, err.Error())
	}
	log.Printf("Lol API Server Processing task Account: \n%v", Info(in))

	body, err := Get(ctx, region, "/lol/summoner/v4/summoners/by-account/"+accountId)
	if err != nil {
		st := status.Convert(err)
		return nil, st.Err()
	}

	var summoner *RiotSummoner
	err = json.Unmarshal(body, &summoner)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Printf("Retrieved Summoner: %v", Success(summoner.Name))
	return summoner, nil
}

func (s *SummonerServer) Name(ctx context.Context, in *RiotSummonerNameReq) (*RiotSummoner, error) {
	region := in.GetRegion()
	name := in.GetName()

	if region == "" || name == "" {
		return nil, status.Error(codes.InvalidArgument, "Region or Name is empty")
	}

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	log.Print("Checking API Rate Limits")
	err := CheckRates(ctx, region, SummonerNameRates)
	if err != nil {
		return nil, status.Error(codes.ResourceExhausted, err.Error())
	}
	log.Printf("Lol API Server Processing task Name: \n%v", Info(in))

	body, err := Get(ctx, region, "/lol/summoner/v4/summoners/by-name/"+name)
	if err != nil {
		st := status.Convert(err)
		return nil, st.Err()
	}

	var summoner *RiotSummoner
	err = json.Unmarshal(body, &summoner)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Printf("Retrieved Summoner: %v", Success(summoner.Name))
	return summoner, nil
}

func (s *SummonerServer) Puuid(ctx context.Context, in *RiotSummonerPuuidReq) (*RiotSummoner, error) {
	region := in.GetRegion()
	puuid := in.GetPuuid()

	if region == "" || puuid == "" {
		return nil, status.Error(codes.InvalidArgument, "Region or Puuid is empty")
	}

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	log.Print("Checking API Rate Limits")
	err := CheckRates(ctx, region, SummonerPuuidRates)
	if err != nil {
		return nil, status.Error(codes.ResourceExhausted, err.Error())
	}
	log.Printf("Lol API Server Processing task Puuid: \n%v", Info(in))

	body, err := Get(ctx, region, "/lol/summoner/v4/summoners/by-puuid/"+puuid)
	if err != nil {
		st := status.Convert(err)
		return nil, st.Err()
	}

	var summoner *RiotSummoner
	err = json.Unmarshal(body, &summoner)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Printf("Retrieved Summoner: %v", Success(summoner.Name))
	return summoner, nil
}

func (s *SummonerServer) Id(ctx context.Context, in *RiotSummonerIdReq) (*RiotSummoner, error) {
	region := in.GetRegion()
	id := in.GetId()

	if region == "" || id == "" {
		return nil, status.Error(codes.InvalidArgument, "Region or Id is empty")
	}

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	log.Print("Checking API Rate Limits")
	err := CheckRates(ctx, region, SummonerIdRates)
	if err != nil {
		return nil, status.Error(codes.ResourceExhausted, err.Error())
	}
	log.Printf("Lol API Server Processing task Id: \n%v", Info(in))
	body, err := Get(ctx, region, "/lol/summoner/v4/summoners/"+id)
	if err != nil {
		st := status.Convert(err)
		return nil, st.Err()
	}

	var summoner *RiotSummoner
	err = json.Unmarshal(body, &summoner)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Printf("Retrieved Summoner: %v", Success(summoner.Name))
	return summoner, nil
}
