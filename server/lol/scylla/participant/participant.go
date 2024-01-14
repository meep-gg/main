package participant

import (
	"context"
	"fmt"
	"log"

	"github.com/gocql/gocql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	. "meep.gg/log"
	. "meep.gg/protos/scylla/lol/participant/v1"
	. "meep.gg/scylla"
	. "meep.gg/types/lol"
)

type ParticipantServer struct {
	ScyllaParticipantServiceServer
}

func (s *ParticipantServer) GetMatchParticipants(ctx context.Context, req *ScyllaMatchParticipantsReq) (*ScyllaParticipants, error) {
	region := req.GetRegion()
	gameId := req.GetGameId()
	puuids := req.GetPuuids()

	if Session == nil {
		return nil, status.Error(codes.Internal, "Scylla Session is nil")
	}

	if gameId == 0 {
		return nil, status.Error(codes.InvalidArgument, "GameId is empty")
	}

	if len(puuids) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Puuids is empty")
	}

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	log.Printf(Info("Getting participants"))
	var participants = []*ScyllaParticipant{}
	filter := fmt.Sprintf(`SELECT * FROM %s_lol.participant WHERE gameId = ? AND puuid IN ?`, region)
	iter := Session.Query(filter, gameId, puuids).Iter()
	for {
		var participant DBParticipant
		if !iter.Scan(
			&participant.Puuid,
			&participant.GameId,
			&participant.Assists,
			&participant.ChampionId,
			&participant.ChampLevel,
			&participant.DamagePerMinute,
			&participant.Deaths,
			&participant.GoldEarned,
			&participant.GoldPerMinute,
			&participant.IndividualPosition,
			&participant.Item0,
			&participant.Item1,
			&participant.Item2,
			&participant.Item3,
			&participant.Item4,
			&participant.Item5,
			&participant.Item6,
			&participant.Kda,
			&participant.KillParticipation,
			&participant.Kills,
			&participant.NeutralMinionsKilled,
			&participant.ParticipantId,
			&participant.Placement,
			&participant.PlayerAugment1,
			&participant.PlayerAugment2,
			&participant.PlayerAugment3,
			&participant.PlayerAugment4,
			&participant.PlayerSubteamId,
			&participant.PrimaryStyle,
			&participant.Rank,
			&participant.SubStyle,
			&participant.Summoner1Id,
			&participant.Summoner2Id,
			&participant.TeamDamagePercentage,
			&participant.TeamId,
			&participant.TimeCCingOthers,
			&participant.TotalDamageDealtToChampions,
			&participant.TotalDamageTaken,
			&participant.TotalMinionsKilled,
			&participant.VisionScorePerMinute,
			&participant.Win,
		) {
			break
		}
		participants = append(participants, DeserializeParticipant(&participant))
	}
	if err := iter.Close(); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Printf("Got Participants: %v", Success(gameId))
	return &ScyllaParticipants{Participants: participants}, nil
}

func (s *ParticipantServer) PaginateParticipant(ctx context.Context, req *ScyllaPaginateParticipantsReq) (*ScyllaParticipants, error) {
	region := req.GetRegion()
	puuid := req.GetPuuid()
	gameId := req.GetGameId()
	count := req.GetCount()

	if Session == nil {
		return nil, status.Error(codes.Internal, "Scylla Session is nil")
	}

	if puuid == "" {
		return nil, status.Error(codes.InvalidArgument, "Puuid is empty")
	}

	if count < 0 {
		return nil, status.Error(codes.InvalidArgument, "Count is less than 0")
	}

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	log.Printf(Info("Getting participants"))
	var participants = []*ScyllaParticipant{}
	var filter string
	if gameId == 0 {
		filter = fmt.Sprintf(`SELECT * FROM %s_lol.participant WHERE puuid = ? LIMIT %d`, region, count)
	} else {
		filter = fmt.Sprintf(`SELECT * FROM %s_lol.participant WHERE puuid = ? AND gameId < %d LIMIT %d`, region, gameId, count)
	}

	iter := Session.Query(filter, puuid).Iter()
	for {
		var participant DBParticipant
		if !iter.Scan(
			&participant.Puuid,
			&participant.GameId,
			&participant.Assists,
			&participant.ChampionId,
			&participant.ChampLevel,
			&participant.DamagePerMinute,
			&participant.Deaths,
			&participant.GoldEarned,
			&participant.GoldPerMinute,
			&participant.IndividualPosition,
			&participant.Item0,
			&participant.Item1,
			&participant.Item2,
			&participant.Item3,
			&participant.Item4,
			&participant.Item5,
			&participant.Item6,
			&participant.Kda,
			&participant.KillParticipation,
			&participant.Kills,
			&participant.NeutralMinionsKilled,
			&participant.ParticipantId,
			&participant.Placement,
			&participant.PlayerAugment1,
			&participant.PlayerAugment2,
			&participant.PlayerAugment3,
			&participant.PlayerAugment4,
			&participant.PlayerSubteamId,
			&participant.PrimaryStyle,
			&participant.Rank,
			&participant.SubStyle,
			&participant.Summoner1Id,
			&participant.Summoner2Id,
			&participant.TeamDamagePercentage,
			&participant.TeamId,
			&participant.TimeCCingOthers,
			&participant.TotalDamageDealtToChampions,
			&participant.TotalDamageTaken,
			&participant.TotalMinionsKilled,
			&participant.VisionScorePerMinute,
			&participant.Win,
		) {
			break
		}
		participants = append(participants, DeserializeParticipant(&participant))
	}

	if err := iter.Close(); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Printf("Got Participants: %v", Success(puuid))
	return &ScyllaParticipants{Participants: participants}, nil
}

