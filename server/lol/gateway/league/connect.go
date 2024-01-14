package league

import (
	"context"

	"google.golang.org/grpc"
	. "meep.gg/protos/riot/lol/league/v1"
	. "meep.gg/protos/scylla/lol/league/v1"
)

var ScyllaLeagueClient ScyllaLeagueServiceClient
var RiotLeagueClient RiotLeagueServiceClient

func LeagueConnect(ctx context.Context, scyllaConn, riotConn *grpc.ClientConn) {
	ScyllaLeagueClient = NewScyllaLeagueServiceClient(scyllaConn)
	RiotLeagueClient = NewRiotLeagueServiceClient(riotConn)
}
