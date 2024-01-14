package lol

import (
	"context"
	"encoding/json"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	. "meep.gg/http/riot"
	. "meep.gg/log"
	. "meep.gg/protos/riot/lol/league/v1"
	. "meep.gg/rates"
	. "meep.gg/types/lol"
)

type LeagueServer struct {
	RiotLeagueServiceServer
}

func (s *LeagueServer) Challenger(ctx context.Context, in *RiotLeagueApexReq) (*RiotLeagueList, error) {
	region := in.GetRegion()
	queue := in.GetQueue()

	if region == "" || queue == "" {
		return nil, status.Error(codes.InvalidArgument, "Region or Queue is empty")
	}

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	if !IsQueue(queue) {
		return nil, status.Error(codes.InvalidArgument, "Queue is not of type Queue")
	}

	err := CheckRates(ctx, region, LeagueChallengerRates)
	if err != nil {
		return nil, status.Error(codes.ResourceExhausted, err.Error())
	}

	log.Printf("Lol API Server Processing task Challenger: %v", Info(in))
	body, err := Get(ctx, region, "/lol/league/v4/challengerleagues/by-queue/"+queue)
	if err != nil {
		st := status.Convert(err)
		return nil, st.Err()
	}

	log.Printf("Unmarshalling data")
	var leagueList *RiotLeagueList
	err = json.Unmarshal(body, &leagueList)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Printf("Retrieved Challenger Entries: %v", Success(len(leagueList.Entries)))
	return leagueList, nil
}

func (s *LeagueServer) Grandmaster(ctx context.Context, in *RiotLeagueApexReq) (*RiotLeagueList, error) {
	region := in.GetRegion()
	queue := in.GetQueue()

	if region == "" || queue == "" {
		return nil, status.Error(codes.InvalidArgument, "Region or Queue is empty")
	}

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	if !IsQueue(queue) {
		return nil, status.Error(codes.InvalidArgument, "Queue is not of type Queue")
	}

	log.Printf("Checking API Rate Limits")
	err := CheckRates(ctx, region, LeagueGrandmasterRates)
	if err != nil {
		return nil, status.Error(codes.ResourceExhausted, err.Error())
	}

	log.Printf("Lol API Server Processing task Grandmaster: %v", Info(in))
	body, err := Get(ctx, region, "/lol/league/v4/grandmasterleagues/by-queue/"+queue)
	if err != nil {
		st := status.Convert(err)
		return nil, st.Err()
	}

	log.Printf("Unmarshalling data")
	var leagueList *RiotLeagueList
	err = json.Unmarshal(body, &leagueList)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Printf("Retrieved Grandmaster Entries: %v", Success(len(leagueList.Entries)))
	return leagueList, nil
}

func (s *LeagueServer) Master(ctx context.Context, in *RiotLeagueApexReq) (*RiotLeagueList, error) {
	region := in.GetRegion()
	queue := in.GetQueue()

	if region == "" || queue == "" {
		return nil, status.Error(codes.InvalidArgument, "Region or Queue is empty")
	}

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	if !IsQueue(queue) {
		return nil, status.Error(codes.InvalidArgument, "Queue is not of type Queue")
	}

	log.Printf("Checking API Rate Limits")
	err := CheckRates(ctx, region, LeagueMasterRates)
	if err != nil {
		return nil, status.Error(codes.ResourceExhausted, err.Error())
	}

	log.Printf("Lol API Server Processing task Master: %v", Info(in))
	body, err := Get(ctx, region, "/lol/league/v4/masterleagues/by-queue/"+queue)
	if err != nil {
		st := status.Convert(err)
		return nil, st.Err()
	}

	log.Printf("Unmarshalling data")
	var leagueList *RiotLeagueList
	err = json.Unmarshal(body, &leagueList)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Printf("Retrieved Master Entries: %v", Success(len(leagueList.Entries)))
	return leagueList, nil
}

