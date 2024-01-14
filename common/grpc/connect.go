package grpc

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	. "meep.gg/log"
)

func GRPCConnect(addr string) *grpc.ClientConn {
	conn, err := grpc.Dial(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", Error(err))
	}
	return conn
}

func GRPCConnectRetry(ctx context.Context, addr string) *grpc.ClientConn {
	var conn *grpc.ClientConn
	var err error
	for i := 0; i < 10; i++ {
		conn, err = grpc.DialContext(ctx, addr,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithBlock())
		if err == nil {
			return conn
		}
		log.Printf("Failed to connect to %s, retrying...", Warn(addr))
		time.Sleep(2 * time.Second) // Wait for 2 seconds before retrying
	}
	log.Fatalf("Failed to connect to %s: %v", Error(addr), Error(err))
	return nil
}
