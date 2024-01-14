package account

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	. "meep.gg/log"
	. "meep.gg/protos/gateway/riot/account/v1"
	. "meep.gg/protos/riot/account/v1"
	. "meep.gg/protos/scylla/riot/account/v1"
	. "meep.gg/types/lol"
)

type AccountServer struct {
	GatewayAccountServiceServer
}

func (s *AccountServer) UpdateAccount(ctx context.Context, in *GatewayAccountReq) (*ScyllaAccount, error) {
	region := in.GetRegion()
	puuid := in.GetPuuid()
	gameName := in.GetGameName()
	tagLine := in.GetTagLine()

	regional, err := PlatformToRegional(region)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	var RiotAccount *RiotAccount
	if puuid != "" {
		log.Printf("getting account by puuid from riot: %v", Info(puuid))
		RiotAccount, err = RiotAccountClient.Puuid(ctx, &RiotAccountReq{
			Region: regional,
			Puuid:  puuid,
		})
		if err != nil {
			st := status.Convert(err)
			return nil, st.Err()
		}
	} else if gameName != "" && tagLine != "" {
		log.Printf("getting account by riotid from riot: %v", Info(gameName+"#"+tagLine))
		RiotAccount, err = RiotAccountClient.RiotId(ctx, &RiotAccountReq{
			Region:   regional,
			GameName: gameName,
			TagLine:  tagLine,
		})
		if err != nil {
			st := status.Convert(err)
			return nil, st.Err()
		}
	} else {
		return nil, status.Error(codes.InvalidArgument, "Puuid or RiotId is empty")
	}
	log.Printf("Got account from riot: %s\n", Info(RiotAccount))
	account := SerializeRiotAccount(RiotAccount)
	if RiotAccount != nil {
		_, err := ScyllaAccountClient.UpsertAccount(ctx, &ScyllaUpsertAccountReq{
			Region:  region,
			Account: account,
		})
		if err != nil {
			st := status.Convert(err)
			return nil, st.Err()
		}
		log.Printf("Updated and Inserted Account: %v",
			Success(RiotAccount.GameName+"#"+RiotAccount.TagLine))
		return account, nil
	} else {
		return nil, status.Error(codes.NotFound, "No Account found from Riot")
	}
}

func (s *AccountServer) GetAccount(ctx context.Context, in *GatewayAccountReq) (*ScyllaAccount, error) {
	region := in.GetRegion()
	puuid := in.GetPuuid()

	if !IsPlatform(region) {
		return nil, status.Error(codes.InvalidArgument, "Region is not of type Platform")
	}

	if puuid == "" {
		return nil, status.Error(codes.InvalidArgument, "Puuid is empty")
	}

	log.Printf("getting account by puuid from ScyllaDB: %v", Info(puuid))
	account, err := ScyllaAccountClient.GetAccount(ctx, &ScyllaAccountReq{
		Region: region,
		Puuid:  puuid,
	})

	if err != nil {
		st := status.Convert(err)
		if st.Code() != codes.NotFound {
			return nil, st.Err()
		}
	}

	if account == nil {
		regional, err := PlatformToRegional(region)
		log.Printf("getting account by puuid from riot: %v", Info(puuid))
		RiotAccount, err := RiotAccountClient.Puuid(ctx, &RiotAccountReq{
			Region: regional,
			Puuid:  puuid,
		})
		if err != nil {
			st := status.Convert(err)
			if st.Code() != codes.NotFound {
				return nil, st.Err()
			}
		}

		if RiotAccount != nil {
			scyllaAccount := SerializeRiotAccount(RiotAccount)
			_, err := ScyllaAccountClient.UpsertAccount(ctx, &ScyllaUpsertAccountReq{
				Region:  region,
				Account: scyllaAccount,
			})
			if err != nil {
				st := status.Convert(err)
				return nil, st.Err()
			}
			log.Printf("Updated and Inserted Account: %v",
				Success(RiotAccount.GameName+"#"+RiotAccount.TagLine))
			return account, nil
		} else {
			return nil, status.Error(codes.NotFound, "No Account found from Riot")
		}
	}

	log.Printf("Retrieved Account from ScyllaDB: %v", Success(account.GameName+"#"+account.TagLine))
	return account, nil
}