func (s *LeagueServer) SummonerId(ctx context.Context, in *RiotLeagueSummonerIdReq) (*RiotLeagueEntries, error) {
	region := in.GetRegion()
	summonerId := in.GetSummonerId()

	if region == "" || summonerId == "" {
		return nil, status.Error(codes.InvalidArgument, "Region or SummonerId is empty")
	}

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	log.Printf("Checking API Rate Limits")
	err := CheckRates(ctx, region, LeagueSummonerIdRates)
	if err != nil {
		return nil, status.Error(codes.ResourceExhausted, err.Error())
	}

	log.Printf("Lol API Server Processing task SummonerId: %v", Info(in))
	body, err := Get(ctx, region, "/lol/league/v4/entries/by-summoner/"+summonerId)
	if err != nil {
		st := status.Convert(err)
		return nil, st.Err()
	}

	var leagueEntries []*RiotLeagueEntry
	err = json.Unmarshal(body, &leagueEntries)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Printf("Retrieved SummonerId Entries: %v", Success(len(leagueEntries)))
	return &RiotLeagueEntries{Leagues: leagueEntries}, nil
}

func (s *LeagueServer) LeagueId(ctx context.Context, in *RiotLeagueIdReq) (*RiotLeagueList, error) {
	region := in.GetRegion()
	leagueId := in.GetLeagueId()

	if region == "" || leagueId == "" {
		return nil, status.Error(codes.InvalidArgument, "Region or LeagueId is empty")
	}

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	log.Printf("Checking API Rate Limits")
	err := CheckRates(ctx, region, LeagueIdRates)
	if err != nil {
		return nil, status.Error(codes.ResourceExhausted, err.Error())
	}

	log.Printf("Lol API Server Processing task LeagueId: %v", Info(in))
	body, err := Get(ctx, region, "/lol/league/v4/leagues/"+leagueId)
	if err != nil {
		st := status.Convert(err)
		return nil, st.Err()
	}

	var leagueList *RiotLeagueList
	err = json.Unmarshal(body, &leagueList)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Printf("Retrieved LeagueId Entries: %v", Success(len(leagueList.Entries)))
	return leagueList, nil
}

func (s *LeagueServer) Entry(ctx context.Context, in *RiotLeagueEntryReq) (*RiotLeagueEntries, error) {
	region := in.GetRegion()
	division := in.GetDivision()
	tier := in.GetTier()
	queue := in.GetQueue()
	page := in.GetPage()

	if region == "" || division == "" || tier == "" || queue == "" {
		return nil, status.Error(codes.InvalidArgument, "Region, Division, Tier, or Queue is empty")
	}

	if !IsTier(tier) {
		return nil, status.Error(codes.InvalidArgument, "Tier is not of type Tier")
	}

	if !IsDivision(division) {
		return nil, status.Error(codes.InvalidArgument, "Division is not of type Division")
	}

	if !IsQueue(queue) {
		return nil, status.Error(codes.InvalidArgument, "Queue is not of type Queue")
	}

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	log.Printf("Checking API Rate Limits")
	err := CheckRates(ctx, region, LeagueEntryRates)
	if err != nil {
		return nil, status.Error(codes.ResourceExhausted, err.Error())
	}

	if page == 0 {
		page = 1
	}
	filter := map[string]interface{}{
		"page": page,
	}
	log.Printf("Lol API Server Processing task Entry: %v", Info(in))
	body, err := Get(ctx, region, "/lol/league/v4/entries/"+queue+"/"+tier+"/"+division, filter)
	if err != nil {
		st := status.Convert(err)
		return nil, st.Err()
	}

	var leagueEntries []*RiotLeagueEntry
	err = json.Unmarshal(body, &leagueEntries)
	if err != nil {
		return nil, err
	}

	log.Printf("Retrieved Entries: %v", Success(len(leagueEntries)))
	return &RiotLeagueEntries{Leagues: leagueEntries}, nil
}
