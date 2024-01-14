package main

import (
	"flag"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"

	"os"

	"github.com/joho/godotenv"
	. "meep.gg/http/riot"
	. "meep.gg/log"
	. "meep.gg/protos/riot/account/v1"
	"meep.gg/rates"
	. "meep.gg/riot-api/account"
)

func main() {
	err := godotenv.Load() // This will load your .env file
	if err != nil {
		log.Fatalf("Error loading .env file: %v", Error(err))
	}
	ImportKey(os.Getenv("RIOT_API_KEY"))
	prod := flag.Bool("prod", false, "Preps for production")
	addr := flag.String("addr", ":50051", "Address to listen on")
	redis := flag.String("redis", ":6379", "Redis address")
	flag.Parse()
	if *prod {
		log.SetOutput(io.Discard)
	}
	rates.InitRedis(*redis)

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", Error(err))
	}

	s := grpc.NewServer()

	RegisterRiotAccountServiceServer(s, &AccountServer{})
	log.Printf("Riot Api Server listening at %v", Success(lis.Addr()))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", Error(err))
	}

}
