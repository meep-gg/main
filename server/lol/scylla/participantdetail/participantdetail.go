package participantdetail

import (
	"context"
	"log"

	. "meep.gg/protos/scylla/lol/participantdetail/v1"

	"github.com/gocql/gocql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	. "meep.gg/log"
	. "meep.gg/scylla"
	. "meep.gg/types/lol"
)

type ParticipantdetailServer struct {
	ScyllaParticipantdetailServiceServer
}

func (s *ParticipantdetailServer) GetMatchParticipantsdetail(ctx context.Context, in *ScyllaGetParticipantdetailsReq) (*ScyllaParticipantdetails, error) {
	region := in.GetRegion()
	gameId := in.GetGameId()
	count := in.GetCount()

	if Session == nil {
		return nil, status.Error(codes.Internal, "Scylla Session is nil")
	}

	if gameId == 0 {
		return nil, status.Error(codes.InvalidArgument, "GameId is empty")
	}

	if count == 0 {
		return nil, status.Error(codes.InvalidArgument, "Count is empty")
	}

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	participantIds := make([]int32, count+1)
	for i := 0; i < int(count+1); i++ {
		participantIds[i] = int32(i)
	}
	log.Printf(Info("Getting participant details"))
	var participantdetails []*ScyllaParticipantdetail
	filter := `SELECT * FROM ` + region + `_lol.participant_detail WHERE gameId = ? AND participantId IN ?`
	iter := Session.Query(filter, gameId, participantIds).Iter()
	for {
		var participantdetail DBParticipantdetail
		if !iter.Scan(
			&participantdetail.GameId,
			&participantdetail.ParticipantId,
			&participantdetail.DamageDealtToBuildings,
			&participantdetail.DamageDealtToObjectives,
			&participantdetail.DamageDealtToTurrets,
			&participantdetail.DamageSelfMitigated,
			&participantdetail.DamageTakenOnTeamPercentage,
			&participantdetail.MagicDamageDealt,
			&participantdetail.MagicDamageDealtToChampions,
			&participantdetail.MagicDamageTaken,
			&participantdetail.MaxCsAdvantageOnLaneOpponent,
			&participantdetail.MaxLevelLeadLaneOpponent,
			&participantdetail.NeutralMinionsKilled,
			&participantdetail.Perks,
			&participantdetail.PhysicalDamageDealt,
			&participantdetail.PhysicalDamageDealtToChampions,
			&participantdetail.PhysicalDamageTaken,
			&participantdetail.Spell1Casts,
			&participantdetail.Spell2Casts,
			&participantdetail.Spell3Casts,
			&participantdetail.Spell4Casts,
			&participantdetail.Summoner1Casts,
			&participantdetail.Summoner2Casts,
		) {
			break
		}
		participantdetails = append(participantdetails, DeserializeParticipantdetail(&participantdetail))
	}
	if err := iter.Close(); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	log.Printf("Got Participant details: %v", Success(gameId))
	return &ScyllaParticipantdetails{
		Details: participantdetails,
	}, nil
}

