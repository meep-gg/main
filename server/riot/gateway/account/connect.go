package account

import (
	"context"

	"google.golang.org/grpc"
	. "meep.gg/grpc"
	. "meep.gg/protos/riot/account/v1"
	. "meep.gg/protos/scylla/riot/account/v1"
)

var ScyllaAccountClient ScyllaAccountServiceClient
var RiotAccountClient RiotAccountServiceClient

func AccountConnect(ctx context.Context, scylla, riot string) (*grpc.ClientConn, *grpc.ClientConn) {
	scyllaConn := GRPCConnectRetry(ctx, scylla)
	riotConn := GRPCConnectRetry(ctx, riot)

	ScyllaAccountClient = NewScyllaAccountServiceClient(scyllaConn)
	RiotAccountClient = NewRiotAccountServiceClient(riotConn)

	return scyllaConn, riotConn
}
