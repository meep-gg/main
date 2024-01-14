package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	. "meep.gg/log"
	. "meep.gg/scylla"
	. "meep.gg/types/lol"

	. "meep.gg/lol-scylla/league"
	. "meep.gg/lol-scylla/match"
	. "meep.gg/lol-scylla/participant"
	. "meep.gg/lol-scylla/participantdetail"
	. "meep.gg/lol-scylla/profile"
	. "meep.gg/lol-scylla/summoner"

	. "meep.gg/protos/scylla/lol/league/v1"
	. "meep.gg/protos/scylla/lol/match/v1"
	. "meep.gg/protos/scylla/lol/participant/v1"
	. "meep.gg/protos/scylla/lol/participantdetail/v1"
	. "meep.gg/protos/scylla/lol/profile/v1"
	. "meep.gg/protos/scylla/lol/summoner/v1"
)

func main() {
	err := godotenv.Load() // This will load the .env file
	if err != nil {
		log.Fatalf(Error("Error loading .env file"))
	}
	prod := flag.Bool("prod", false, "Preps for production")
	addr := flag.String("addr", ":50062", "Address to listen on")
	flag.Parse()
	if *prod {
		log.SetOutput(io.Discard)
	}

	InitScylla(os.Getenv("SCYLLA_ADDR"))
	defer Session.Close()
	log.Printf("Lol ScyllaDB at %v", Highlight("g", *addr))

	s := grpc.NewServer()
	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Printf("Lol Scylla Server failed to listen: %v", Error(err))
	}

	RegisterScyllaSummonerServiceServer(s, &SummonerServer{})
	RegisterScyllaLeagueServiceServer(s, &LeagueServer{})
	RegisterScyllaMatchServiceServer(s, &MatchServer{})
	RegisterScyllaParticipantServiceServer(s, &ParticipantServer{})
	RegisterScyllaParticipantdetailServiceServer(s, &ParticipantdetailServer{})
	RegisterScyllaProfileServiceServer(s, &ProfileServer{})

	for platform := range Platform {
		CreateKeyspace(platform, "lol")
		// Create tables if they don't exist
		InitSummoner(platform)
		InitLeague(platform)
		InitMatch(platform)
		InitParticipant(platform)
		InitParticipantdetail(platform)
		InitProfile(platform)
	}

	log.Printf("Lol Scylla Server listening at %v", Success(lis.Addr()))
	if err := s.Serve(lis); err != nil {
		log.Printf("Lol Scylla Server failed to serve: %v", Error(err))
	}
}
