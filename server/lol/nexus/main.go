package main

import (
	"context"
	"flag"
	"io"
	"log"
	"net"
	"os"
	"time"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	. "meep.gg/grpc"
	. "meep.gg/log"
	. "meep.gg/lol-nexus/connect"
	. "meep.gg/lol-nexus/match"
	. "meep.gg/lol-nexus/profile"
	. "meep.gg/protos/nexus/lol/match/v1"
	. "meep.gg/protos/nexus/lol/profile/v1"
)

func main() {
	err := godotenv.Load() // Load the .env file
	if err != nil {
		log.Fatalf(Error("Error loading .env file"))
	}
	prod := flag.Bool("prod", false, "Preps for production")
	addr := flag.String("addr", ":50064", "Address to listen on")
	flag.Parse()
	if *prod {
		log.SetOutput(io.Discard)
	}

	riotApi := os.Getenv("RIOT_API")
	riotScylla := os.Getenv("RIOT_SCYLLA")
	riotGateway := os.Getenv("RIOT_GATEWAY")
	lolScylla := os.Getenv("LOL_SCYLLA")
	lolGateway := os.Getenv("LOL_GATEWAY")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	riotApiConn := GRPCConnectRetry(ctx, riotApi)
	riotScyllaConn := GRPCConnectRetry(ctx, riotScylla)
	riotGatewayConn := GRPCConnectRetry(ctx, riotGateway)
	lolScyllaConn := GRPCConnectRetry(ctx, lolScylla)
	lolGatewayConn := GRPCConnectRetry(ctx, lolGateway)
	defer riotScyllaConn.Close()
	defer riotGatewayConn.Close()
	defer lolScyllaConn.Close()
	defer lolGatewayConn.Close()

	ClientConnect(ctx, &GRPCConnList{
		RiotApi:     riotApiConn,
		RiotScylla:  riotScyllaConn,
		RiotGateway: riotGatewayConn,
		LolScylla:   lolScyllaConn,
		LolGateway:  lolGatewayConn,
	})

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	RegisterMatchServiceServer(s, &MatchServer{})
	RegisterProfileServiceServer(s, &ProfileServer{})
	log.Printf("Lol Nexus Server listening at %v", Success(lis.Addr()))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", Error(err))
	}
}
