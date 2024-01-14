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
	. "meep.gg/log"
	. "meep.gg/protos/gateway/riot/account/v1"
	. "meep.gg/riot-gateway/account"
)

func main() {
	err := godotenv.Load() // Load the .env file
	if err != nil {
		log.Fatalf(Error("Error loading .env file"))
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	scylla := os.Getenv("RIOT_SCYLLA")
	riot := os.Getenv("RIOT_API")

	scyllaAccount, riotAccount := AccountConnect(ctx, scylla, riot)
	// Defer the closure of connections
	defer scyllaAccount.Close()
	defer riotAccount.Close()

	prod := flag.Bool("prod", false, "Preps for production")
	addr := flag.String("addr", ":50053", "Address to listen on")
	flag.Parse()
	if *prod {
		log.SetOutput(io.Discard)
	}

	s := grpc.NewServer()
	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Printf("Riot-Gateway-Server failed to listen: %v", Error(err))
	}

	RegisterGatewayAccountServiceServer(s, &AccountServer{})
	log.Printf("Riot Gateway Server listening at %v", Success(lis.Addr()))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Riot-Gateway-Server failed to serve: %v", Error(err))
	}
}
