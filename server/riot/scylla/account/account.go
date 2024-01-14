package account

import (
	"context"
	"log"

	"github.com/gocql/gocql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	. "meep.gg/log"
	. "meep.gg/protos/scylla/riot/account/v1"
	. "meep.gg/scylla"
	. "meep.gg/types/lol"
)

type AccountServer struct {
	ScyllaAccountServiceServer
}

func (s *AccountServer) UpsertAccount(ctx context.Context, in *ScyllaUpsertAccountReq) (*emptypb.Empty, error) {
	region := in.GetRegion()
	account := in.GetAccount()

	if account.Puuid == "" || account.GameName == "" || account.TagLine == "" {
		return nil, status.Error(codes.InvalidArgument, "Puuid or GameName or TagLine is empty")
	}

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	filter := `INSERT INTO ` + region + `_riot.account (puuid, gameName, tagLine) VALUES (?, ?, ?)`
	err := Session.Query(filter,
		account.Puuid,
		account.GameName,
		account.TagLine).Exec()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Printf("Updated Account: %v", Success(account.GameName+"#"+account.TagLine))
	return &emptypb.Empty{}, nil
}

func (s *AccountServer) GetAccount(ctx context.Context, in *ScyllaAccountReq) (*ScyllaAccount, error) {
	region := in.GetRegion()
	puuid := in.GetPuuid()

	if puuid == "" {
		return nil, status.Error(codes.InvalidArgument, "Puuid is empty")
	}

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	var account ScyllaAccount
	filter := `SELECT * FROM ` + region + `_riot.account WHERE puuid = ? LIMIT 1`
	query := Session.Query(filter, puuid)
	err := query.Scan(
		&account.Puuid,
		&account.GameName,
		&account.TagLine,
	)
	if err != nil {
		if err == gocql.ErrNotFound {
			return nil, status.Error(codes.NotFound, err.Error())
		} else {
			return nil, status.Error(codes.Internal, err.Error())

		}
	}
	log.Printf("Retrieved Account: %v", Success(account.GameName+"#"+account.TagLine))
	return &account, nil
}

func (s *AccountServer) DeleteAccount(ctx context.Context, in *ScyllaAccountReq) (*emptypb.Empty, error) {
	region := in.GetRegion()
	puuid := in.GetPuuid()

	if puuid == "" {
		return nil, status.Error(codes.InvalidArgument, "Puuid is empty")
	}

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	log.Printf("Deleting account")
	filter := `DELETE FROM ` + region + `_riot.account WHERE puuid = ?`
	err := Session.Query(filter, puuid).Exec()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	log.Printf("Deleted Account: %v", Success(puuid))
	return &emptypb.Empty{}, nil
}
