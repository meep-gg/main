package summoner

import (
	"context"

	"google.golang.org/grpc"
	. "meep.gg/protos/riot/lol/summoner/v1"
	. "meep.gg/protos/scylla/lol/summoner/v1"
)

var ScyllaSummonerClient ScyllaSummonerServiceClient
var RiotSummonerClient RiotSummonerServiceClient

func SummonerConnect(ctx context.Context, scyllaConn, riotConn *grpc.ClientConn) {
	ScyllaSummonerClient = NewScyllaSummonerServiceClient(scyllaConn)
	RiotSummonerClient = NewRiotSummonerServiceClient(riotConn)
}