func (s *ParticipantdetailServer) GetParticipantdetail(ctx context.Context, in *ScyllaParticipantdetailReq) (*ScyllaParticipantdetail, error) {
	region := in.GetRegion()
	gameId := in.GetGameId()
	participantId := in.GetParticipantId()

	if Session == nil {
		return nil, status.Error(codes.Internal, "Scylla Session is nil")
	}

	if gameId == 0 || participantId == 0 {
		return nil, status.Error(codes.InvalidArgument, "GameId or ParticipantId is empty")
	}

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	log.Printf(Info("Getting participant detail"))
	var participantdetail = DBParticipantdetail{}
	filter := `SELECT * FROM ` + region + `_lol.participant_detail WHERE gameId = ? AND participantId = ? LIMIT 1`
	err := Session.Query(filter, gameId, participantId).Scan(
		&participantdetail.GameId,
		&participantdetail.ParticipantId,
		&participantdetail.DamageDealtToBuildings,
		&participantdetail.DamageDealtToObjectives,
		&participantdetail.DamageDealtToTurrets,
		&participantdetail.DamageSelfMitigated,
		&participantdetail.DamageTakenOnTeamPercentage,
		&participantdetail.MagicDamageDealt,
		&participantdetail.MagicDamageDealtToChampions,
		&participantdetail.MagicDamageTaken,
		&participantdetail.MaxCsAdvantageOnLaneOpponent,
		&participantdetail.MaxLevelLeadLaneOpponent,
		&participantdetail.NeutralMinionsKilled,
		&participantdetail.Perks,
		&participantdetail.PhysicalDamageDealt,
		&participantdetail.PhysicalDamageDealtToChampions,
		&participantdetail.PhysicalDamageTaken,
		&participantdetail.Spell1Casts,
		&participantdetail.Spell2Casts,
		&participantdetail.Spell3Casts,
		&participantdetail.Spell4Casts,
		&participantdetail.Summoner1Casts,
		&participantdetail.Summoner2Casts,
	)
	if err != nil {
		if err == gocql.ErrNotFound {
			log.Printf("Participant detail not found: %v", Warn(gameId))
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Printf("Got Participant detail: %v", Success(participantdetail.GameId))
	return DeserializeParticipantdetail(&participantdetail), nil
}

func (s *ParticipantdetailServer) UpsertParticipantdetail(ctx context.Context, in *ScyllaUpsertParticipantdetailReq) (*emptypb.Empty, error) {
	region := in.GetRegion()
	participantdetail := in.GetDetail()

	if Session == nil {
		return nil, status.Error(codes.Internal, "Scylla Session is nil")
	}

	if participantdetail == nil {
		return nil, status.Error(codes.InvalidArgument, "Participantdetail is empty")
	}

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	log.Printf(Info("Upserting participant detail"))
	filter := `INSERT INTO ` + region + `_lol.participant_detail (
		gameId,
		participantId,
		damageDealtToBuildings,
		damageDealtToObjectives,
		damageDealtToTurrets,
		damageSelfMitigated,
		damageTakenOnTeamPercentage,
		magicDamageDealt,
		magicDamageDealtToChampions,
		magicDamageTaken,
		maxCsAdvantageOnLaneOpponent,
		maxLevelLeadLaneOpponent,
		neutralMinionsKilled,
		perks,
		physicalDamageDealt,
		physicalDamageDealtToChampions,
		physicalDamageTaken,
		spell1Casts,
		spell2Casts,
		spell3Casts,
		spell4Casts,
		summoner1Casts,
		summoner2Casts
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	err := Session.Query(filter,
		participantdetail.GameId,
		participantdetail.ParticipantId,
		participantdetail.DamageDealtToBuildings,
		participantdetail.DamageDealtToObjectives,
		participantdetail.DamageDealtToTurrets,
		participantdetail.DamageSelfMitigated,
		participantdetail.DamageTakenOnTeamPercentage,
		participantdetail.MagicDamageDealt,
		participantdetail.MagicDamageDealtToChampions,
		participantdetail.MagicDamageTaken,
		participantdetail.MaxCsAdvantageOnLaneOpponent,
		participantdetail.MaxLevelLeadLaneOpponent,
		participantdetail.NeutralMinionsKilled,
		SerializePerks(participantdetail.Perks),
		participantdetail.PhysicalDamageDealt,
		participantdetail.PhysicalDamageDealtToChampions,
		participantdetail.PhysicalDamageTaken,
		participantdetail.Spell1Casts,
		participantdetail.Spell2Casts,
		participantdetail.Spell3Casts,
		participantdetail.Spell4Casts,
		participantdetail.Summoner1Casts,
		participantdetail.Summoner2Casts,
	).Exec()
	if err != nil {
		if err == gocql.ErrNotFound {
			log.Printf("Participant detail not found: %v", Warn(participantdetail.GameId))
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Printf("Upserted Participant detail: %v", Success(participantdetail.GameId))
	return &emptypb.Empty{}, nil
}

func (s *ParticipantdetailServer) DeleteParticipantdetail(ctx context.Context, in *ScyllaParticipantdetailReq) (*emptypb.Empty, error) {
	region := in.GetRegion()
	gameId := in.GetGameId()
	participantId := in.GetParticipantId()

	if Session == nil {
		return nil, status.Error(codes.Internal, "Scylla Session is nil")
	}

	if gameId == 0 && participantId == 0 {
		return nil, status.Error(codes.InvalidArgument, "GameId or ParticipantId is empty")
	}

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	log.Printf(Info("Deleting participant detail"))
	filter := `DELETE FROM ` + region + `_lol.participant_detail WHERE gameId = ? AND participantId = ?`
	err := Session.Query(filter, gameId, participantId).Exec()
	if err != nil {
		if err == gocql.ErrNotFound {
			log.Printf("Participant detail not found: %v", Warn(gameId))
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Printf("Deleted Participant detail: %v", Success(gameId))
	return &emptypb.Empty{}, nil
}
