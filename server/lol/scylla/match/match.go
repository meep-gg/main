package match

import (
	"context"
	"log"
	"sync"

	"github.com/gocql/gocql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	. "meep.gg/log"
	. "meep.gg/protos/scylla/lol/match/v1"
	. "meep.gg/scylla"
	. "meep.gg/types/lol"
)

type MatchServer struct {
	ScyllaMatchServiceServer
}

func (s *MatchServer) CheckMatch(ctx context.Context, in *ScyllaMatchReq) (*emptypb.Empty, error) {
	region := in.GetRegion()
	gameId := in.GetGameId()

	if Session == nil {
		return nil, status.Error(codes.Internal, "Scylla Session is nil")
	}

	if gameId == 0 {
		return nil, status.Error(codes.InvalidArgument, "GameId is empty")
	}

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	log.Printf(Info("Checking match"))
	var id string
	filter := `SELECT gameId FROM ` + region + `_lol.match WHERE gameId = ? LIMIT 1`
	err := Session.Query(filter, gameId).Scan(&id)
	if err != nil {
		if err == gocql.ErrNotFound {
			log.Printf("Match not found: %v", Success(gameId))
			return nil, status.Error(codes.NotFound, "Match not found")
		}
		log.Printf("error checking match: %v", Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Printf("Match found: %v", Success(gameId))
	return &emptypb.Empty{}, nil
}

func (s *MatchServer) GetMatch(ctx context.Context, in *ScyllaMatchReq) (*ScyllaMatch, error) {
	region := in.GetRegion()
	gameId := in.GetGameId()

	if Session == nil {
		return nil, status.Error(codes.Internal, "Scylla Session is nil")
	}

	if gameId == 0 {
		return nil, status.Error(codes.InvalidArgument, "GameId is empty")
	}

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	log.Printf(Info("Getting match"))
	var match = DBMatch{}
	filter := `SELECT * FROM ` + region + `_lol.match WHERE gameId = ? LIMIT 1`
	err := Session.Query(filter, gameId).Scan(
		&match.GameId,
		&match.Bans,
		&match.GameDuration,
		&match.GameEndTimestamp,
		&match.Participants,
		&match.Patch,
		&match.QueueId,
		&match.Rank,
		&match.Season,
	)
	if err != nil {
		if err == gocql.ErrNotFound {
			log.Printf("Match not found: %v", Success(gameId))
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Printf("Got Match: %v", Success(match.GameId))
	return DeserializeMatch(&match), nil
}

func (s *MatchServer) GetMatches(ctx context.Context, in *ScyllaGetMatchesReq) (*ScyllaMatches, error) {
	region := in.GetRegion()
	gameIds := in.GetGameIds()

	if Session == nil {
		return nil, status.Error(codes.Internal, "Scylla Session is nil")
	}

	if len(gameIds) == 0 {
		return nil, status.Error(codes.InvalidArgument, "GameIds is empty")
	}

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	log.Printf(Info("Getting matches"))
	results := make([]*ScyllaMatch, 0, len(gameIds))
	var mutex sync.Mutex
	var wg sync.WaitGroup

	for _, gameId := range gameIds {
		wg.Add(1)
		go func(gameId int64) {
			defer wg.Done()
			var match = DBMatch{}
			filter := `SELECT * FROM ` + region + `_lol.match WHERE gameId = ? LIMIT 1`
			err := Session.Query(filter, gameId).Scan(
				&match.GameId,
				&match.Bans,
				&match.GameDuration,
				&match.GameEndTimestamp,
				&match.Participants,
				&match.Patch,
				&match.QueueId,
				&match.Rank,
				&match.Season,
			)
			if err != nil {
				if err == gocql.ErrNotFound {
					log.Printf("Match not found: %v", Success(gameId))
					return
				}
				log.Printf("error getting match: %v", Error(err))
				return
			}

			mutex.Lock()
			results = append(results, DeserializeMatch(&match))
			mutex.Unlock()
		}(gameId)
	}

	wg.Wait()
	log.Printf("Got %v Matches", Success(len(results)))
	return &ScyllaMatches{Matches: results}, nil
}

func (s *MatchServer) UpsertMatch(ctx context.Context, in *ScyllaUpsertMatchReq) (*emptypb.Empty, error) {
	region := in.GetRegion()
	match := in.GetMatch()

	if match.GameId == 0 {
		return nil, status.Error(codes.InvalidArgument, "GameId is empty")
	}

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	log.Printf(Info("Updating match"))

	filter := `INSERT INTO ` + region + `_lol.match
	(
		gameId,
		bans,
		gameDuration,
		gameEndTimestamp,
		participants,
		patch,
		queueId,
		rank,
		season
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	err := Session.Query(filter,
		match.GameId,
		match.Bans,
		match.GameDuration,
		match.GameEndTimestamp,
		SerializeMParticipants(match.Participants),
		match.Patch,
		match.QueueId,
		match.Rank,
		match.Season,
	).Exec()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Printf("Updated Match: %v", Success(match.GameId))
	return &emptypb.Empty{}, nil
}

func (s *MatchServer) DeleteMatch(ctx context.Context, in *ScyllaMatchReq) (*emptypb.Empty, error) {
	region := in.GetRegion()
	gameId := in.GetGameId()

	if Session == nil {
		return nil, status.Error(codes.Internal, "Scylla Session is nil")
	}

	if gameId == 0 {
		return nil, status.Error(codes.InvalidArgument, "GameId is empty")
	}

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	log.Printf(Info("Deleting match"))
	filter := `DELETE FROM ` + region + `_lol.match WHERE gameId = ?`
	err := Session.Query(filter, gameId).Exec()
	if err != nil {
		if err == gocql.ErrNotFound {
			log.Printf("Match not found: %v", Success(gameId))
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Printf("Deleted Match: %v", Success(gameId))
	return &emptypb.Empty{}, nil
}
