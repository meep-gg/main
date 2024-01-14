package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/joho/godotenv"
	. "meep.gg/http/riot"
	. "meep.gg/log"
	. "meep.gg/lol-api/summoner"
	. "meep.gg/protos/riot/lol/summoner/v1"
	"meep.gg/rates"

	. "meep.gg/lol-api/league"
	. "meep.gg/protos/riot/lol/league/v1"

	. "meep.gg/lol-api/match"
	. "meep.gg/protos/riot/lol/match/v1"
)

func main() {
	err := godotenv.Load() // This will load your .env file
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ImportKey(os.Getenv("RIOT_API_KEY"))
	prod := flag.Bool("prod", false, "Preps for production")
	addr := flag.String("addr", ":50061", "Address to listen on")
	redis := flag.String("redis", ":6379", "Redis address")
	flag.Parse()
	if *prod {
		log.SetOutput(io.Discard)
	}
	rates.InitRedis(*redis)

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	RegisterRiotMatchServiceServer(s, &MatchServer{})
	RegisterRiotLeagueServiceServer(s, &LeagueServer{})
	RegisterRiotSummonerServiceServer(s, &SummonerServer{})
	log.Printf("Lol Api Server listening at %v", Success(lis.Addr()))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", Error(err))
	}

}
