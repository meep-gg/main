package connect

import (
	"context"

	"google.golang.org/grpc"
	. "meep.gg/protos/gateway/lol/league/v1"
	. "meep.gg/protos/gateway/lol/match/v1"
	. "meep.gg/protos/gateway/lol/summoner/v1"
	. "meep.gg/protos/gateway/riot/account/v1"
	. "meep.gg/protos/riot/account/v1"
	. "meep.gg/protos/scylla/lol/league/v1"
	. "meep.gg/protos/scylla/lol/match/v1"
	. "meep.gg/protos/scylla/lol/participant/v1"
	. "meep.gg/protos/scylla/lol/participantdetail/v1"
	. "meep.gg/protos/scylla/lol/profile/v1"
	. "meep.gg/protos/scylla/riot/account/v1"
)

type GRPCConnList struct {
	RiotApi     *grpc.ClientConn
	RiotScylla  *grpc.ClientConn
	RiotGateway *grpc.ClientConn
	LolScylla   *grpc.ClientConn
	LolGateway  *grpc.ClientConn
}

// Account
var ScyllaAccountClient ScyllaAccountServiceClient
var GatewayAccountClient GatewayAccountServiceClient
var RiotAccountClient RiotAccountServiceClient

// Profile
var ScyllaProfileClient ScyllaProfileServiceClient

// League
var ScyllaLeagueClient ScyllaLeagueServiceClient
var GatewayLeagueClient GatewayLeagueServiceClient

// Summoner
var GatewaySummonerClient GatewaySummonerServiceClient

// Match
var ScyllaMatchClient ScyllaMatchServiceClient
var GatewayMatchClient GatewayMatchServiceClient

// Participant
var ScyllaParticipantClient ScyllaParticipantServiceClient

// ParticipantDetail
var ScyllaParticipantdetailClient ScyllaParticipantdetailServiceClient

func ClientConnect(ctx context.Context, list *GRPCConnList) {
	// Account
	RiotAccountClient = NewRiotAccountServiceClient(list.RiotApi)
	ScyllaAccountClient = NewScyllaAccountServiceClient(list.RiotScylla)
	GatewayAccountClient = NewGatewayAccountServiceClient(list.RiotGateway)

	// Profile
	ScyllaProfileClient = NewScyllaProfileServiceClient(list.LolScylla)

	// League
	ScyllaLeagueClient = NewScyllaLeagueServiceClient(list.LolScylla)
	GatewayLeagueClient = NewGatewayLeagueServiceClient(list.LolGateway)

	// Summoner
	GatewaySummonerClient = NewGatewaySummonerServiceClient(list.LolGateway)

	// Match
	GatewayMatchClient = NewGatewayMatchServiceClient(list.LolGateway)
	ScyllaMatchClient = NewScyllaMatchServiceClient(list.LolScylla)

	// Participant
	ScyllaParticipantClient = NewScyllaParticipantServiceClient(list.LolScylla)

	// ParticipantDetail
	ScyllaParticipantdetailClient = NewScyllaParticipantdetailServiceClient(list.LolScylla)
}