func (s *ParticipantServer) GetParticipant(ctx context.Context, in *ScyllaParticipantReq) (*ScyllaParticipant, error) {
	region := in.GetRegion()
	gameId := in.GetGameId()
	puuid := in.GetPuuid()

	if Session == nil {
		return nil, status.Error(codes.Internal, "Scylla Session is nil")
	}

	if gameId == 0 && puuid == "" {
		return nil, status.Error(codes.InvalidArgument, "GameId or Puuid is empty")
	}

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	log.Printf(Info("Getting participant"))
	var participant = DBParticipant{}
	filter := `SELECT * FROM ` + region + `_lol.participant WHERE puuid = ? AND gameId = ? LIMIT 1`
	err := Session.Query(filter, puuid, gameId).Scan(
		&participant.Puuid,
		&participant.GameId,
		&participant.Assists,
		&participant.ChampionId,
		&participant.ChampLevel,
		&participant.DamagePerMinute,
		&participant.Deaths,
		&participant.GoldEarned,
		&participant.GoldPerMinute,
		&participant.IndividualPosition,
		&participant.Item0,
		&participant.Item1,
		&participant.Item2,
		&participant.Item3,
		&participant.Item4,
		&participant.Item5,
		&participant.Item6,
		&participant.Kda,
		&participant.KillParticipation,
		&participant.Kills,
		&participant.NeutralMinionsKilled,
		&participant.ParticipantId,
		&participant.Placement,
		&participant.PlayerAugment1,
		&participant.PlayerAugment2,
		&participant.PlayerAugment3,
		&participant.PlayerAugment4,
		&participant.PlayerSubteamId,
		&participant.PrimaryStyle,
		&participant.Rank,
		&participant.SubStyle,
		&participant.Summoner1Id,
		&participant.Summoner2Id,
		&participant.TeamDamagePercentage,
		&participant.TeamId,
		&participant.TimeCCingOthers,
		&participant.TotalDamageDealtToChampions,
		&participant.TotalDamageTaken,
		&participant.TotalMinionsKilled,
		&participant.VisionScorePerMinute,
		&participant.Win,
	)
	if err != nil {
		if err == gocql.ErrNotFound {
			log.Printf("Participant not found: %v", Warn(puuid))
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Printf("Got Participant: %v", Success(participant.Puuid))
	return DeserializeParticipant(&participant), nil
}

func (s *ParticipantServer) GetParticipants(ctx context.Context, in *ScyllaGetParticipantsReq) (*ScyllaParticipants, error) {
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

	log.Printf(Info("Getting participants"))
	var participants = []*ScyllaParticipant{}
	filter := `SELECT * FROM ` + region + `_lol.participant WHERE puuid = ?`
	iter := Session.Query(filter, puuid).Iter()
	for {
		var participant DBParticipant
		if !iter.Scan(
			&participant.Puuid,
			&participant.GameId,
			&participant.Assists,
			&participant.ChampionId,
			&participant.ChampLevel,
			&participant.DamagePerMinute,
			&participant.Deaths,
			&participant.GoldEarned,
			&participant.GoldPerMinute,
			&participant.IndividualPosition,
			&participant.Item0,
			&participant.Item1,
			&participant.Item2,
			&participant.Item3,
			&participant.Item4,
			&participant.Item5,
			&participant.Item6,
			&participant.Kda,
			&participant.KillParticipation,
			&participant.Kills,
			&participant.NeutralMinionsKilled,
			&participant.ParticipantId,
			&participant.Placement,
			&participant.PlayerAugment1,
			&participant.PlayerAugment2,
			&participant.PlayerAugment3,
			&participant.PlayerAugment4,
			&participant.PlayerSubteamId,
			&participant.PrimaryStyle,
			&participant.Rank,
			&participant.SubStyle,
			&participant.Summoner1Id,
			&participant.Summoner2Id,
			&participant.TeamDamagePercentage,
			&participant.TeamId,
			&participant.TimeCCingOthers,
			&participant.TotalDamageDealtToChampions,
			&participant.TotalDamageTaken,
			&participant.TotalMinionsKilled,
			&participant.VisionScorePerMinute,
			&participant.Win,
		) {
			break
		}
		participants = append(participants, DeserializeParticipant(&participant))
	}
	if err := iter.Close(); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Printf("Got Participants: %v", Success(puuid))
	return &ScyllaParticipants{Participants: participants}, nil
}

func (s *ParticipantServer) UpsertParticipant(ctx context.Context, in *ScyllaUpsertParticipantReq) (*emptypb.Empty, error) {
	region := in.GetRegion()
	participant := in.GetParticipant()

	if participant.GameId == 0 && participant.Puuid == "" {
		return nil, status.Error(codes.InvalidArgument, "GameId or Puuid is empty")
	}

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	log.Printf(Info("Updating participant"))

	filter := `INSERT INTO ` + region + `_lol.participant
	(
		puuid,
		gameId,
		assists,
		championId,
		champLevel,
		damagePerMinute,
		deaths,
		goldEarned,
		goldPerMinute,
		individualPosition,
		item0,
		item1,
		item2,
		item3,
		item4,
		item5,
		item6,
		kda,
		killParticipation,
		kills,
		neutralMinionsKilled,
		participantId,
		placement,
		playerAugment1,
		playerAugment2,
		playerAugment3,
		playerAugment4,
		playerSubteamId,
		primaryStyle,
		rank,
		subStyle,
		summoner1Id,
		summoner2Id,
		teamDamagePercentage,
		teamId,
		timeCCingOthers,
		totalDamageDealtToChampions,
		totalDamageTaken,
		totalMinionsKilled,
		visionScorePerMinute,
		win
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, 
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	if err := Session.Query(filter,
		participant.Puuid,
		participant.GameId,
		participant.Assists,
		participant.ChampionId,
		participant.ChampLevel,
		participant.DamagePerMinute,
		participant.Deaths,
		participant.GoldEarned,
		participant.GoldPerMinute,
		participant.IndividualPosition,
		participant.Item0,
		participant.Item1,
		participant.Item2,
		participant.Item3,
		participant.Item4,
		participant.Item5,
		participant.Item6,
		participant.Kda,
		participant.KillParticipation,
		participant.Kills,
		participant.NeutralMinionsKilled,
		participant.ParticipantId,
		participant.Placement,
		participant.PlayerAugment1,
		participant.PlayerAugment2,
		participant.PlayerAugment3,
		participant.PlayerAugment4,
		participant.PlayerSubteamId,
		participant.PrimaryStyle,
		participant.Rank,
		participant.SubStyle,
		participant.Summoner1Id,
		participant.Summoner2Id,
		participant.TeamDamagePercentage,
		participant.TeamId,
		participant.TimeCCingOthers,
		participant.TotalDamageDealtToChampions,
		participant.TotalDamageTaken,
		participant.TotalMinionsKilled,
		participant.VisionScorePerMinute,
		participant.Win,
	).Exec(); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Printf("Updated Participant: %v", Success(participant.Puuid))
	return &emptypb.Empty{}, nil
}

func (s *ParticipantServer) DeleteParticipant(ctx context.Context, in *ScyllaParticipantReq) (*emptypb.Empty, error) {
	region := in.GetRegion()
	gameId := in.GetGameId()
	puuid := in.GetPuuid()

	if Session == nil {
		return nil, status.Error(codes.Internal, "Scylla Session is nil")
	}

	if gameId == 0 && puuid == "" {
		return nil, status.Error(codes.InvalidArgument, "GameId or Puuid is empty")
	}

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	log.Printf(Info("Deleting participant"))
	filter := `DELETE FROM ` + region + `_lol.participant WHERE gameId = ? AND puuid = ?`
	err := Session.Query(filter, gameId, puuid).Exec()
	if err != nil {
		if err == gocql.ErrNotFound {
			log.Printf("Participant not found: %v", Warn(puuid))
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Printf("Deleted Participant: %v", Success(puuid))
	return &emptypb.Empty{}, nil
}
