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
	gl "meep.gg/lol-gateway/league"
	gm "meep.gg/lol-gateway/match"
	gs "meep.gg/lol-gateway/summoner"
	. "meep.gg/protos/gateway/lol/league/v1"
	. "meep.gg/protos/gateway/lol/match/v1"
	. "meep.gg/protos/gateway/lol/summoner/v1"
)

func main() {
	err := godotenv.Load() // Load the .env file
	if err != nil {
		log.Fatalf(Error("Error loading .env file"))
	}
	prod := flag.Bool("prod", false, "Preps for production")
	addr := flag.String("addr", ":50063", "Address to listen on")
	flag.Parse()
	if *prod {
		log.SetOutput(io.Discard)
	}

	scylla := os.Getenv("LOL_SCYLLA")
	riot := os.Getenv("LOL_API")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	scyllaConn := GRPCConnectRetry(ctx, scylla)
	riotConn := GRPCConnectRetry(ctx, riot)
	defer scyllaConn.Close()
	defer riotConn.Close()

	gs.SummonerConnect(ctx, scyllaConn, riotConn)
	gl.LeagueConnect(ctx, scyllaConn, riotConn)
	gm.MatchConnect(ctx, scyllaConn, riotConn)

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	RegisterGatewayMatchServiceServer(s, &gm.GatewayMatchServer{})
	RegisterGatewaySummonerServiceServer(s, &gs.GatewaySummonerServer{})
	RegisterGatewayLeagueServiceServer(s, &gl.GatewayLeagueServer{})
	log.Printf("Lol Gateway Server listening at %v", Success(lis.Addr()))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", Error(err))
	}
}
