package league

import (
	"context"
	"log"

	"github.com/gocql/gocql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	. "meep.gg/log"
	. "meep.gg/protos/scylla/lol/league/v1"
	. "meep.gg/scylla"
	. "meep.gg/types/lol"
)

type LeagueServer struct {
	ScyllaLeagueServiceServer
}

func (s *LeagueServer) GetLeague(ctx context.Context, in *ScyllaLeagueReq) (*ScyllaLeagueEntry, error) {
	region := in.GetRegion()
	summonerId := in.GetSummonerId()
	queueType := in.GetQueueType()

	if Session == nil {
		return nil, status.Error(codes.Internal, "Scylla Session is nil")
	}

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}
	if summonerId == "" {
		return nil, status.Error(codes.InvalidArgument, "SummonerId is empty")
	}

	if !IsQueue(queueType) {
		return nil, status.Error(codes.InvalidArgument, "QueueType is not of type Queue")
	}

	log.Printf(Info("Getting leagues"))
	var league = &DBLeague{}
	filter := `SELECT * FROM ` + region + `_lol.league WHERE summonerId = ? AND queueType = ? LIMIT 1`
	err := Session.Query(filter, summonerId, queueType).Scan(
		&league.SummonerId,
		&league.QueueType,
		&league.Current,
		&league.LeaguePoints,
		&league.Losses,
		&league.Rank,
		&league.Tier,
		&league.Wins,
	)
	if err != nil {
		if err == gocql.ErrNotFound {
			log.Printf("League not found: %v", Warn(summonerId))
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Printf("Got League: %v", Success(summonerId))

	return DeserializeLeague(league), nil
}

func (s *LeagueServer) UpsertLeague(ctx context.Context, in *ScyllaUpsertLeagueReq) (*emptypb.Empty, error) {
	region := in.GetRegion()
	league := in.GetLeague()

	if Session == nil {
		return nil, status.Error(codes.Internal, "Scylla Session is nil")
	}
	if league == nil {
		return nil, status.Error(codes.InvalidArgument, "League is nil")
	}
	if league.SummonerId == "" {
		return nil, status.Error(codes.InvalidArgument, "SummonerId is empty")
	}

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	log.Printf(Info("Upserting league"))
	filter := `INSERT INTO ` + region + `_lol.league (
    summonerId,
		queueType,
		current,
    leaguePoints,
    losses,
    rank,
    tier,
    wins
) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	err := Session.Query(filter,
		league.SummonerId,
		league.QueueType,
		league.Current,
		league.LeaguePoints,
		league.Losses,
		league.Rank,
		league.Tier,
		league.Wins,
	).Exec()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Printf("Updated League: %v", Success(league.SummonerId))
	return &emptypb.Empty{}, nil
}

func (s *LeagueServer) GetLeagues(ctx context.Context, in *ScyllaLeagueReq) (*ScyllaLeagueEntries, error) {
	region := in.GetRegion()
	summonerId := in.GetSummonerId()

	if Session == nil {
		return nil, status.Error(codes.Internal, "Scylla Session is nil")
	}

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	if summonerId == "" {
		return nil, status.Error(codes.InvalidArgument, "SummonerId is empty")
	}

	log.Printf(Info("Getting leagues"))
	var leagues = []*ScyllaLeagueEntry{}
	filter := `SELECT * FROM ` + region + `_lol.league WHERE summonerId = ?`
	iter := Session.Query(filter, summonerId).Iter()
	for {
		var league = &ScyllaLeagueEntry{}
		if !iter.Scan(
			&league.SummonerId,
			&league.QueueType,
			&league.Current,
			&league.LeaguePoints,
			&league.Losses,
			&league.Rank,
			&league.Tier,
			&league.Wins,
		) {
			break
		}
		leagues = append(leagues, league)
	}

	log.Printf("Got Leagues: %v", Success(summonerId))
	return &ScyllaLeagueEntries{
		Entries: leagues,
	}, nil
}

func (s *LeagueServer) DeleteLeague(ctx context.Context, in *ScyllaLeagueReq) (*emptypb.Empty, error) {
	region := in.GetRegion()
	summonerId := in.GetSummonerId()
	queueType := in.GetQueueType()

	if Session == nil {
		return nil, status.Error(codes.Internal, "Scylla Session is nil")
	}

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	if summonerId == "" {
		return nil, status.Error(codes.InvalidArgument, "SummonerId is empty")
	}

	if !IsQueue(queueType) {
		return nil, status.Error(codes.InvalidArgument, "QueueType is not of type Queue")
	}

	log.Printf(Info("Deleting league"))
	filter := `DELETE FROM ` + region + `_lol.league WHERE summonerId = ? AND queueType = ?`
	err := Session.Query(filter, summonerId, queueType).Exec()
	if err != nil {
		if err == gocql.ErrNotFound {
			log.Printf("League not found: %v", Warn(summonerId))
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Printf("Deleted League: %v", Success(summonerId))
	return &emptypb.Empty{}, nil
}
