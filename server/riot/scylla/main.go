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
	. "meep.gg/protos/scylla/riot/account/v1"
	. "meep.gg/riot-scylla/account"
	. "meep.gg/scylla"
	. "meep.gg/types/lol"
)

func main() {
	err := godotenv.Load() // This will load the .env file
	if err != nil {
		log.Fatalf("Error loading .env file: %v", Error(err))
	}
	prod := flag.Bool("prod", false, "Preps for production")
	addr := flag.String("addr", ":50052", "Address to listen on")
	flag.Parse()
	if *prod {
		log.SetOutput(io.Discard)
	}

	InitScylla(os.Getenv("SCYLLA_ADDR"))
	defer Session.Close()
	log.Printf("Riot ScyllaDB at %v", Highlight("g", *addr))

	s := grpc.NewServer()
	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", Error(err))
	}

	RegisterScyllaAccountServiceServer(s, &AccountServer{})
	for platform := range Platform {
		CreateKeyspace(platform, "riot")
		CreateAccountTable(platform)
	}
	log.Printf("Scylla Server listening at %v", Success(lis.Addr()))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", Error(err))
	}

}
