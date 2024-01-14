package match

import (
	"context"

	"google.golang.org/grpc"
	. "meep.gg/protos/riot/lol/match/v1"
	. "meep.gg/protos/scylla/lol/match/v1"
	. "meep.gg/protos/scylla/lol/participant/v1"
	. "meep.gg/protos/scylla/lol/participantdetail/v1"

	. "meep.gg/protos/scylla/lol/league/v1"
)

var ScyllaMatchClient ScyllaMatchServiceClient
var ScyllaParticipantClient ScyllaParticipantServiceClient
var ScyllaParticipantdetailClient ScyllaParticipantdetailServiceClient
var RiotMatchClient RiotMatchServiceClient

var ScyllaLeagueClient ScyllaLeagueServiceClient

func MatchConnect(ctx context.Context, scyllaConn, riotConn *grpc.ClientConn) {
	ScyllaMatchClient = NewScyllaMatchServiceClient(scyllaConn)
	ScyllaParticipantClient = NewScyllaParticipantServiceClient(scyllaConn)
	ScyllaParticipantdetailClient = NewScyllaParticipantdetailServiceClient(scyllaConn)
	RiotMatchClient = NewRiotMatchServiceClient(riotConn)

	ScyllaLeagueClient = NewScyllaLeagueServiceClient(scyllaConn)
}
