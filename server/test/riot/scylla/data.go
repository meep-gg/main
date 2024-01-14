package scylla

import . "meep.gg/protos/scylla/riot/account/v1"

var ScyllaAddr string = ":50052"

var TestAccount = &ScyllaAccount{
	Puuid:    "9WX5pU7pGi8EQvSFfC5hsiDJRpqdvmp1Lf57qbq1zSpSIiz38oakVBcFR_dMC2Rd6nbYf4n_pLI8sw",
	GameName: "Leong",
	TagLine:  "NA1",
}
